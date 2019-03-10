package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
	log "github.com/sirupsen/logrus"
)

var (
	history  = kingpin.Flag("history", "Number of files to back up for").Short('l').Default("3").Int()
	interval = kingpin.Flag("interval", "Number of seconds between backups").Short('i').Default("300").Int()
	path     = kingpin.Arg("path", "Path to file for backing up").Required().String()
	verbose  = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
)

func shift() {
	// Shift backups
	for i := *history - 1; i > 0; i-- {
		f := getPath(*path, i)
		t := getPath(*path, i+1)
		rm(t)
		move(f, t)
	}

	// New backup
	rm(getPath(*path, 1))
	copy(*path, getPath(*path, 1))
}

func move(f string, t string) {
	if _, err := os.Stat(f); err == nil {
		log.Debugf("mv %v -> %v", f, t)
		cmd := exec.Command("mv", "-f", f, t)
		err := cmd.Run()
		if err != nil {
			log.Errorln(err)
		}
	}
}

func rm(f string) {
	log.Debugf("rm %v", f)
	cmd := exec.Command("rm", "-rf", f)
	err := cmd.Run()
	if err != nil {
		log.Errorln(err)
	}
}

func copy(f string, t string) {
	if _, err := os.Stat(f); err == nil {
		log.Debugf("cp %v -> %v", f, t)
		cmd := exec.Command("cp", "-rf", f, t)
		err := cmd.Run()
		if err != nil {
			log.Errorln(err)
		}
	}
}

func getPath(path string, number int) string {
	return fmt.Sprintf("%s.bak%v", path, number)
}

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Set logging level
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}
	strings.TrimSuffix(*path, "/")

	log.Debugf("history: %v", *history)
	log.Debugf("interval: %v", *interval)
	log.Debugf("path: %v", *path)
	log.Debugf("verbose: %v", *verbose)
	log.Debugf("verbosity: %v", log.GetLevel())

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	for range ticker.C {
		log.Debugln("Tick")
		shift()
	}
}
