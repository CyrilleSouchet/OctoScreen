package ui

import (
	"github.com/mcuadros/go-octoprint"
)

var toolchangerPanelInstance *toolchangerPanel

type toolchangerPanel struct {
	CommonPanel
	activeTool int
}

func ToolchangerPanel(ui *UI, parent Panel) Panel {
	if toolchangerPanelInstance == nil {
		m := &toolchangerPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.panelH = 3
		m.initialize()

		toolchangerPanelInstance = m
	}

	return toolchangerPanelInstance
}

func (m *toolchangerPanel) initialize() {
	defer m.Initialize()

	if m.UI.Settings == nil || len(m.UI.Settings.ApplicationsStructure) == 0 {
		Logger.Info("Loading default menu")
	} else {
		var menuItems []octoprint.ApplicationsItem
		menuItems = m.UI.Settings.ApplicationsStructure
		buttons := MustGrid()
		buttons.SetRowHomogeneous(true)
		buttons.SetColumnHomogeneous(true)
		m.Grid().Attach(buttons, 3, 0, 2, 2)
		m.arrangeApplicationsItems(buttons, menuItems, 1)

	}
}
