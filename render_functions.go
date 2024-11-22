package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/otiai10/copy"
	"github.com/pkg/errors"
)

func hastenVideo(inVideoPath, outVideoPath, ffmpegCmd string) error {
	if !DoesPathExists(inVideoPath) {
		return errors.New(inVideoPath + " does not exists.")
	}

	rootPath, _ := GetRootPath()

	tmpPath := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath, 0777)

	exec.Command(ffmpegCmd, "-i", inVideoPath, filepath.Join(tmpPath, "%d.png")).Run()

	dirFIs, _ := os.ReadDir(tmpPath)

	nameNums := make([]int, 0)
	for _, dirFI := range dirFIs {
		nameNum := strings.ReplaceAll(dirFI.Name(), ".png", "")
		nameNumInt, _ := strconv.Atoi(nameNum)

		rem := math.Mod(float64(nameNumInt), 2)
		rem2 := math.Mod(float64(nameNumInt), 3)
		rem3 := math.Mod(float64(nameNumInt), 5)

		if int(rem) == 0 || int(rem2) == 0 || int(rem3) == 0 {
			os.Remove(filepath.Join(tmpPath, fmt.Sprintf("%d.png", nameNumInt)))
		} else {
			nameNums = append(nameNums, nameNumInt)
		}
	}

	sort.Ints(nameNums)

	tmpPath2 := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath2, 0777)

	for i, num := range nameNums {
		oldPath := filepath.Join(tmpPath, fmt.Sprintf("%d.png", num))
		newPath := filepath.Join(tmpPath2, fmt.Sprintf("%d.png", i))
		copy.Copy(oldPath, newPath)
	}

	out, err := exec.Command(ffmpegCmd, "-framerate", "24", "-i", filepath.Join(tmpPath2, "%d.png"),
		outVideoPath).CombinedOutput()
	if err != nil {
		return errors.New(string(out) + "\n" + err.Error())
	}

	os.RemoveAll(tmpPath)
	os.RemoveAll(tmpPath2)

	return nil
}

func blackAndWhiteVideo(inVideoPath, outVideoPath, ffmpegCmd string) error {
	if !DoesPathExists(inVideoPath) {
		errors.New(inVideoPath + " does not exists.")
	}

	rootPath, _ := GetRootPath()

	tmpPath := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath, 0777)

	exec.Command(ffmpegCmd, "-i", inVideoPath, filepath.Join(tmpPath, "%d.png")).Run()

	dirFIs, _ := os.ReadDir(tmpPath)

	tmpPath2 := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath2, 0777)

	numberOfCPUS := runtime.NumCPU()
	jobsPerThread := int(math.Floor(float64(len(dirFIs)) / float64(numberOfCPUS)))
	var wg sync.WaitGroup

	for threadIndex := 0; threadIndex < numberOfCPUS; threadIndex++ {
		wg.Add(1)

		startIndex := threadIndex * jobsPerThread
		endIndex := (threadIndex + 1) * jobsPerThread

		go func(startIndex, endIndex int, wg *sync.WaitGroup) {
			defer wg.Done()

			for index := startIndex; index < endIndex; index++ {
				workingImgPath := filepath.Join(tmpPath, dirFIs[index].Name())
				workingImg, _ := imaging.Open(workingImgPath)
				workingImg = imaging.Grayscale(workingImg)
				outPath := filepath.Join(tmpPath2, dirFIs[index].Name())
				imaging.Save(workingImg, outPath)
			}

		}(startIndex, endIndex, &wg)
	}
	wg.Wait()

	for index := (jobsPerThread * numberOfCPUS); index < len(dirFIs); index++ {
		workingImgPath := filepath.Join(tmpPath, dirFIs[index].Name())
		workingImg, _ := imaging.Open(workingImgPath)
		workingImg = imaging.Grayscale(workingImg)
		outPath := filepath.Join(tmpPath2, dirFIs[index].Name())
		imaging.Save(workingImg, outPath)
	}

	tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")

	out, err := exec.Command(ffmpegCmd, "-framerate", "24", "-i", filepath.Join(tmpPath2, "%d.png"),
		tmpVideoPath).CombinedOutput()
	if err != nil {
		return errors.New(string(out) + "\n" + err.Error())
	}

	// extract audio from original video
	tmpAudioPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp3")
	out, err = exec.Command(ffmpegCmd, "-i", inVideoPath, "-q:a", "0", "-map", "a", tmpAudioPath).CombinedOutput()
	if err != nil {
		return errors.New(string(out) + "\n" + err.Error())
	}

	out, err = exec.Command(ffmpegCmd, "-i", tmpVideoPath, "-i", tmpAudioPath,
		outVideoPath).CombinedOutput()
	if err != nil {
		return errors.New(string(out) + "\n" + err.Error())
	}

	os.RemoveAll(tmpPath)
	os.RemoveAll(tmpPath2)

	return nil
}
