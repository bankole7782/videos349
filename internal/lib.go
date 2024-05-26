package internal

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
)

func GetDefaultFontPath() string {
	fontPath := filepath.Join(os.TempDir(), "v349_font.ttf")
	os.WriteFile(fontPath, DefaultFont, 0777)
	return fontPath
}

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

func DoesPathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func UntestedRandomString(length int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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

func IsKeyNumeric(key glfw.Key) bool {
	numKeys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4,
		glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9}

	for _, numKey := range numKeys {
		if key == numKey {
			return true
		}
	}

	return false
}

func ExternalLaunch(p string) {
	cmd := "url.dll,FileProtocolHandler"
	runDll32 := filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")

	if runtime.GOOS == "windows" {
		exec.Command(runDll32, cmd, p).Run()
	} else if runtime.GOOS == "linux" {
		exec.Command("xdg-open", p).Run()
	}
}

func LengthOfVideo(p, ffprobePath string) string {
	cmd := exec.Command(ffprobePath, "-v", "quiet", "-print_format", "compact=print_section=0:nokey=1:escape=csv",
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

func SaveProjectCloseCallback(w *glfw.Window) {
	jsonBytes, _ := json.Marshal(Instructions)
	rootPath, _ := GetRootPath()
	outPath := filepath.Join(rootPath, ProjectName)
	os.WriteFile(outPath, jsonBytes, 0777)
}

func GetProjectFiles() []ToSortProject {
	// display some project names
	rootPath, _ := GetRootPath()
	dirEs, _ := os.ReadDir(rootPath)

	projectFiles := make([]ToSortProject, 0)
	for _, dirE := range dirEs {
		if dirE.IsDir() {
			continue
		}

		if strings.HasSuffix(dirE.Name(), ".v3p") {
			fInfo, _ := dirE.Info()
			projectFiles = append(projectFiles, ToSortProject{dirE.Name(), fInfo.ModTime()})
		}
	}

	slices.SortFunc(projectFiles, func(a, b ToSortProject) int {
		return b.ModTime.Compare(a.ModTime)
	})

	return projectFiles
}
