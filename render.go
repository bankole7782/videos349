package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bankole7782/videos349/v3shared"
	"github.com/disintegration/imaging"
)

func render(instructions []map[string]string) string {
	rootPath, _ := v3shared.GetRootPath()

	defer func() {
		dirFIs, _ := os.ReadDir(rootPath)

		for _, dirFI := range dirFIs {
			if strings.HasPrefix(dirFI.Name(), ".") && dirFI.IsDir() {
				os.RemoveAll(filepath.Join(rootPath, dirFI.Name()))
			} else if strings.HasPrefix(dirFI.Name(), ".") && !dirFI.IsDir() {
				os.Remove(filepath.Join(rootPath, dirFI.Name()))
			}
		}
	}()

	ffmpeg := v3shared.GetFFMPEGCommand()

	videoParts := make([]string, 0)
	for _, instructionDesc := range instructions {
		// treat images
		if instructionDesc["kind"] == "image" {
			tmpFramesPath := filepath.Join(rootPath, "."+UntestedRandomString(10))
			os.MkdirAll(tmpFramesPath, 0777)
			endSeconds, _ := strconv.Atoi(instructionDesc["duration"])
			for seconds := 0; seconds < endSeconds; seconds++ {
				img, err := imaging.Open(instructionDesc["image"])
				if err != nil {
					panic(err)
				}
				outPath := filepath.Join(tmpFramesPath, strconv.Itoa(seconds)+".png")
				imaging.Save(img, outPath)
			}

			tmpImageVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
			_, err := exec.Command(ffmpeg, "-framerate", "1", "-i", filepath.Join(tmpFramesPath, "%d.png"),
				"-pix_fmt", "yuv420p", tmpImageVideoPath).CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return "error occured."
			}

			// join audio to video
			if instructionDesc["audio_optional"] != "" {
				tmpAudioPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp3")

				endSeconds, _ := strconv.Atoi(instructionDesc["duration"])
				audioLength := SecondsToTimeFormat(endSeconds)
				out, err := exec.Command(ffmpeg, "-ss", "0:0", "-to", audioLength, "-i",
					instructionDesc["audio_optional"], "-c", "copy", tmpAudioPath).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					return "error occured"
				}

				tmpImageVideoPath2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				_, err = exec.Command(ffmpeg, "-i", tmpImageVideoPath, "-i", tmpAudioPath,
					"-pix_fmt", "yuv420p", tmpImageVideoPath2).CombinedOutput()
				if err != nil {
					return "error occured"
				}
				tmpImageVideoPath = tmpImageVideoPath2

			} else {

				tmp2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				_, err = exec.Command(ffmpeg, "-f", "lavfi", "-i", "anullsrc=channel_layout=stereo:sample_rate=44100",
					"-i", tmpImageVideoPath, "-shortest", "-c:v", "copy", "-c:a", "aac", tmp2).CombinedOutput()
				if err != nil {
					fmt.Println(err)
					return "error occured"
				}

				tmpImageVideoPath = tmp2

			}

			videoParts = append(videoParts, tmpImageVideoPath)
		}

		// treat videos
		if instructionDesc["kind"] == "video" {
			videoPath := instructionDesc["video"]
			if !strings.HasSuffix(videoPath, ".mp4") {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err := exec.Command(ffmpeg, "-i", videoPath, "-pix_fmt", "yuv420p", tmpVideoPath).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					fmt.Println(err)
					return "error occurred"
				}

				videoPath = tmpVideoPath
			}

			videoLength := lengthOfVideo(videoPath)

			// check and do triming
			if instructionDesc["begin"] != "0:00" || instructionDesc["end"] != videoLength {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err := exec.Command(ffmpeg, "-ss", instructionDesc["begin"], "-to", instructionDesc["end"],
					"-i", videoPath, "-c", "copy", "-pix_fmt", "yuv420p", tmpVideoPath).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					fmt.Println(err)
					return "error occurred"
				}

				videoPath = tmpVideoPath
			}

			// add optional audio to video

			if instructionDesc["audio_optional"] != "" {

				tmpVideoPath3 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")

				out, err := exec.Command(ffmpeg, "-i", videoPath, "-i", instructionDesc["audio_optional"],
					"-c:v", "copy", "-map", "0:v:0", "-map", "1:a:0", "-shortest", tmpVideoPath3).CombinedOutput()
				if err != nil {
					fmt.Println(string(out))
					return "error occured"
				}

				videoPath = tmpVideoPath3
			}

			videoParts = append(videoParts, videoPath)
		}
	}

	for i, videoPart := range videoParts {
		// convert to 24 fps
		tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
		out, err := exec.Command(ffmpeg, "-i", videoPart, "-filter:v",
			"fps=24", tmpVideoPath).CombinedOutput()
		if err != nil {
			fmt.Println(out)
			return "error occurred."
		}
		videoParts[i] = tmpVideoPath
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
	out, err := exec.Command(ffmpeg, "-f", "concat", "-safe", "0",
		"-i", tmpVideosTxtPath, "-c", "copy", finalPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return "error occurred"
	}

	return finalPath
}
