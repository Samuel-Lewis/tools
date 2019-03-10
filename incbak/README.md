# `incbak`

`incbak` is a small tool written in [Go](https://golang.org/), which provides a way of incrementally backing up files in place. For a given `file`, it generates a `file.bak1` every so often.

## Usage
1. Download the [latest `incbak` release](https://github.com/Samuel-Lewis/Tools/releases).
2. Copy the `incbak` binary to somewhere within your path
3. Use the command `incbak [args] <file>` from the command line
4. `incbak --help` for more info

`.bak` files will be generated every few minutes (configurable). The most recent backup is the `.bak` with the smallest number (eg, `.bak1`). To remove ambiguity, specifiy an absolute path surrounded in quotes.


## TODO
- [ ] Add better relative paths
- [ ] Add timestamps to the `.bak` file names somehow?
- [ ] Windows compatability
