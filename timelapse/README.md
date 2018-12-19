# `timelapse`

`timelapse` is a small tool written in [Go](https://golang.org/) which provides a lightweight CLI for taking timelapses of your displays. Works with multiple displays and is relatively plug and play.

## Usage
1. Download the [latest `timelapse` release](https://github.com/Samuel-Lewis/Tools/releases).
2. Copy the `timelapse` binary to somewhere within your path
3. Use the command `timelapse` from the command line
4. Finish the timelapse by quitting the program (such as Ctrl+C)
4. `timelapse --help` for more info

There is the ability to append prefix to your files with `--prefix` which allows naming for sorting. Upon running the command, a screenshot of your display will be saved to a local folder called `raw_<timestamp>`. Within this, a collection of `.png`s will be created, named `d#_000000.png` where `#` is the corresponding display captured. Multiple displays can be specified by repeating the argument (`-d`), eg:
```
timelapse -d 0 -d 1
```

## Stitching with [ffmpeg](https://www.ffmpeg.org/)
Currently, the tool only outputs the frames at the full resolution, allowing you to assemble the video yourself. A common way to do this is using [ffmpeg](https://www.ffmpeg.org/). ffmpeg is a very powerful and configurable tool, and has lots of configurations. An *example* of stitching a video to a `.mp4` can be done by following these steps:
1. Download and install [ffmpeg](https://www.ffmpeg.org/)
2. Navigate to the directory created by the timelapse
3. Use the command:
```
ffmpeg -r 24 -i d0_%06d.png -vcodec libx264 -crf 25 -s 1920x1080 my_video.mp4
```

## Building from source
1. Ensure you have a working and up to date copy of [Go](https://golang.org/)
2. Clone [this repository](https://github.com/Samuel-Lewis/Tools) and navigate to the `Tools/timelapse` directory
3. Use `go build timelapse.go`

## TODO
- [ ] Add ability to stitch the frames together with [ffmpeg](https://www.ffmpeg.org/) from within the tool
- [ ] Add option for specifiying the output directory
