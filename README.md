# Peerflix Runner

A simple program that runs `go-peerflix` and `mpv` with a single command. Although `go-peerflix` supports this feature, it needs a lot of typing. I could've made a small PowerShell script to solve this, but I wanted to have some fancy progress indicators, which I don't know how to make in PowerShell.

## Requirements
1. Go
1. `go-peerflix` installed in PATH
	```
	go get -u github.com/Sioro-Neoku/go-peerflix
	go install github.com/Sioro-Neoku/go-peerflix
	```
1. `mpv` installed in PATH (for windows you can use `chocolatey`)

## Installation
```
go get github.com/mrg0lden/pfr
go install github.com/mrg0lden/pfr
```
### Usage
You can use torrent file links or magnet links
```
pfr magnet:replace-with-ur-link
```
