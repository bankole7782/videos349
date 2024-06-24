package main

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed "ffmpeg/windows/ffmpeg.exe"
var ffmpegBytes []byte

//go:embed "ffmpeg/windows/ffprobe.exe"
var ffprobeBytes []byte

func GetFFMPEGCommand() string {
	homeDir, _ := os.UserHomeDir()

	ffmegDir := filepath.Join(homeDir, ".v349")
	outPath := filepath.Join(ffmegDir, "ffmpeg.exe")
	if !DoesPathExists(outPath) {
		os.MkdirAll(ffmegDir, 0777)

		os.WriteFile(outPath, ffmpegBytes, 0777)
	}

	return outPath
}

func GetFFPCommand() string {
	homeDir, _ := os.UserHomeDir()

	ffmegDir := filepath.Join(homeDir, ".v349")
	outPath := filepath.Join(ffmegDir, "ffprobe.exe")
	if !DoesPathExists(outPath) {
		os.MkdirAll(ffmegDir, 0777)

		os.WriteFile(outPath, ffprobeBytes, 0777)
	}

	return outPath
}
