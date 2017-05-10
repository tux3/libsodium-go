libsodium-go
============
A binding library made in Go for the popular portable cryptography library [Sodium](https://download.libsodium.org/doc/).

This package magically hijacks the go build process to configure and build the native libsodium library, then link to it.
It's much more convenient to install as the whole build process is automated, but this relies on a major hack to run custom build scripts during "go build"!

This will currently only compile on Linux/Mac, since the build scripts are currently in Bash. I may later write some Windows build scripts too.

Purpose
-------
The goal of this binding library is to make use of Sodium in a more Go friendly matter.  And of course making it easier to make secure software.

Acknowledgments
----------------
This repository is a fork of https://github.com/GoKillers/libsodium-go

How to build
------------

Make sure you have a Linux or Mac system with bash, a C/C++ compiler toolchain, and the autotools.
Then simply go get the package, wait a couple minutes for the build to (silently) finish, and use it normally:

1. `go get github.com/tux3/libsodium-go`

License
---------
Originally Copyright 2015 - GoKillers
