package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
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
	toolsCount := m.defineToolsCount()

	m.Grid().Attach(m.createChangeToolButton(0), 1, 0, 1, 1)
	if toolsCount >= 2 {
		m.Grid().Attach(m.createChangeToolButton(1), 2, 0, 1, 1)
		if toolsCount >= 3 {
			m.Grid().Attach(m.createChangeToolButton(2), 3, 0, 1, 1)
			if toolsCount >= 4 {
				m.Grid().Attach(m.createChangeToolButton(3), 4, 0, 1, 1)
			}
		}
	}

	m.Grid().Attach(m.createHomeButton(), 1, 1, 1, 1)

	if m.UI.Settings == nil || len(m.UI.Settings.MenuStructure) == 0 {
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

	//m.Grid().Attach(m.createUSBOnButton(), 3, 1, 1, 1)
	//m.Grid().Attach(m.createUSBOffButton(), 4, 1, 1, 1)
	m.Grid().Attach(m.createZCalibrationButton(), 2, 2, 1, 1)
}

func (m *toolchangerPanel) defineToolsCount() int {
	c, err := (&octoprint.ConnectionRequest{}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return 0
	}

	profile, err := (&octoprint.PrinterProfilesRequest{Id: c.Current.PrinterProfile}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return 0
	}

	if profile.Extruder.SharedNozzle {
		return 1
	}

	return profile.Extruder.Count
}

func (m *toolchangerPanel) createZCalibrationButton() gtk.IWidget {
	b := MustButtonImageStyle("Z Offsets", "z-calibration.svg", "color2", func() {
		m.UI.Add(NozzleCalibrationPanel(m.UI, m))
	})

	return b
}

func (m *toolchangerPanel) createHomeButton() gtk.IWidget {
	return MustButtonImageStyle("Home XYZ", "home.svg", "color3", func() {
		cmd := &octoprint.CommandRequest{}
		cmd.Commands = []string{
			"G28 Z",
			"G28 X",
			"G28 Y",
		}
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
		}
	})
}

func (m *toolchangerPanel) createChangeToolButton(num int) gtk.IWidget {
	style := fmt.Sprintf("color%d", num+1)
	name := fmt.Sprintf("Tool%d", num)
	gcode := fmt.Sprintf("T%d", num)
	img := fmt.Sprintf("tool%d.svg", num)
	return MustButtonImageStyle(name, img, style, func() {
		m.command(gcode)
		Logger.Infof("Envoi de la commande %s", gcode)
	})
}

func (m *toolchangerPanel) createUSBOnButton() gtk.IWidget {
	return MustButtonImageStyle("USB On", "usb.svg", "color4", func() {
		cmd := &octoprint.CommandRequest{}
		cmd.Commands = []string{"OCTO1"}

		Logger.Info("USB On")
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})
}

func (m *toolchangerPanel) createUSBOffButton() gtk.IWidget {
	return MustButtonImageStyle("USB Off", "usb.svg", "color3", func() {
		cmd := &octoprint.CommandRequest{}
		//cmd.Commands = []string{"SET_PIN PIN=sol VALUE=0"}
		cmd.Commands = []string{"OCTO2"}

		Logger.Info("USB Off")
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})
}
