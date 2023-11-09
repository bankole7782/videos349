package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
)

func GetRootPath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "os error")
	}

	dd := os.Getenv("SNAP_USER_COMMON")

	if strings.HasPrefix(dd, filepath.Join(hd, "snap", "go")) || dd == "" {
		dd = filepath.Join(hd, "Videos349")
		os.MkdirAll(dd, 0777)
	}

	return dd, nil
}

func TimeFormatToSeconds(s string) int {
	// calculate total duration of the song
	parts := strings.Split(s, ":")
	minutesPartConverted, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	secondsPartConverted, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	totalSecondsOfSong := (60 * minutesPartConverted) + secondsPartConverted
	return totalSecondsOfSong
}

func SecondsToTimeFormat(seconds int) string {
	minutes := seconds / 60
	leftSeconds := math.Mod(float64(seconds), 60)

	return fmt.Sprintf("%d:%d", minutes, int(leftSeconds))
}

func UntestedRandomString(length int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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

func isKeyNumeric(key glfw.Key) bool {
	numKeys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4,
		glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9}

	for _, numKey := range numKeys {
		if key == numKey {
			return true
		}
	}

	return false
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
	return SecondsToTimeFormat(tmp)
}

func externalLaunch(p string) {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/C", "start", p).Run()
	} else if runtime.GOOS == "linux" {
		exec.Command("xdg-open", p).Run()
	}
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
