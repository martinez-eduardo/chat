package main

import (
	"time"
	"strings"
	"fmt"
	"os/user"
	"io/ioutil"
	"net/http"
	"net/url"
	ui "github.com/VladimirMarkelov/clui"
)

var (
	urlchat = "http://ravr.webcindario.com/5_chat/consola.php"
	usuario string
	saladechat []byte
)

//Funcion para actualizar saladechat y para enviar mensajes
func Chatear(mensaje string) bool {
	linkparseado, conexionerror := url.Parse(urlchat)
	if conexionerror != nil {
		return false
	}
	parametros := url.Values{}
	parametros.Add("nombre", usuario)
	parametros.Add("mensaje", mensaje)
	linkparseado.RawQuery = parametros.Encode()
	respuesta, conexionerror := http.Get(linkparseado.String()) 
	if conexionerror != nil {
		return false
	}
	defer respuesta.Body.Close()
	if respuesta.StatusCode != 200 {
		return false
	}
	saladechat, conexionerror = ioutil.ReadAll(respuesta.Body)
	if conexionerror != nil {
		return false
	}
	respuesta.Body.Close()
	return true
}

//Funcion para crear interfaz grafica
func createView() *ui.TextDisplay {
	//Creamos ventana
	view := ui.AddWindow(0, 0, 30, 7, "Chat")
	view.SetPack(ui.Vertical)
	view.SetGaps(0, 1)
	view.SetPaddings(2, 2)
	view.SetMaximized(true)
	view.SetSizable(false)
	//MENSAJERIA
	txtchat := ui.CreateTextDisplay(view, 10, 10, 1)
	ui.ActivateControl(view, txtchat)
	//MENSAJE
	form1 := ui.CreateFrame(view, 1, 1, ui.BorderNone, ui.Fixed)
	form1.SetPack(ui.Horizontal)
	form1.SetGaps(1, 0)
	ui.CreateLabel(form1, ui.AutoSize, ui.AutoSize, "Mensaje:", ui.Fixed)
	txtusuario := ui.CreateEditField(form1, 8, "", 1)
	ui.ActivateControl(view, form1)
	//BOTONES
	form2 := ui.CreateFrame(view, 1, 1, ui.BorderNone, ui.Fixed)
	form2.SetPack(ui.Horizontal)
	btnenviar := ui.CreateButton(form2, ui.AutoSize, 2, "Enviar", ui.Fixed)
	btnlimpiar := ui.CreateButton(form2, ui.AutoSize, 2, "Limpiar", ui.Fixed)
	btnsalir := ui.CreateButton(form2, ui.AutoSize, 2, "Salir", ui.Fixed)
	ui.ActivateControl(view, form2)
	//Eventos de los botones
	btnlimpiar.OnClick(func(ev ui.Event) { txtusuario.SetTitle("") })
	btnsalir.OnClick(func(ev ui.Event) { ui.Stop() })
	btnenviar.OnClick(func(ev ui.Event) { if Chatear(txtusuario.Title()) { txtusuario.SetTitle("") } })
	return txtchat
}

//Funcion para aperturar la interfaz grafica y sus procesos en paralelo
func mainLoop() {
	user, errorfatal := user.Current()
	if errorfatal == nil{
		usuario = user.Username //Usuario del Sistema Operativo
	} else {
		usuario = "USUARIO" //En caso de no obtenerlo, Usuario generico
	}
	if !Chatear("") {
		fmt.Println("No se logro conectar a:",urlchat) //Si el sitio web del chat no funcion
	} else {
		ui.InitLibrary()
		defer ui.DeinitLibrary()
		b := createView()
		_ = b
		//PROCESO EN PARALELO PARA ACTUALIZAR LA MENSAJERIA CADA SEGUNDO
		go func() {
			for {
				if Chatear("") {
					lineas_nuevas := strings.Split(string(saladechat) , "\n")
					b.SetLineCount(10) //siempre mostrar solo 10 lineas
					b.OnDrawLine(func(ind int) string {
						switch ind {
							case 0, 2, 4, 6, 8, 10:
								return fmt.Sprintf(" ")
							case 1, 3, 5, 7, 9:
								variable := (ind - 1) - ((ind - 1)/2) //sucesion: 0,1,2,3,4 a partir de: 1,3,5,7,9
								if len(lineas_nuevas)< (variable+1) { return fmt.Sprintf(" ") } 
								campos := strings.Split(lineas_nuevas[variable], ";")
								if len(campos) == 3 { return fmt.Sprintf("%s: %s", campos[1], campos[2]) } else { return fmt.Sprintf(" ") }
							default:
								return fmt.Sprintf("FIN")
						}
					})
					ui.PutEvent(ui.Event{Type: ui.EventRedraw})
					time.Sleep(1 * time.Second) //ESPERAR UN SEGUNDO PARA VOLVER A ACTUALIZAR MENSAJES
				}  
			}
		}()
		ui.MainLoop()
	}
}

//Funcion Main
func main() {
	mainLoop()
}

