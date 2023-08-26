package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

func render(instructions []map[string]string) string {
	rootPath, _ := GetRootPath()
	videoParts := make([]string, 0)
	for _, instructionDesc := range instructions {
		// treat images
		if instructionDesc["kind"] == "image" {
			tmpFramesPath := filepath.Join(rootPath, "."+UntestedRandomString(10))
			os.MkdirAll(tmpFramesPath, 0777)
			endSeconds, _ := strconv.Atoi(instructionDesc["length (in seconds)"])
			for seconds := 0; seconds < endSeconds; seconds++ {
				img, err := imaging.Open(instructionDesc["image file"])
				if err != nil {
					panic(err)
				}
				outPath := filepath.Join(tmpFramesPath, strconv.Itoa(seconds)+".png")
				imaging.Save(img, outPath)
			}

			tmpImageVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
			_, err := exec.Command("ffmpeg", "-framerate", "1", "-i", filepath.Join(tmpFramesPath, "%d.png"),
				"-pix_fmt", "yuv420p", tmpImageVideoPath).CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return "error occured."
			}

			if instructionDesc["sound file (optional)"] != "" {
				tmpImageVideoPath2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")

				// join audio to video
				_, err = exec.Command("ffmpeg", "-i", tmpImageVideoPath, "-i", instructionDesc["sound file (optional)"],
					"-pix_fmt", "yuv420p", tmpImageVideoPath2).CombinedOutput()
				if err != nil {
					return "error occured"
				}
				tmpImageVideoPath = tmpImageVideoPath2
			}

			videoParts = append(videoParts, tmpImageVideoPath)
		}

		// treat videos
		if instructionDesc["kind"] == "video" {
			videoPath := instructionDesc["video file"]
			fmt.Println(videoPath)
			if !strings.HasSuffix(videoPath, ".mp4") {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err := exec.Command("ffmpeg", "-i", videoPath, "-pix_fmt", "yuv420p", tmpVideoPath).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					fmt.Println(err)
					return "error occurred"
				}

				videoPath = tmpVideoPath
			}

			// get the length of the video
			cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "compact=print_section=0:nokey=1:escape=csv",
				"-show_entries", "format=duration", videoPath)

			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(out))
				fmt.Println(err)
				return "error occured"
			}

			trueOut := strings.TrimSpace(string(out))
			seconds, _ := strconv.ParseFloat(trueOut, 64)
			tmp := int(math.Ceil(seconds))
			videoLength := SecondsToTimeFormat(tmp)

			// check and do triming
			if instructionDesc["begin (mm:ss)"] != "0:0" || instructionDesc["end (mm:ss)"] != videoLength {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err := exec.Command("ffmpeg", "-ss", instructionDesc["begin (mm:ss)"], "-to", instructionDesc["end (mm:ss)"],
					"-i", videoPath, "-c", "copy", "-pix_fmt", "yuv420p", tmpVideoPath).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					fmt.Println(err)
					return "error occurred"
				}

				videoPath = tmpVideoPath
			}

			videoParts = append(videoParts, videoPath)
		}
	}

	tmpVideosTxtPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".txt")
	outStr := ""
	for i, videoPart := range videoParts {
		outStr += fmt.Sprintf("file '%s'", videoPart)
		if i != len(videoParts)-1 {
			outStr += "\n"
		}
	}
	os.WriteFile(tmpVideosTxtPath, []byte(outStr), 0777)

	finalPath := filepath.Join(rootPath, "video_"+time.Now().Format("20060102T150405")+".mp4")
	out, err := exec.Command("ffmpeg", "-f", "concat", "-safe", "0",
		"-i", tmpVideosTxtPath, "-c", "copy", finalPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return "error occurred"
	}

	return finalPath
}
