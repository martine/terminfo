package terminfo

import (
	"encoding/binary"
	"io"
	"os"
	"strings"
)

// Use types wrapping int so that we don't accidentally look up a bool
// attribute in the numbers list.
type BoolAttr int
type NumberAttr int
type StringAttr int

const (
	AutoLeftMargin BoolAttr = iota
	AutoRightMargin
	NoEscCtlc
	CeolStandoutGlitch
	EatNewlineGlitch
	EraseOverstrike
	GenericType
	HardCopy
	HasMetaKey
	HasStatusLine
	InsertNullGlitch
	MemoryAbove
	MemoryBelow
	MoveInsertMode
	MoveStandoutMode
	OverStrike
	StatusLineEscOk
	DestTabsMagicSmso
	TildeGlitch
	TransparentUnderline
	XonXoff
	NeedsXonXoff
	PrtrSilent
	HardCursor
	NonRevRmcup
	NoPadChar
	NonDestScrollRegion
	CanChange
	BackColorErase
	HueLightnessSaturation
	ColAddrGlitch
	CrCancelsMicroMode
	HasPrintWheel
	RowAddrGlitch
	SemiAutoRightMargin
	CpiChangesRes
	LpiChangesRes  // 36

	// Internal caps.
	BackspacesWithBs
	CrtNoScrolling
	NoCorrectlyWorkingCr
	GnuHasMetaKey
	LinefeedIsNewline
)
var BoolAttrNames = [...]string {
	"AutoLeftMargin",
	"AutoRightMargin",
	"NoEscCtlc",
	"CeolStandoutGlitch",
	"EatNewlineGlitch",
	"EraseOverstrike",
	"GenericType",
	"HardCopy",
	"HasMetaKey",
	"HasStatusLine",
	"InsertNullGlitch",
	"MemoryAbove",
	"MemoryBelow",
	"MoveInsertMode",
	"MoveStandoutMode",
	"OverStrike",
	"StatusLineEscOk",
	"DestTabsMagicSmso",
	"TildeGlitch",
	"TransparentUnderline",
	"XonXoff",
	"NeedsXonXoff",
	"PrtrSilent",
	"HardCursor",
	"NonRevRmcup",
	"NoPadChar",
	"NonDestScrollRegion",
	"CanChange",
	"BackColorErase",
	"HueLightnessSaturation",
	"ColAddrGlitch",
	"CrCancelsMicroMode",
	"HasPrintWheel",
	"RowAddrGlitch",
	"SemiAutoRightMargin",
	"CpiChangesRes",
	"LpiChangesRes",

	// Internal caps.
	"BackspacesWithBs",
	"CrtNoScrolling",
	"NoCorrectlyWorkingCr",
	"GnuHasMetaKey",
	"LinefeedIsNewline",
}

const (
	Columns NumberAttr = iota
	InitTabs
	Lines
	LinesOfMemory
	MagicCookieGlitch
	PaddingBaudRate
	VirtualTerminal
	WidthStatusLine
	NumLabels
	LabelHeight
	LabelWidth
	MaxAttributes
	MaximumWindows
	MaxColors
	MaxPairs
	NoColorVideo
	BufferCapacity
	DotVertSpacing
	DotHorzSpacing
	MaxMicroAddress
	MaxMicroJump
	MicroColSize
	MicroLineSize
	NumberOfPins
	OutputResChar
	OutputResLine
	OutputResHorzInch
	OutputResVertInch
	PrintRate
	WideCharSize
	Buttons
	BitImageEntwining
	BitImageType

	// Internal caps.
	MagicCookieGlitchUl
	CarriageReturnDelay
	NewLineDelay
	BackspaceDelay
	HorizontalTabDelay
	NumberOfFunctionKeys
)

var NumberAttrNames = [...]string {
	"Columns",
	"InitTabs",
	"Lines",
	"LinesOfMemory",
	"MagicCookieGlitch",
	"PaddingBaudRate",
	"VirtualTerminal",
	"WidthStatusLine",
	"NumLabels",
	"LabelHeight",
	"LabelWidth",
	"MaxAttributes",
	"MaximumWindows",
	"MaxColors",
	"MaxPairs",
	"NoColorVideo",
	"BufferCapacity",
	"DotVertSpacing",
	"DotHorzSpacing",
	"MaxMicroAddress",
	"MaxMicroJump",
	"MicroColSize",
	"MicroLineSize",
	"NumberOfPins",
	"OutputResChar",
	"OutputResLine",
	"OutputResHorzInch",
	"OutputResVertInch",
	"PrintRate",
	"WideCharSize",
	"Buttons",
	"BitImageEntwining",
	"BitImageType",  // 32

	// Internal caps.
	"MagicCookieGlitchUl",
	"CarriageReturnDelay",
	"NewLineDelay",
	"BackspaceDelay",
	"HorizontalTabDelay",
	"NumberOfFunctionKeys",
}

const (
	BackTab StringAttr = iota
	Bell
	CarriageReturn
	ChangeScrollRegion
	ClearAllTabs
	ClearScreen
	ClrEol
	ClrEos
	ColumnAddress
	CommandCharacter
	CursorAddress
	CursorDown
	CursorHome
	CursorInvisible
	CursorLeft
	CursorMemAddress
	CursorNormal
	CursorRight
	CursorToLl
	CursorUp
	CursorVisible
	DeleteCharacter
	DeleteLine
	DisStatusLine
	DownHalfLine
	EnterAltCharsetMode
	EnterBlinkMode
	EnterBoldMode
	EnterCaMode
	EnterDeleteMode
	EnterDimMode
	EnterInsertMode
	EnterSecureMode
	EnterProtectedMode
	EnterReverseMode
	EnterStandoutMode
	EnterUnderlineMode
	EraseChars
	ExitAltCharsetMode
	ExitAttributeMode
	ExitCaMode
	ExitDeleteMode
	ExitInsertMode
	ExitStandoutMode
	ExitUnderlineMode
	FlashScreen
	FormFeed
	FromStatusLine
	Init1string
	Init2string
	Init3string
	InitFile
	InsertCharacter
	InsertLine
	InsertPadding
	KeyBackspace
	KeyCatab
	KeyClear
	KeyCtab
	KeyDc
	KeyDl
	KeyDown
	KeyEic
	KeyEol
	KeyEos
	KeyF0
	KeyF1
	KeyF10
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyHome
	KeyIc
	KeyIl
	KeyLeft
	KeyLl
	KeyNpage
	KeyPpage
	KeyRight
	KeySf
	KeySr
	KeyStab
	KeyUp
	KeypadLocal
	KeypadXmit
	LabF0
	LabF1
	LabF10
	LabF2
	LabF3
	LabF4
	LabF5
	LabF6
	LabF7
	LabF8
	LabF9
	MetaOff
	MetaOn
	Newline
	PadChar
	ParmDch
	ParmDeleteLine
	ParmDownCursor
	ParmIch
	ParmIndex
	ParmInsertLine
	ParmLeftCursor
	ParmRightCursor
	ParmRindex
	ParmUpCursor
	PkeyKey
	PkeyLocal
	PkeyXmit
	PrintScreen
	PrtrOff
	PrtrOn
	RepeatChar
	Reset1string
	Reset2string
	Reset3string
	ResetFile
	RestoreCursor
	RowAddress
	SaveCursor
	ScrollForward
	ScrollReverse
	SetAttributes
	SetTab
	SetWindow
	Tab
	ToStatusLine
	UnderlineChar
	UpHalfLine
	InitProg
	KeyA1
	KeyA3
	KeyB2
	KeyC1
	KeyC3
	PrtrNon
	CharPadding
	AcsChars
	PlabNorm
	KeyBtab
	EnterXonMode
	ExitXonMode
	EnterAmMode
	ExitAmMode
	XonCharacter
	XoffCharacter
	EnaAcs
	LabelOn
	LabelOff
	KeyBeg
	KeyCancel
	KeyClose
	KeyCommand
	KeyCopy
	KeyCreate
	KeyEnd
	KeyEnter
	KeyExit
	KeyFind
	KeyHelp
	KeyMark
	KeyMessage
	KeyMove
	KeyNext
	KeyOpen
	KeyOptions
	KeyPrevious
	KeyPrint
	KeyRedo
	KeyReference
	KeyRefresh
	KeyReplace
	KeyRestart
	KeyResume
	KeySave
	KeySuspend
	KeyUndo
	KeySbeg
	KeyScancel
	KeyScommand
	KeyScopy
	KeyScreate
	KeySdc
	KeySdl
	KeySelect
	KeySend
	KeySeol
	KeySexit
	KeySfind
	KeyShelp
	KeyShome
	KeySic
	KeySleft
	KeySmessage
	KeySmove
	KeySnext
	KeySoptions
	KeySprevious
	KeySprint
	KeySredo
	KeySreplace
	KeySright
	KeySrsume
	KeySsave
	KeySsuspend
	KeySundo
	ReqForInput
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyF26
	KeyF27
	KeyF28
	KeyF29
	KeyF30
	KeyF31
	KeyF32
	KeyF33
	KeyF34
	KeyF35
	KeyF36
	KeyF37
	KeyF38
	KeyF39
	KeyF40
	KeyF41
	KeyF42
	KeyF43
	KeyF44
	KeyF45
	KeyF46
	KeyF47
	KeyF48
	KeyF49
	KeyF50
	KeyF51
	KeyF52
	KeyF53
	KeyF54
	KeyF55
	KeyF56
	KeyF57
	KeyF58
	KeyF59
	KeyF60
	KeyF61
	KeyF62
	KeyF63
	ClrBol
	ClearMargins
	SetLeftMargin
	SetRightMargin
	LabelFormat
	SetClock
	DisplayClock
	RemoveClock
	CreateWindow
	GotoWindow
	Hangup
	DialPhone
	QuickDial
	Tone
	Pulse
	FlashHook
	FixedPause
	WaitTone
	User0
	User1
	User2
	User3
	User4
	User5
	User6
	User7
	User8
	User9
	OrigPair
	OrigColors
	InitializeColor
	InitializePair
	SetColorPair
	SetForeground
	SetBackground
	ChangeCharPitch
	ChangeLinePitch
	ChangeResHorz
	ChangeResVert
	DefineChar
	EnterDoublewideMode
	EnterDraftQuality
	EnterItalicsMode
	EnterLeftwardMode
	EnterMicroMode
	EnterNearLetterQuality
	EnterNormalQuality
	EnterShadowMode
	EnterSubscriptMode
	EnterSuperscriptMode
	EnterUpwardMode
	ExitDoublewideMode
	ExitItalicsMode
	ExitLeftwardMode
	ExitMicroMode
	ExitShadowMode
	ExitSubscriptMode
	ExitSuperscriptMode
	ExitUpwardMode
	MicroColumnAddress
	MicroDown
	MicroLeft
	MicroRight
	MicroRowAddress
	MicroUp
	OrderOfPins
	ParmDownMicro
	ParmLeftMicro
	ParmRightMicro
	ParmUpMicro
	SelectCharSet
	SetBottomMargin
	SetBottomMarginParm
	SetLeftMarginParm
	SetRightMarginParm
	SetTopMargin
	SetTopMarginParm
	StartBitImage
	StartCharSetDef
	StopBitImage
	StopCharSetDef
	SubscriptCharacters
	SuperscriptCharacters
	TheseCauseCr
	ZeroMotion
	CharSetNames
	KeyMouse
	MouseInfo
	ReqMousePos
	GetMouse
	SetAForeground
	SetABackground
	PkeyPlab
	DeviceType
	CodeSetInit
	Set0DesSeq
	Set1DesSeq
	Set2DesSeq
	Set3DesSeq
	SetLrMargin
	SetTbMargin
	BitImageRepeat
	BitImageNewline
	BitImageCarriageReturn
	ColorNames
	DefineBitImageRegion
	EndBitImageRegion
	SetColorBand
	SetPageLength
	DisplayPcChar
	EnterPcCharsetMode
	ExitPcCharsetMode
	EnterScancodeMode
	ExitScancodeMode
	PcTermOptions
	ScancodeEscape
	AltScancodeEsc
	EnterHorizontalHlMode
	EnterLeftHlMode
	EnterLowHlMode
	EnterRightHlMode
	EnterTopHlMode
	EnterVerticalHlMode
	SetAAttributes
	SetPglenInch  // 393

	// Internal caps.
	TermcapInit2
	TermcapReset
	LinefeedIfNotLf
	BackspaceIfNotBs
	OtherNonFunctionKeys
)

var StringAttrNames = [...]string {
	"BackTab",
	"Bell",
	"CarriageReturn",
	"ChangeScrollRegion",
	"ClearAllTabs",
	"ClearScreen",
	"ClrEol",
	"ClrEos",
	"ColumnAddress",
	"CommandCharacter",
	"CursorAddress",
	"CursorDown",
	"CursorHome",
	"CursorInvisible",
	"CursorLeft",
	"CursorMemAddress",
	"CursorNormal",
	"CursorRight",
	"CursorToLl",
	"CursorUp",
	"CursorVisible",
	"DeleteCharacter",
	"DeleteLine",
	"DisStatusLine",
	"DownHalfLine",
	"EnterAltCharsetMode",
	"EnterBlinkMode",
	"EnterBoldMode",
	"EnterCaMode",
	"EnterDeleteMode",
	"EnterDimMode",
	"EnterInsertMode",
	"EnterSecureMode",
	"EnterProtectedMode",
	"EnterReverseMode",
	"EnterStandoutMode",
	"EnterUnderlineMode",
	"EraseChars",
	"ExitAltCharsetMode",
	"ExitAttributeMode",
	"ExitCaMode",
	"ExitDeleteMode",
	"ExitInsertMode",
	"ExitStandoutMode",
	"ExitUnderlineMode",
	"FlashScreen",
	"FormFeed",
	"FromStatusLine",
	"Init1string",
	"Init2string",
	"Init3string",
	"InitFile",
	"InsertCharacter",
	"InsertLine",
	"InsertPadding",
	"KeyBackspace",
	"KeyCatab",
	"KeyClear",
	"KeyCtab",
	"KeyDc",
	"KeyDl",
	"KeyDown",
	"KeyEic",
	"KeyEol",
	"KeyEos",
	"KeyF0",
	"KeyF1",
	"KeyF10",
	"KeyF2",
	"KeyF3",
	"KeyF4",
	"KeyF5",
	"KeyF6",
	"KeyF7",
	"KeyF8",
	"KeyF9",
	"KeyHome",
	"KeyIc",
	"KeyIl",
	"KeyLeft",
	"KeyLl",
	"KeyNpage",
	"KeyPpage",
	"KeyRight",
	"KeySf",
	"KeySr",
	"KeyStab",
	"KeyUp",
	"KeypadLocal",
	"KeypadXmit",
	"LabF0",
	"LabF1",
	"LabF10",
	"LabF2",
	"LabF3",
	"LabF4",
	"LabF5",
	"LabF6",
	"LabF7",
	"LabF8",
	"LabF9",
	"MetaOff",
	"MetaOn",
	"Newline",
	"PadChar",
	"ParmDch",
	"ParmDeleteLine",
	"ParmDownCursor",
	"ParmIch",
	"ParmIndex",
	"ParmInsertLine",
	"ParmLeftCursor",
	"ParmRightCursor",
	"ParmRindex",
	"ParmUpCursor",
	"PkeyKey",
	"PkeyLocal",
	"PkeyXmit",
	"PrintScreen",
	"PrtrOff",
	"PrtrOn",
	"RepeatChar",
	"Reset1string",
	"Reset2string",
	"Reset3string",
	"ResetFile",
	"RestoreCursor",
	"RowAddress",
	"SaveCursor",
	"ScrollForward",
	"ScrollReverse",
	"SetAttributes",
	"SetTab",
	"SetWindow",
	"Tab",
	"ToStatusLine",
	"UnderlineChar",
	"UpHalfLine",
	"InitProg",
	"KeyA1",
	"KeyA3",
	"KeyB2",
	"KeyC1",
	"KeyC3",
	"PrtrNon",
	"CharPadding",
	"AcsChars",
	"PlabNorm",
	"KeyBtab",
	"EnterXonMode",
	"ExitXonMode",
	"EnterAmMode",
	"ExitAmMode",
	"XonCharacter",
	"XoffCharacter",
	"EnaAcs",
	"LabelOn",
	"LabelOff",
	"KeyBeg",
	"KeyCancel",
	"KeyClose",
	"KeyCommand",
	"KeyCopy",
	"KeyCreate",
	"KeyEnd",
	"KeyEnter",
	"KeyExit",
	"KeyFind",
	"KeyHelp",
	"KeyMark",
	"KeyMessage",
	"KeyMove",
	"KeyNext",
	"KeyOpen",
	"KeyOptions",
	"KeyPrevious",
	"KeyPrint",
	"KeyRedo",
	"KeyReference",
	"KeyRefresh",
	"KeyReplace",
	"KeyRestart",
	"KeyResume",
	"KeySave",
	"KeySuspend",
	"KeyUndo",
	"KeySbeg",
	"KeyScancel",
	"KeyScommand",
	"KeyScopy",
	"KeyScreate",
	"KeySdc",
	"KeySdl",
	"KeySelect",
	"KeySend",
	"KeySeol",
	"KeySexit",
	"KeySfind",
	"KeyShelp",
	"KeyShome",
	"KeySic",
	"KeySleft",
	"KeySmessage",
	"KeySmove",
	"KeySnext",
	"KeySoptions",
	"KeySprevious",
	"KeySprint",
	"KeySredo",
	"KeySreplace",
	"KeySright",
	"KeySrsume",
	"KeySsave",
	"KeySsuspend",
	"KeySundo",
	"ReqForInput",
	"KeyF11",
	"KeyF12",
	"KeyF13",
	"KeyF14",
	"KeyF15",
	"KeyF16",
	"KeyF17",
	"KeyF18",
	"KeyF19",
	"KeyF20",
	"KeyF21",
	"KeyF22",
	"KeyF23",
	"KeyF24",
	"KeyF25",
	"KeyF26",
	"KeyF27",
	"KeyF28",
	"KeyF29",
	"KeyF30",
	"KeyF31",
	"KeyF32",
	"KeyF33",
	"KeyF34",
	"KeyF35",
	"KeyF36",
	"KeyF37",
	"KeyF38",
	"KeyF39",
	"KeyF40",
	"KeyF41",
	"KeyF42",
	"KeyF43",
	"KeyF44",
	"KeyF45",
	"KeyF46",
	"KeyF47",
	"KeyF48",
	"KeyF49",
	"KeyF50",
	"KeyF51",
	"KeyF52",
	"KeyF53",
	"KeyF54",
	"KeyF55",
	"KeyF56",
	"KeyF57",
	"KeyF58",
	"KeyF59",
	"KeyF60",
	"KeyF61",
	"KeyF62",
	"KeyF63",
	"ClrBol",
	"ClearMargins",
	"SetLeftMargin",
	"SetRightMargin",
	"LabelFormat",
	"SetClock",
	"DisplayClock",
	"RemoveClock",
	"CreateWindow",
	"GotoWindow",
	"Hangup",
	"DialPhone",
	"QuickDial",
	"Tone",
	"Pulse",
	"FlashHook",
	"FixedPause",
	"WaitTone",
	"User0",
	"User1",
	"User2",
	"User3",
	"User4",
	"User5",
	"User6",
	"User7",
	"User8",
	"User9",
	"OrigPair",
	"OrigColors",
	"InitializeColor",
	"InitializePair",
	"SetColorPair",
	"SetForeground",
	"SetBackground",
	"ChangeCharPitch",
	"ChangeLinePitch",
	"ChangeResHorz",
	"ChangeResVert",
	"DefineChar",
	"EnterDoublewideMode",
	"EnterDraftQuality",
	"EnterItalicsMode",
	"EnterLeftwardMode",
	"EnterMicroMode",
	"EnterNearLetterQuality",
	"EnterNormalQuality",
	"EnterShadowMode",
	"EnterSubscriptMode",
	"EnterSuperscriptMode",
	"EnterUpwardMode",
	"ExitDoublewideMode",
	"ExitItalicsMode",
	"ExitLeftwardMode",
	"ExitMicroMode",
	"ExitShadowMode",
	"ExitSubscriptMode",
	"ExitSuperscriptMode",
	"ExitUpwardMode",
	"MicroColumnAddress",
	"MicroDown",
	"MicroLeft",
	"MicroRight",
	"MicroRowAddress",
	"MicroUp",
	"OrderOfPins",
	"ParmDownMicro",
	"ParmLeftMicro",
	"ParmRightMicro",
	"ParmUpMicro",
	"SelectCharSet",
	"SetBottomMargin",
	"SetBottomMarginParm",
	"SetLeftMarginParm",
	"SetRightMarginParm",
	"SetTopMargin",
	"SetTopMarginParm",
	"StartBitImage",
	"StartCharSetDef",
	"StopBitImage",
	"StopCharSetDef",
	"SubscriptCharacters",
	"SuperscriptCharacters",
	"TheseCauseCr",
	"ZeroMotion",
	"CharSetNames",
	"KeyMouse",
	"MouseInfo",
	"ReqMousePos",
	"GetMouse",
	"SetAForeground",
	"SetABackground",
	"PkeyPlab",
	"DeviceType",
	"CodeSetInit",
	"Set0DesSeq",
	"Set1DesSeq",
	"Set2DesSeq",
	"Set3DesSeq",
	"SetLrMargin",
	"SetTbMargin",
	"BitImageRepeat",
	"BitImageNewline",
	"BitImageCarriageReturn",
	"ColorNames",
	"DefineBitImageRegion",
	"EndBitImageRegion",
	"SetColorBand",
	"SetPageLength",
	"DisplayPcChar",
	"EnterPcCharsetMode",
	"ExitPcCharsetMode",
	"EnterScancodeMode",
	"ExitScancodeMode",
	"PcTermOptions",
	"ScancodeEscape",
	"AltScancodeEsc",
	"EnterHorizontalHlMode",
	"EnterLeftHlMode",
	"EnterLowHlMode",
	"EnterRightHlMode",
	"EnterTopHlMode",
	"EnterVerticalHlMode",
	"SetAAttributes",
	"SetPglenInch",

	// Internal caps.
	"TermcapInit2",
	"TermcapReset",
	"LinefeedIfNotLf",
	"BackspaceIfNotBs",
	"OtherNonFunctionKeys",
}

type Term struct {
	Names []string
	Bools map[BoolAttr]bool
	Numbers map[NumberAttr]int
	Strings map[StringAttr]string
}

func Parse(r io.Reader) (term *Term, err os.Error) {
	var header [6]uint16
	err = binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return
	}
	magic := header[0]
	namesLen := header[1]
	boolCount := header[2]
	intCount := header[3]
	stringCount := header[4]
	stringTableSize := header[5]

	if magic != 0432 {
		return nil, os.NewError("bad magic")
	}

	term = new(Term)

	buf := make([]byte, namesLen)
	_, err = io.ReadFull(r, buf)
	if err != nil {
                // I often do this, but you might want to return nil when returning with err != nil. You could also have a working Term with a different name so that plain "return" still returns a nil Term.
		return nil, err
	}
	// buf is NUL-terminated in the file; drop that.
	term.Names = strings.Split(string(buf[0:len(buf)-1]), "|", -1)

	bools := make([]byte, boolCount)
	err = binary.Read(r, binary.LittleEndian, &bools)
	if err != nil {
		return nil, err
	}
	term.Bools = make(map[BoolAttr]bool)
	for i, b := range bools {
		if b != 0 {
			term.Bools[BoolAttr(i)] = true
		}
	}

	// Skip padding byte if necessary.
	if boolCount % 2 != 0 {
                var scratch [1]byte
		n, err := r.Read(scratch[:])
		if err != nil {
			return nil, err
		}
		if n < 1 {
			return nil, io.ErrUnexpectedEOF
		}
	}

	nums := make([]uint16, intCount)
	err = binary.Read(r, binary.LittleEndian, &nums)
	if err != nil {
		return nil, err
	}
	term.Numbers = make(map[NumberAttr]int)
	for i, n := range(nums) {
		if n != 0xFFFF {
			term.Numbers[NumberAttr(i)] = int(n)
		}
	}

	stringOffsets := make([]uint16, stringCount)
	err = binary.Read(r, binary.LittleEndian, &stringOffsets)
	if err != nil {
		return nil, err
	}

	buf = make([]byte, stringTableSize)
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	if n < 1 {
		return nil, io.ErrUnexpectedEOF
	}

	term.Strings = make(map[StringAttr]string)
	for i, ofs := range stringOffsets {
		if ofs == 0xFFFF {
			continue
		}
		start := int(ofs)
		var end int
		for end = start; end < len(buf); end++ {
			if buf[end] == 0 {
				break
			}
		}
		term.Strings[StringAttr(i)] = string(buf[start:end])
	}

	return
}