package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
	log "github.com/sirupsen/logrus"
)

var (
	cache     = kingpin.Flag("cache", "Use a warmup run to test cache").Short('c').Bool()
	calibrate = kingpin.Flag("calibrate", "Runs a 1 second calibration to adjust for potential overhead").Short('a').Bool()
	command   = kingpin.Arg("command", "Command to execute").Required().Strings()
	output    = kingpin.Flag("output", "Enable output of original command to stdout and stderr").Short('o').Bool()
	runs      = kingpin.Flag("runs", "Number of instances to run").Default("1").Short('r').Int()
	shell     = kingpin.Flag("shell", "Shell environment to execute command").Default("sh -c").String()
	verbose   = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
)

// TimeCmd executes a command cmd and times its duration
func TimeCmd(c []string, label string) time.Duration {
	cmd := exec.Command(c[0], c[1:]...)

	var outc, errc bytes.Buffer
	cmd.Stdout = &outc
	cmd.Stderr = &errc

	if *output {
		outc.WriteTo(os.Stdout)
		errc.WriteTo(os.Stderr)
	}

	// Predeclare variables to save allocation time in time sensitive
	var start time.Time
	var elapsed time.Duration
	var err error

	log.Debugln("Running instance:", label)
	start = time.Now()
	err = cmd.Run()
	elapsed = time.Since(start)

	if err != nil {
		log.Fatalln("Command failed to run:", err)
	}

	return elapsed
}

func mean(times []time.Duration) time.Duration {
	var sum int64
	for _, v := range times {
		sum += v.Nanoseconds()
	}
	return time.Duration(sum / int64(len(times)))
}

func stddev(times []time.Duration, mean time.Duration) time.Duration {
	var total float64
	for _, v := range times {
		total += math.Pow(float64((v - mean).Nanoseconds()), 2)
	}
	return time.Duration(math.Sqrt(total / float64(len(times))))
}

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Set logging level
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	// Construct command
	c := append(strings.Split(*shell, " "), *command...)
	log.Debugf("cache: %v", *cache)
	log.Debugf("calibrate: %v", *calibrate)
	log.Debugf("command: %v", *command)
	log.Debugf("output: %v", *output)
	log.Debugf("runs: %v", *runs)
	log.Debugf("shell: %v", *shell)
	log.Debugf("verbosity: %v", log.GetLevel())

	// Offset calibration
	offset := time.Duration(0)
	if *calibrate {
		log.Debugf("Calibrating...")
		s, _ := time.ParseDuration("2s")
		offset = TimeCmd([]string{"sleep", "2"}, "CALIBRATION") - s
		log.Debugln("Calibration time set: ", offset)
	}

	// Run warm up run
	if *cache {
		TimeCmd(c, "CACHE")
	}

	// Actual timings
	times := make([]time.Duration, *runs)
	for i := 0; i < *runs; i++ {
		times[i] = TimeCmd(c, strconv.Itoa(i+1)) - offset
	}

	// Find summary
	m := mean(times)
	sd := stddev(times, m)

	fmt.Fprintln(os.Stdout, "mean", m)
	fmt.Fprintln(os.Stdout, "stddev", sd)
}
