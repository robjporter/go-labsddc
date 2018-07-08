package app

import (
	"github.com/robjporter/go-xtools/xgraphics"
)

func (a *Application) displayBanner() {
	g := xgraphics.New()
	g.Display.ClearScreen()

	var lines []string
	lines = append(lines, "")
	lines = append(lines, "         .:.:. Cisco .:.:.")
	lines = append(lines, "Software Defined DataCenter Manager")
	lines = append(lines, "")

	g.Borders.SetBorderStyle("thick")
	g.Borders.SetContent(lines)
	g.Borders.SetSpacer(12)

	g.Borders.PrintBorder(true)
}
