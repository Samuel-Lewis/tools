package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/kbinani/screenshot"
	log "github.com/sirupsen/logrus"
)

var (
	display  = kingpin.Flag("display", "List of displays to capture").Default("0").Short('d').Ints()
	interval = kingpin.Flag("interval", "Seconds between screenshots").Default("5").Short('i').Int()
	limit    = kingpin.Flag("limit", "Limit the number of screenshots (-1 is inifite)").Default("-1").Short('l').Int()
	prefix   = kingpin.Flag("prefix", "Prefix all file names (useful for sorting or naming)").Default("").Short('p').String()
	verbose  = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()
	dirName  = ""
)

func capture(d int, i int) {
	bounds := screenshot.GetDisplayBounds(d)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Panicln("Failed to take screenshot", err)
	}

	fileName := fmt.Sprintf("%vd%d_%06d.png", *prefix, d, i)
	path := filepath.Join(dirName, fileName)
	file, _ := os.Create(path)
	defer file.Close()
	png.Encode(file, img)
	log.Debugf("Screenshot: %v", path)
}

func record() {
	t := *limit
	for i := 0; i < t || t == -1; i++ {
		log.Infof("Taking screenshot %06d", i)

		for _, d := range *display {
			if d >= screenshot.NumActiveDisplays() {
				continue
			}
			capture(d, i)
		}

		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("display: %v", *display)
	log.Debugf("interval: %v", *interval)
	log.Debugf("limit: %v", *limit)
	log.Debugf("prefix: %v", *prefix)
	log.Debugf("verbosity: %v", log.GetLevel())

	dirName = fmt.Sprintf("raw_%v", time.Now().Format("2006-01-02T150405"))

	if *prefix != "" {
		dirName = fmt.Sprintf("%v_%v", dirName, *prefix)
		*prefix = fmt.Sprintf("%v_", *prefix)
	}

	log.Infof("Recording to %v", dirName)
	os.MkdirAll(dirName, os.ModePerm)

	record()
}
