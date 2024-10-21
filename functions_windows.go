package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
)

//go:embed "ffmpeg/ffmpeg.exe"
var ffmpegBytes []byte

//go:embed "ffmpeg/ffprobe.exe"
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

func PickVideoFile() string {
	filename, err := dialog.File().Filter("MP4 Video", "mp4").Filter("WEBM Video", "webm").Filter("MKV Video", "mkv").Load()
	if filename == "" || err != nil {
		log.Println(err)
		return ""
	}
	return filename
}

func PickAudioFile() string {
	filename, err := dialog.File().Filter("MP3 Audio", "mp3").Filter("FLAC Audio", "flac").
		Filter("WAV Audio", "wav").Load()
	if filename == "" || err != nil {
		log.Println(err)
		return ""
	}
	return filename
}

func PickImageFile() string {
	filename, err := dialog.File().Filter("PNG Image", "png").Filter("JPEG Image", "jpg").Load()
	if filename == "" || err != nil {
		log.Println(err)
		return ""
	}
	return filename
}
