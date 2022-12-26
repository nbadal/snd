//go:build WEBVIEW
// +build WEBVIEW

package main

import (
	"github.com/BigJk/snd/database"

	"github.com/webview/webview"
)

func init() {
	startFunc = startWebview
}

var w webview.WebView

func startWebview(db database.Database, debug bool) {
	// Start the S&D Backend in separate go-routine.
	go func() {
		startServer(db, debug)
	}()

	w = webview.New(debug)
	defer w.Destroy()
	// w.SetTitle("WebView Example with Bindings")
	w.SetSize(800, 600, webview.HintNone)

	// Load the start page, index.html
	w.Navigate("http://localhost:7123/index.html")

	// Run the app.
	w.Run()
}
