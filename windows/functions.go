package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

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

func lengthOfVideo(p string) string {
	ffprobe := GetFFPCommand()

	cmd := exec.Command(ffprobe, "-v", "quiet", "-print_format", "compact=print_section=0:nokey=1:escape=csv",
		"-show_entries", "format=duration", p)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	trueOut := strings.TrimSpace(string(out))
	seconds, _ := strconv.ParseFloat(trueOut, 64)
	tmp := int(math.Ceil(seconds))
	return internal.SecondsToTimeFormat(tmp)
}
