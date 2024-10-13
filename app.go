package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	g "xabbo.b7c.io/goearth"
)

//go:embed all:frontend/dist
var assets embed.FS

var ext = g.NewExt(g.ExtInfo{
	Title:       "[JTD] G-Starter Kit [Wails_V2]",
	Description: "G Extension Starter Kit [Wails_V2]",
	Version:     "1.0.0",
	Author:      "JTD",
})

var app *App

func main() {
	// Create an instance of the app structure
	app := NewApp(ext, assets)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "[JTD] G-Starter Kit [Wails_V2]", // This is the title of the window
		Width:  850,
		Height: 1050,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// This method Shows the application.
func (a *App) ShowWindow() {
	// https://wails.io/docs/reference/runtime/intro/#show
	runtime.Show(a.ctx) // On Mac, this will bring the application back into the foreground. For Windows and Linux, this is currently the same as WindowShow.
}

func setupExt() {
	ext.Initialized(func(e g.InitArgs) {
		log.Printf("initialized (connected=%t)", e.Connected)
	})

	// This method is called when the green launch button is pressed in G-Earth
	ext.Activated(func() {
		log.Printf("activated")
		app.ShowWindow()
	})

	ext.Connected(func(e g.ConnectArgs) {
		log.Printf("connected (%s:%d)", e.Host, e.Port)
		log.Printf("client %s (%s)", e.Client.Identifier, e.Client.Version)
	})

	ext.Disconnected(func() {
		log.Printf("connection lost")
	})
}
