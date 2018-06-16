#!/bin/bash

mkdir -p ./builds

program="dockerclean"
winprogram="dockerclean.exe"
tag="$1"

# build Linux 64bit program
env GOOS=linux GOARCH=amd64 go build -o $program main.go
linux64=$(printf "%s_%s_linux_amd64.tar.gz" "$program" "$tag")
tar -cvzf ./builds/$linux64 $program

# build Mac 64bit program
env GOOS=darwin GOARCH=amd64 go build -o $program main.go
maczip=$(printf "%s_%s_darwin_amd64.zip" "$program" "$tag")
zip -r ./builds/$maczip $program

# build Windows 64bit program
env GOOS=windows GOARCH=amd64 go build -o $winprogram main.go
winzip64=$(printf "%s_%s_windows_amd64.zip" "$program" "$tag")
zip -r ./builds/$winzip64 $winprogram

# build FreeBSD 64bit program
env GOOS=freebsd GOARCH=amd64 go build -o $program main.go
freebsd64=$(printf "%s_%s_freebsd_amd64.tar.gz" "$program" "$tag")
tar -cvzf ./builds/$freebsd64 $program