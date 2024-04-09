package main

import (
	_ "embed"
)

//go:embed Roboto-Light.ttf
var DefaultFont []byte

//go:embed "ffmpeg/ffmpeg.exe"
var ffmpegBytes []byte

//go:embed "ffmpeg/ffprobe.exe"
var ffprobeBytes []byte
