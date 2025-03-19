package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
)

func GetFFMPEGCommand() string {
	execPath, _ := os.Executable()
	cmdPath := filepath.Join(filepath.Dir(execPath), "ffmpeg.exe")

	return cmdPath
}

func GetFFPCommand() string {
	execPath, _ := os.Executable()
	cmdPath := filepath.Join(filepath.Dir(execPath), "ffprobe.exe")

	return cmdPath
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
