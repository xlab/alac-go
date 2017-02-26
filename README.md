alac-go
=======

The package provides Go bindings for ALAC decoder (C version by David Hammerton) with community patches.<br />
All the binding code has automatically been generated with rules defined in [alac.yml](/alac.yml).

### Usage

```
$ go get github.com/xlab/alac-go/alac
```

### Demo

There is a player implemented in Go that can read M4A ALAC files and play them via [portaudio-go](https://github.com/xlab/portaudio-go). So you will need to get portaudio installed first. [03-Morcheeba-Otherwise.m4a](http://dl.xlab.is/music/alac/03-Morcheeba-Otherwise.m4a) direct link.

```
$ brew install portaudio
$ go get github.com/xlab/alac-go/cmd/alac-player

$ ./alac-player 03-Morcheeba-Otherwise.m4a
ALAC header: 616C616300000000000010000010280A0E0200FF000034DF000DBCB40000AC44
Audio duration: 224.026s
[==>-----------------------------------------------------------------------------------------------------------------] 2.20 %
```

### Rebuilding the package

You will need to get the [c-for-go](https://git.io/c-for-go) tool installed first.

```
$ git clone https://github.com/xlab/alac-go && cd alac-go
$ make clean
$ make
```
