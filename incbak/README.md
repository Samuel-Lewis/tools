# `incbak`

`incbak` is a small tool written in [Go](https://golang.org/), which provides a way of incrementally backing up files in place. For a given `file`, it generates a `file.bak` every so often.

## Usage
1. Download the [latest `incbak` release](https://github.com/Samuel-Lewis/Tools/releases).
2. Copy the `incbak` binary to somewhere within your path
3. Use the command `incbak [args] <file>` from the command line
4. `incbak --help` for more info

`.bak` files will be generated every few minutes (configurable). Each `.bak` is generated with a timestamp of when the backup was made. To revert, to the most recent backup, go to the latest timestamp. To remove ambiguity when running the command, specifiy an absolute path surrounded in quotes.


## TODO
- [ ] Add better relative paths
- [x] Add timestamps to the `.bak` file names somehow?
- [ ] Windows compatability
