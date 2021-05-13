# mdir-cmd

cmd of mdir

## build

```
$ GOOS=target-OS GOARCH=target-architecture go build -o mdir-cmd
```

## usage

```
$ mdir-cmd -h
NAME:
  mdir-cmd - cmd of mdir

USAGE:
  mdir-cmd /src /dest 2 3 4

GLOBAL OPTIONS:
  -f             force to mv/cp files (default: false)
  -c             copy instead of move (default: false)
  -p             show progress bar (default: false)
  --help, -h     show help (default: false)
  --version, -v  print the version (default: false)
```