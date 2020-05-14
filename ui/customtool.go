package ui

import (
	"github.com/mcuadros/go-octoprint"
)

var customtoolPanelInstance *customtoolPanel

type customtoolPanel struct {
	CommonPanel
	activeTool int
}

func CustomToolPanel(ui *UI, parent Panel) Panel {
	if customtoolPanelInstance == nil {
		m := &customtoolPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.panelH = 2
		m.initialize()

		customtoolPanelInstance = m
	}

	return customtoolPanelInstance
}

func (m *customtoolPanel) initialize() {
	defer m.Initialize()

	if m.UI.Settings == nil || len(m.UI.Settings.ApplicationsStructure) == 0 {
		Logger.Info("Loading default menu")
	} else {
		var menuItems []octoprint.ApplicationsItem
		menuItems = m.UI.Settings.ApplicationsStructure
		/*buttons := MustGrid()
		buttons.SetRowHomogeneous(true)
		buttons.SetColumnHomogeneous(true)
		m.Grid().Attach(buttons, 3, 0, 2, 2)*/
		m.arrangeApplicationsItems(m.Grid(), menuItems, 4)

	}
}
