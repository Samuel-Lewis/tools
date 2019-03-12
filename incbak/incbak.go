package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
	log "github.com/sirupsen/logrus"
)

var (
	backups  = kingpin.Flag("backups", "Number of files to back up for").Short('b').Default("3").Int()
	interval = kingpin.Flag("interval", "Number of seconds between backups").Short('i').Default("300").Int()
	path     = kingpin.Arg("path", "Path to file for backing up").Required().String()
	verbose  = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
)

func shift() {
	// Create backup of original
	copy(*path, getPathNow(*path))

	// Select all .bak.*
	matches, err := filepath.Glob(fmt.Sprintf("%v.bak.*", *path))
	if err != nil {
		log.Errorln(err)
	}
	sort.Strings(matches)

	// Remove extras
	for i := 0; i < len(matches)-*backups; i++ {
		rm(matches[i])
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

func getPathNow(path string) string {
	n := time.Now().Format("2006.01.02.15.04.05")
	return fmt.Sprintf("%v.bak.%v", path, n)
}

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Set logging level
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}
	strings.TrimSuffix(*path, "/")

	log.Debugf("backups: %v", *backups)
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
