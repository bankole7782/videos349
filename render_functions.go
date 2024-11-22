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
)

func hastenVideo(inVideoPath, outVideoPath, ffmpegCmd string) {
	if !DoesPathExists(inVideoPath) {
		panic(inVideoPath + " does not exists.")
	}

	rootPath, _ := GetRootPath()

	tmpPath := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath, 0777)

	exec.Command(ffmpegCmd, "-i", inVideoPath, filepath.Join(tmpPath, "%d.png")).Run()

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
		fmt.Println(string(out))
		panic(err)
	}

	os.RemoveAll(tmpPath)
	os.RemoveAll(tmpPath2)
}

func blackAndWhiteVideo(inVideoPath, outVideoPath, ffmpegCmd string) {
	if !DoesPathExists(inVideoPath) {
		panic(inVideoPath + " does not exists.")
	}

	rootPath, _ := GetRootPath()

	tmpPath := filepath.Join(rootPath, ".tmp_"+UntestedRandomString(10))
	os.MkdirAll(tmpPath, 0777)

	exec.Command(ffmpegCmd, "-i", inVideoPath, filepath.Join(tmpPath, "%d.png")).Run()

	dirFIs, err := os.ReadDir(tmpPath)
	if err != nil {
		panic(err)
	}

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
				workingImg, err := imaging.Open(workingImgPath)
				if err != nil {
					fmt.Println(err)
				}
				workingImg = imaging.Grayscale(workingImg)
				outPath := filepath.Join(tmpPath2, dirFIs[index].Name())
				imaging.Save(workingImg, outPath)
			}

		}(startIndex, endIndex, &wg)
	}
	wg.Wait()

	for index := (jobsPerThread * numberOfCPUS); index < len(dirFIs); index++ {
		workingImgPath := filepath.Join(tmpPath, dirFIs[index].Name())
		workingImg, err := imaging.Open(workingImgPath)
		if err != nil {
			fmt.Println(err)
		}
		workingImg = imaging.Grayscale(workingImg)
		outPath := filepath.Join(tmpPath2, dirFIs[index].Name())
		imaging.Save(workingImg, outPath)
	}

	tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")

	out, err := exec.Command(ffmpegCmd, "-framerate", "24", "-i", filepath.Join(tmpPath2, "%d.png"),
		tmpVideoPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	// extract audio from original video
	tmpAudioPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp3")
	out, err = exec.Command(ffmpegCmd, "-i", inVideoPath, tmpAudioPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	_, err = exec.Command(ffmpegCmd, "-i", tmpVideoPath, "-i", tmpAudioPath, outVideoPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	os.RemoveAll(tmpPath)
	os.RemoveAll(tmpPath2)
}
