package main

import (
	"os"
	"path/filepath"

	"github.com/bankole7782/videos349/internal"
)

func GetFFMPEGCommand() string {
	homeDir, _ := os.UserHomeDir()

	ffmegDir := filepath.Join(homeDir, ".v349")
	outPath := filepath.Join(ffmegDir, "ffmpeg.exe")
	if !internal.DoesPathExists(outPath) {
		os.MkdirAll(ffmegDir, 0777)

		os.WriteFile(outPath, ffmpegBytes, 0777)
	}

	return outPath
}

func GetFFPCommand() string {
	homeDir, _ := os.UserHomeDir()

	ffmegDir := filepath.Join(homeDir, ".v349")
	outPath := filepath.Join(ffmegDir, "ffprobe.exe")
	if !internal.DoesPathExists(outPath) {
		os.MkdirAll(ffmegDir, 0777)

		os.WriteFile(outPath, ffprobeBytes, 0777)
	}

	return outPath
}
