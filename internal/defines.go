package internal

import (
	"image"

	g143 "github.com/bankole7782/graphics143"
)

const (
	FPS       = 10
	FontSize  = 20
	AddImgBtn = 101
	AddVidBtn = 102
	OpenWDBtn = 103
	RenderBtn = 104
	OurSite   = 105

	VAI_SelectImg   = 21
	VAI_SelectAudio = 22
	VAI_DurInput    = 23
	VAI_AddBtn      = 24
	VAI_CloseBtn    = 25

	VAV_AddBtn     = 31
	VAV_CloseBtn   = 32
	VAV_PickVideo  = 33
	VAV_BeginInput = 34
	VAV_EndInput   = 35
	VAV_PickAudio  = 36
)

var (
	ObjCoords          map[int]g143.RectSpecs
	CurrentWindowFrame image.Image
	Instructions       []map[string]string

	// tmpFrame image.Image
	InChannel        chan bool
	ClearAfterRender bool

	VaiObjCoords   map[int]g143.RectSpecs
	VaiInputsStore map[string]string
	VaiEnteredText string

	VavObjCoords   map[int]g143.RectSpecs
	VavInputsStore map[string]string

	BeginInputEnteredTxt string = "0:00"
	EndInputEnteredTxt   string = "0:00"
	SelectedInput        int
)
