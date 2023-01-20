package style

import (
	"text/template"

	"github.com/gookit/color"
)

var (
	FgWhite        = FromBasicFg(color.FgWhite)
	FgLightWhite   = FromBasicFg(color.FgLightWhite)
	FgBlack        = FromBasicFg(color.FgBlack)
	FgBlackLighter = FromBasicFg(color.FgBlack.Light())
	FgCyan         = FromBasicFg(color.FgCyan)
	FgRed          = FromHexFg("#FF7733")
	FgGreen        = FromHexFg("#B8CC52")
	FgBlue         = FromHexFg("#36A3D9")
	FgYellow       = FromHexFg("#E6B673")
	FgMagenta      = FromBasicFg(color.FgMagenta)
	FgDefault      = FromBasicFg(color.FgDefault)

	BgWhite   = FromBasicBg(color.BgWhite)
	BgBlack   = FromBasicBg(color.BgBlack)
	BgRed     = FromHexBg("#FF7733")
	BgGreen   = FromHexBg("#B8CC52")
	BgYellow  = FromHexBg("#E6B673")
	BgBlue = FromHexBg("#253340")
	BgMagenta = FromBasicBg(color.BgMagenta)
	BgCyan    = FromBasicBg(color.BgCyan)
	BgDefault = FromBasicBg(color.BgDefault)

	// will not print any colour escape codes, including the reset escape code
	Nothing = New()

	AttrUnderline = New().SetUnderline()
	AttrBold      = New().SetBold()

	ColorMap = map[string]struct {
		Foreground TextStyle
		Background TextStyle
	}{
		"default": {FgWhite, BgBlack},
		"black":   {FgBlack, BgBlack},
		"red":     {FgRed, BgRed},
		"green":   {FgGreen, BgGreen},
		"yellow":  {FgYellow, BgYellow},
		"blue":    {FgBlue, BgBlue},
		"magenta": {FgMagenta, BgMagenta},
		"cyan":    {FgCyan, BgCyan},
		"white":   {FgWhite, BgWhite},
	}
)

func FromBasicFg(fg color.Color) TextStyle {
	return New().SetFg(NewBasicColor(fg))
}

func FromRGBFg(fg color.RGBColor) TextStyle {
	return New().SetFg(NewRGBColor(fg))
}

func FromHexFg(fg string) TextStyle {
	return New().SetFg(NewRGBColor(color.Hex(fg, false)))
}

func FromBasicBg(bg color.Color) TextStyle {
	return New().SetBg(NewBasicColor(bg))
}

func FromRGBBg(bg color.RGBColor) TextStyle {
	return New().SetBg(NewRGBColor(bg))
}

func FromHexBg(bg string) TextStyle {
	return New().SetBg(NewRGBColor(color.Hex(bg, false)))
}

func TemplateFuncMapAddColors(m template.FuncMap) template.FuncMap {
	for k, v := range ColorMap {
		m[k] = v.Foreground.Sprint
	}
	m["underline"] = color.OpUnderscore.Sprint
	m["bold"] = color.OpBold.Sprint
	return m
}
