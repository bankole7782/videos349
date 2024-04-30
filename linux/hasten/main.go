package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/bankole7782/videos349/internal"
	"github.com/otiai10/copy"
)

func main() {
	if len(os.Args) != 2 {
		panic("expecting only a video path")
	}
	inVideoFilename := os.Args[1]
	rootPath, _ := internal.GetRootPath()

	inVideoPath := filepath.Join(rootPath, inVideoFilename)

	if !internal.DoesPathExists(inVideoPath) {
		panic(inVideoPath + " does not exists.")
	}

	ffmpegCmd := GetFFMPEGCommand()

	tmpPath := filepath.Join(rootPath, ".tmp_"+internal.UntestedRandomString(10))
	os.MkdirAll(tmpPath, 0777)

	exec.Command(ffmpegCmd, "-i", inVideoPath, filepath.Join(tmpPath, "%d.png")).Run()
	fmt.Println("finished generating frames from your video")

	dirFIs, err := os.ReadDir(tmpPath)
	if err != nil {
		panic(err)
	}

	nameNums := make([]int, 0)
	for _, dirFI := range dirFIs {
		nameNum := strings.ReplaceAll(dirFI.Name(), ".png", "")
		nameNumInt, _ := strconv.Atoi(nameNum)

		rem := math.Mod(float64(nameNumInt), 2)
		rem2 := math.Mod(float64(nameNumInt), 3)
		rem3 := math.Mod(float64(nameNumInt), 5)

		if int(rem) == 0 || int(rem2) == 0 || int(rem3) == 0 {
			err = os.Remove(filepath.Join(tmpPath, fmt.Sprintf("%d.png", nameNumInt)))
			if err != nil {
				panic(err)
			}

		} else {
			nameNums = append(nameNums, nameNumInt)
		}
	}

	sort.Ints(nameNums)

	tmpPath2 := filepath.Join(rootPath, ".tmp_"+internal.UntestedRandomString(10))
	os.MkdirAll(tmpPath2, 0777)

	for i, num := range nameNums {
		oldPath := filepath.Join(tmpPath, fmt.Sprintf("%d.png", num))
		newPath := filepath.Join(tmpPath2, fmt.Sprintf("%d.png", i))
		copy.Copy(oldPath, newPath)
	}

	fmt.Println("finished preparing the frames")
	parts := strings.Split(inVideoFilename, ".")
	newVideoPath := filepath.Join(rootPath, parts[0]+"_hasten."+parts[1])
	os.RemoveAll(newVideoPath)

	out, err := exec.Command(ffmpegCmd, "-framerate", "24", "-i", filepath.Join(tmpPath2, "%d.png"),
		"-pix_fmt", "yuv420p", newVideoPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	fmt.Println("Hastened video: " + newVideoPath)

	os.RemoveAll(tmpPath)
	os.RemoveAll(tmpPath2)
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
