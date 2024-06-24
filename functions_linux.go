package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetFFMPEGCommand() string {
	var cmdPath string
	begin := os.Getenv("SNAP")
	cmdPath = "ffmpeg"
	if begin != "" && !strings.HasPrefix(begin, "/snap/go/") {
		cmdPath = filepath.Join(begin, "bin", "ffmpeg")
	}

	return cmdPath
}

func GetFFPCommand() string {
	var cmdPath string
	begin := os.Getenv("SNAP")
	cmdPath = "ffprobe"
	if begin != "" && !strings.HasPrefix(begin, "/snap/go/") {
		cmdPath = filepath.Join(begin, "bin", "ffprobe")
	}

	return cmdPath
}

func GetPickerPath() string {
	homeDir, _ := os.UserHomeDir()
	var cmdPath string
	begin := os.Getenv("SNAP")
	cmdPath = filepath.Join(homeDir, "bin", "fpicker")
	if begin != "" && !strings.HasPrefix(begin, "/snap/go/") {
		cmdPath = filepath.Join(begin, "bin", "fpicker")
	}

	return cmdPath
}

func pickFileUbuntu(exts string) string {
	fPickerPath := GetPickerPath()

	rootPath, _ := GetRootPath()
	cmd := exec.Command(fPickerPath, rootPath, exts)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return strings.TrimSpace(string(out))
}

func PickVideoFile() string {
	return pickFileUbuntu("mp4|mkv|webm")
}

func PickAudioFile() string {
	return pickFileUbuntu("mp3|flac|wav")
}

func PickImageFile() string {
	return pickFileUbuntu("png|jpg")
}
