# `timeavg`

`timeavg` is a small tool written in [Go](https://golang.org/), which aims to extend upon the system call `time`. This is primarly done by allowing multiple runs of a program and averaging the result.

## Usage
1. Download the [latest `timeavg` release](https://github.com/Samuel-Lewis/Tools/releases).
2. Copy the `timeavg` binary to somewhere within your path
3. Use the command `timeavg <command>` from the command line
4. `timeavg --help` for more info

The command you want to time must be the final argument to the `timeavg` command. If you want to include arguments to the command, consider placing the entire command in quotes, eg
```
timeavg --runs=2 "ls -R /pictures | grep 'dog'"
```

If the command is super verbose or full of strange characters, consider placing it in its own script and calling that instead.


## Building from source
1. Ensure you have a working and up to date copy of [Go](https://golang.org/)
2. Clone [this repository](https://github.com/Samuel-Lewis/Tools) and navigate to the `Tools/timelapse` directory
3. Use `go build timelapse.go`


## TODO
- [ ] Add support for using system `time` for each timing run
- [ ] Test the shell options and flag on different systems
