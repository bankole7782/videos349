package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

func Render(instructions []map[string]string, ffmpeg, ffprobe string) (string, error) {
	rootPath, _ := GetRootPath()

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

	videoParts := make([]string, 0)
	for i, instructionDesc := range instructions {
		// treat images
		if instructionDesc["kind"] == "image" {
			tmpFramesPath := filepath.Join(rootPath, "."+UntestedRandomString(10))
			os.MkdirAll(tmpFramesPath, 0777)
			var endSeconds int
			if tmpAudio, ok := instructionDesc["audio"]; ok && tmpAudio != "" {
				tmp1 := TimeFormatToSeconds(instructionDesc["audio_begin"])
				tmp2 := TimeFormatToSeconds(instructionDesc["audio_end"])
				endSeconds = tmp2 - tmp1
			} else {
				endSeconds, _ = strconv.Atoi(instructionDesc["duration"])
			}
			for seconds := 0; seconds < endSeconds; seconds++ {
				img, err := imaging.Open(instructionDesc["image"])
				if err != nil {
					return "", errors.New(err.Error())
				}
				outPath := filepath.Join(tmpFramesPath, strconv.Itoa(seconds)+".png")
				imaging.Save(img, outPath)
			}

			tmpImageVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
			out, err := exec.Command(ffmpeg, "-framerate", "1", "-i", filepath.Join(tmpFramesPath, "%d.png"),
				tmpImageVideoPath).CombinedOutput()
			if err != nil {
				return "", errors.New(string(out) + "\n" + err.Error())
			}

			// convert to 24 fps
			tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
			out, err = exec.Command(ffmpeg, "-i", tmpImageVideoPath, "-filter:v",
				"fps=24", tmpVideoPath).CombinedOutput()
			if err != nil {
				return "", errors.New(string(out) + "\n" + err.Error())
			}

			tmpImageVideoPath = tmpVideoPath

			// join audio to video
			if tmpAudio, ok := instructionDesc["audio"]; ok && tmpAudio != "" {

				startAudioPath := instructionDesc["audio"]
				// do conversion to mp3 if necessary
				if strings.HasSuffix(instructionDesc["audio"], ".wav") || strings.HasSuffix(instructionDesc["audio"], ".flac") {
					tmpAudioPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp3")
					out, err := exec.Command(ffmpeg, "-i", instructionDesc["audio"], tmpAudioPath).CombinedOutput()
					if err != nil {
						return "", errors.New(string(out) + "\n" + err.Error())
					}

					startAudioPath = tmpAudioPath
				}

				// trim audio if necessary
				audioLength := LengthOfVideo(startAudioPath, ffprobe)
				if audioLength != instructionDesc["audio_end"] || instructionDesc["audio_begin"] != "0:00" {
					tmpAudioPath2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp3")

					out, err := exec.Command(ffmpeg, "-ss", instructionDesc["audio_begin"], "-to",
						instructionDesc["audio_end"], "-i", startAudioPath, "-c", "copy", tmpAudioPath2).CombinedOutput()
					if err != nil {
						return "", errors.New(string(out) + "\n" + err.Error())
					}

					startAudioPath = tmpAudioPath2
				}

				tmpImageVideoPath2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err = exec.Command(ffmpeg, "-i", tmpImageVideoPath, "-i", startAudioPath,
					tmpImageVideoPath2).CombinedOutput()
				if err != nil {
					return "", errors.New(string(out) + "\n" + err.Error())
				}
				tmpImageVideoPath = tmpImageVideoPath2

			} else {

				tmp2 := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err = exec.Command(ffmpeg, "-f", "lavfi", "-i", "anullsrc=channel_layout=stereo:sample_rate=44100",
					"-i", tmpImageVideoPath, "-shortest", "-c:v", "copy", "-c:a", "aac", tmp2).CombinedOutput()
				if err != nil {
					return "", errors.New(string(out) + "\n" + err.Error())
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
				out, err := exec.Command(ffmpeg, "-i", videoPath, tmpVideoPath).CombinedOutput()
				if err != nil {
					return "", errors.New(string(out) + "\n" + err.Error())
				}

				videoPath = tmpVideoPath
			}

			videoLength := LengthOfVideo(videoPath, ffprobe)

			// check and do triming
			if instructionDesc["begin"] != "0:00" || instructionDesc["end"] != videoLength {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				out, err := exec.Command(ffmpeg, "-ss", instructionDesc["begin"], "-to", instructionDesc["end"],
					"-i", videoPath, "-c", "copy", tmpVideoPath).CombinedOutput()
				if err != nil {
					return "", errors.New(string(out) + "\n" + err.Error())
				}

				videoPath = tmpVideoPath
			}

			// check and do speed up
			if instructionDesc["speedup"] == "true" {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				hastenVideo(videoPath, tmpVideoPath, ffmpeg)
				videoPath = tmpVideoPath
			}

			if instructionDesc["blackwhite"] == "true" {
				tmpVideoPath := filepath.Join(rootPath, "."+UntestedRandomString(10)+".mp4")
				blackAndWhiteVideo(videoPath, tmpVideoPath, ffmpeg)
				videoPath = tmpVideoPath
			}

			videoParts = append(videoParts, videoPath)
		}

		RenderProgress = i + 1
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

	stub := strings.ReplaceAll(ProjectName, ".v3p", "") + "_" + time.Now().Format("20060102T150405") + ".mp4"
	finalPath := filepath.Join(rootPath, "renders", stub)

	out, err := exec.Command(ffmpeg, "-f", "concat", "-safe", "0",
		"-i", tmpVideosTxtPath, "-c", "copy", finalPath).CombinedOutput()
	if err != nil {
		return "", errors.New(string(out) + "\n" + err.Error())
	}

	return finalPath, nil
}
