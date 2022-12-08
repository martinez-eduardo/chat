package main

import (
	"fmt"
	ui "github.com/VladimirMarkelov/clui"
)

func createView() {
	view := ui.AddWindow(0, 0, 30, 7, "Chat")
	view.SetPack(ui.Vertical)
	view.SetGaps(0, 1)
	view.SetPaddings(2, 2)
	view.SetMaximized(true)
	view.SetSizable(false)

	/*MENSAJERIA*/
	btnDlgs := ui.CreateTextDisplay(view, 10, 10, 1)
	ui.ActivateControl(view, btnDlgs)

	/*MENSAJE PROMPT*/
	frmBtnse := ui.CreateFrame(view, 1, 1, ui.BorderNone, ui.Fixed)
	frmBtnse.SetPack(ui.Horizontal)
	frmBtnse.SetGaps(1, 0)
	ui.CreateLabel(frmBtnse, ui.AutoSize, ui.AutoSize, "Mensaje:", ui.Fixed)
	edUser := ui.CreateEditField(frmBtnse, 8, "", 1)
	ui.ActivateControl(view, frmBtnse)

	/*BOTONES*/
	frmBtns := ui.CreateFrame(view, 1, 1, ui.BorderNone, ui.Fixed)
	frmBtns.SetPack(ui.Horizontal)
	btnDlg := ui.CreateButton(frmBtns, ui.AutoSize, 2, "Enviar", ui.Fixed)
	ads := ui.CreateButton(frmBtns, ui.AutoSize, 2, "Limpiar", ui.Fixed)
	adss := ui.CreateButton(frmBtns, ui.AutoSize, 2, "Salir", ui.Fixed)
	ui.ActivateControl(view, frmBtns)
	
	

	btnDlg.OnClick(func(ev ui.Event) {
		Texto := edUser.Title()
		btnDlgs.SetLineCount(10) 
		btnDlgs.OnDrawLine(func(ind int) string {
			switch ind {
				case 0:
					return fmt.Sprintf(" ")
				case 1:
					return fmt.Sprintf("Ricardo: %s", Texto)
				case 2:
					return fmt.Sprintf(" ")
				case 3:
					return fmt.Sprintf("Veronica: mensaje")

				case 4:
					return fmt.Sprintf(" ")
				case 5:
					return fmt.Sprintf("Desconocido: mensaje")
				case 6:
					return fmt.Sprintf(" ")
				case 7:
					return fmt.Sprintf("Desconocido: mensaje")
				case 8:
					return fmt.Sprintf(" ")
				case 9:
					return fmt.Sprintf("Veronica: Hola")

				default:
					return fmt.Sprintf("FIN")
			}
		})
	})

	ads.OnClick(func(ev ui.Event) {
		edUser.SetTitle("")
		/*btnDlgs.SetLineCount(2) 
		btnDlgs.OnDrawLine(func(ind int) string {
			switch ind {
				case 0:
					return fmt.Sprintf("Ricardo:",)
				case 1:
					return fmt.Sprintf("Veronica:")
				default:
					return fmt.Sprintf("FIN")
			}
		})*/
	})

	adss.OnClick(func(ev ui.Event) {
			go ui.Stop()
		})
	}

func mainLoop() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	createView()
	ui.MainLoop()
}

func main() {
	mainLoop()
}
