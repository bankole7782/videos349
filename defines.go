package main

import (
	"image"
	"time"

	g143 "github.com/bankole7782/graphics143"
)

const (
	FPS       = 24
	FontSize  = 20
	fontColor = "#444"
	PageSize  = 20

	AddImgBtn      = 101
	AddImgSoundBtn = 102
	AddVidBtn      = 103
	OpenWDBtn      = 104
	RenderBtn      = 105

	VAI_SelectImg  = 21
	VAI_DurInput   = 23
	VAI_AddBtn     = 24
	VAI_CloseBtn   = 25
	VAI_AudioBegin = 26

	VAIS_SelectImg       = 31
	VAIS_SelectAudio     = 32
	VAIS_AddBtn          = 33
	VAIS_CloseBtn        = 34
	VAIS_AudioBeginInput = 35
	VAIS_AudioEndInput   = 36

	VAV_AddBtn                = 41
	VAV_CloseBtn              = 42
	VAV_PickVideo             = 43
	VAV_BeginInput            = 44
	VAV_EndInput              = 45
	VAV_SpeedUpCheckbox       = 46
	VAV_BlackAndWhiteCheckbox = 47

	PROJ_NameInput  = 51
	PROJ_NewProject = 52

	ProgTitle = "videos349: a simple video editor for teachers"
)

var (
	ObjCoords          map[int]g143.Rect
	CurrentWindowFrame image.Image
	Instructions       []map[string]string
	ProjectName        string
	ToUpdateInstrNum   int // for updating an instruction
	IsUpdateDialog     bool
	CurrentPage        int

	ToInsertBefore       int // for inbetween dialogs
	IsInsertBeforeDialog bool

	// tmpFrame image.Image
	InChannel        chan bool
	ClearAfterRender bool

	// view add image
	VaiObjCoords             map[int]g143.Rect
	VaiInputsStore           map[string]string
	VAI_DurationEnteredTxt   string
	VAI_SelectedInput        int
	VAI_AudioBeginEnteredTxt string = "0:00"

	// view add image + sound
	VaisObjCoords   map[int]g143.Rect
	VaisInputsStore map[string]string

	VaisBeginInputEnteredTxt string = "0:00"
	VaisEndInputEnteredTxt   string = "0:00"
	VAIS_SelectedInput       int

	// view add video
	VavObjCoords   map[int]g143.Rect
	VavInputsStore map[string]string

	BeginInputEnteredTxt              string = "0:00"
	EndInputEnteredTxt                string = "0:00"
	VAV_SelectedInput                 int
	VAV_SpeedUpCheckboxSelected       bool
	VAV_BlackAndWhiteCheckboxSelected bool

	// view projects
	ProjObjCoords map[int]g143.Rect

	NameInputEnteredTxt string

	cursorEventsCount = 0
)

type ToSortProject struct {
	Name    string
	ModTime time.Time
}
