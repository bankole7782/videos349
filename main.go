package main

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	myWindow := app.NewWindow("Videos349: A video editor")

	rootPath, err := GetRootPath()
	if err != nil {
		panic(err)
	}

	openWDBtn := widget.NewButton("Open Working Directory", func() {
		if runtime.GOOS == "windows" {
			exec.Command("cmd", "/C", "start", rootPath).Run()
		} else if runtime.GOOS == "linux" {
			exec.Command("xdg-open", rootPath).Run()
		}
	})

	openFileInDefaultViewer := func(p string) {
		if runtime.GOOS == "windows" {
			exec.Command("cmd", "/C", "start", p).Run()
		} else if runtime.GOOS == "linux" {
			exec.Command("xdg-open", p).Run()
		}
	}

	saeBtn := widget.NewButton("sae.ng", func() {
		openFileInDefaultViewer("https://sae.ng")
	})

	aboutBtn := widget.NewButton("About Us", func() {
		img, _, err := image.Decode(bytes.NewReader(SaeLogoBytes))
		if err != nil {
			panic(err)
		}
		logoImage := canvas.NewImageFromImage(img)
		logoImage.FillMode = canvas.ImageFillOriginal

		boxes := container.NewVBox(
			container.NewCenter(logoImage),
			widget.NewLabelWithStyle("Brought to You with Love by", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			saeBtn,
		)
		dialog.ShowCustom("About keys117", "Close", boxes, myWindow)
	})
	topBar := container.NewHBox(openWDBtn)

	instructions := make([]map[string]string, 0)

	instructionsBox := container.NewVBox()

	bottomBar := container.NewHBox()

	updateInstructionsBox := func() {
		instructionsBox.RemoveAll()

		viewImageAssetBtn := func(i int) *widget.Button {
			val := i
			return widget.NewButton("View Image Asset", func() {
				imagePath := instructions[val]["image file"]
				openFileInDefaultViewer(imagePath)
			})
		}

		viewAudioAssetBtn := func(i int) *widget.Button {
			val := i
			return widget.NewButton("View Audio Asset", func() {
				imagePath := instructions[val]["sound file (optional)"]
				openFileInDefaultViewer(imagePath)
			})
		}

		viewVideoAssetBtn := func(i int) *widget.Button {
			val := i
			return widget.NewButton("View Video Asset", func() {
				videoPath := instructions[val]["video file"]
				openFileInDefaultViewer(videoPath)
			})
		}

		for k, instructionsDesc := range instructions {

			outStr := "kind: " + instructionsDesc["kind"] + "\n"
			for inputName, inputValue := range instructionsDesc {
				if inputName == "kind" {
					continue
				}
				outStr += inputName + " :" + inputValue + "\n"
			}
			innerBtnsBox := container.NewVBox()

			if instructionsDesc["kind"] == "image" {
				innerBtnsBox.Add(viewImageAssetBtn(k))
				if instructionsDesc["sound file (optional)"] != "" {
					innerBtnsBox.Add(viewAudioAssetBtn(k))
				}
			}

			if instructionsDesc["kind"] == "video" {
				innerBtnsBox.Add(viewVideoAssetBtn(k))
			}

			innerBox := container.NewHBox(
				widget.NewLabel(strconv.Itoa(k+1)),
				widget.NewLabel(outStr[:len(outStr)-1]), innerBtnsBox,
			)
			instructionsBox.Add(innerBox)
		}

	}

	inChannel := make(chan bool)
	renderDialogObj := dialog.NewProgressInfinite("Rendering", "Rendering your video.\nPlease Wait.", myWindow)

	go func() {
		for {
			<-inChannel
			time.Sleep(5 * time.Second)
			renderDialogObj.Hide()
		}
	}()

	renderBtn := widget.NewButton("Render", func() {
		inChannel <- true
		renderDialogObj.Show()
	})

	renderBtn.Importance = widget.HighImportance

	updateBottomBar := func() {
		removeBtnLists := make([]string, 0)
		for k := range instructions {
			removeBtnLists = append(removeBtnLists, strconv.Itoa(k+1))
		}

		// update bottom bar
		assetsSelect := widget.NewSelect(removeBtnLists, nil)
		removeAssetBtn := widget.NewButton("Remove", func() {
			selected, _ := strconv.Atoi(assetsSelect.Selected)

			// instructions = append(instructions[selected-1:], instructions[selected:]...)
			instructions = slices.Delete(instructions, selected-1, selected)

			removeBtnLists := make([]string, 0)
			for k := range instructions {
				removeBtnLists = append(removeBtnLists, strconv.Itoa(k+1))
			}

			assetsSelect.Options = removeBtnLists
			assetsSelect.Refresh()
			updateInstructionsBox()
		})

		bottomBar.RemoveAll()
		bottomBar.Add(widget.NewLabel("Remove Asset"))
		bottomBar.Add(assetsSelect)
		bottomBar.Add(removeAssetBtn)

		bottomBar.Add(widget.NewSeparator())
		bottomBar.Add(renderBtn)
	}

	addImageBtn := widget.NewButton("Add Image", func() {
		pngFiles := getFilesOfType(rootPath, ".png")
		mp3Files := getFilesOfType(rootPath, ".mp3")
		imageForm := widget.NewForm()
		imageForm.Append("image file", widget.NewSelect(pngFiles, nil))
		imageForm.Append("sound file (optional)", widget.NewSelect(mp3Files, nil))
		lengthEntry := widget.NewEntry()
		lengthEntry.SetText("5")
		imageForm.Append("length (in seconds)", lengthEntry)

		callBack := func(ok bool) {
			if ok {
				inputs := getFormInputs(imageForm.Items)

				// "image file" is compulsory
				if inputs["image file"] == "" {
					return
				}
				// complete the paths
				for k, v := range inputs {
					if strings.Contains(k, "file") && v != "" {
						inputs[k] = filepath.Join(rootPath, v)
					}
				}
				inputs["kind"] = "image"

				instructions = append(instructions, inputs)
				updateInstructionsBox()
				updateBottomBar()
			}
		}

		dialog.ShowCustomConfirm("Add Image Configuration", "Add", "Close", imageForm, callBack, myWindow)
	})

	addVideoBtn := widget.NewButton("Add Video", func() {

		endEntry := widget.NewEntry()
		updateEndEntry := func(selected string) {
			fullPath := filepath.Join(rootPath, selected)
			cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "compact=print_section=0:nokey=1:escape=csv",
				"-show_entries", "format=duration", fullPath)

			out, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}

			trueOut := strings.TrimSpace(string(out))
			seconds, _ := strconv.ParseFloat(trueOut, 64)
			tmp := int(math.Ceil(seconds))
			endEntry.SetText(SecondsToTimeFormat(tmp))
		}

		videoFiles := getFilesOfType(rootPath, ".mp4")
		mkvFiles := getFilesOfType(rootPath, ".mkv")
		webmFiles := getFilesOfType(rootPath, ".webm")
		videoFiles = append(videoFiles, mkvFiles...)
		videoFiles = append(videoFiles, webmFiles...)

		videoForm := widget.NewForm()

		videoForm.Append("video file", widget.NewSelect(videoFiles, updateEndEntry))
		beginEntry := widget.NewEntry()
		beginEntry.SetText("0:0")
		videoForm.Append("begin (mm:ss)", beginEntry)
		videoForm.Append("end (mm:ss)", endEntry)
		callBack := func(ok bool) {
			if ok {
				inputs := getFormInputs(videoForm.Items)

				// "video file" is compulsory
				if inputs["video file"] == "" {
					return
				}
				// complete the paths
				for k, v := range inputs {
					if strings.Contains(k, "file") && v != "" {
						inputs[k] = filepath.Join(rootPath, v)
					}
				}
				inputs["kind"] = "video"

				instructions = append(instructions, inputs)
				updateInstructionsBox()
				updateBottomBar()
			}
		}

		dialog.ShowCustomConfirm("Add Video Configuration", "Add", "Close", videoForm, callBack, myWindow)
	})

	topBar.Add(addImageBtn)
	topBar.Add(addVideoBtn)

	topBar.Add(aboutBtn)
	h1 := widget.NewLabelWithStyle("Instructions", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	windowBox := container.NewBorder(container.NewVBox(container.NewCenter(topBar), widget.NewSeparator(), h1),
		container.NewCenter(bottomBar), nil, nil, container.NewScroll(instructionsBox),
	)
	myWindow.SetContent(windowBox)
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}

func getFormInputs(content []*widget.FormItem) map[string]string {
	inputs := make(map[string]string)
	for _, formItem := range content {
		entryWidget, ok := formItem.Widget.(*widget.Entry)
		if ok {
			inputs[formItem.Text] = entryWidget.Text
			continue
		}

		selectWidget, ok := formItem.Widget.(*widget.Select)
		if ok {
			inputs[formItem.Text] = selectWidget.Selected
		}
	}

	return inputs
}

func getFilesOfType(rootPath, ext string) []string {
	dirFIs, err := os.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}
	files := make([]string, 0)
	for _, dirFI := range dirFIs {
		if !dirFI.IsDir() && !strings.HasPrefix(dirFI.Name(), ".") && strings.HasSuffix(dirFI.Name(), ext) {
			files = append(files, dirFI.Name())
		}

		if dirFI.IsDir() && !strings.HasPrefix(dirFI.Name(), ".") {
			innerDirFIs, _ := os.ReadDir(filepath.Join(rootPath, dirFI.Name()))
			innerFiles := make([]string, 0)

			for _, innerDirFI := range innerDirFIs {
				if !innerDirFI.IsDir() && !strings.HasPrefix(innerDirFI.Name(), ".") && strings.HasSuffix(innerDirFI.Name(), ext) {
					innerFiles = append(innerFiles, filepath.Join(dirFI.Name(), innerDirFI.Name()))
				}
			}

			if len(innerFiles) > 0 {
				files = append(files, innerFiles...)
			}
		}

	}

	return files
}
