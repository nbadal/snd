//go:build ASTI
// +build ASTI

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/BigJk/snd/database"
	// "github.com/BigJk/snd/printing/preview"
	// "github.com/BigJk/snd/server"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Constants
const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
This is using the bootstrap and the bundler.`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	w *astilectron.Window
)

// var prev preview.AstiPreview

// This will change the starting routine so that an additional Electron window
// will open with the frontend in it.
func init() {
	startFunc = startElectron
	// serverOptions = append(serverOptions, server.WithPrinter(&prev))
}

func startElectron(db database.Database, debug bool) {
	// Start the S&D Backend in separate go-routine.
	go func() {
		startServer(db, debug)
	}()

	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:          Asset,
		AssetDir:       AssetDir,
		Debug:          debug,
		IgnoredSignals: []os.Signal{syscall.SIGURG},
		Logger:         l,
		RestoreAssets:  RestoreAssets,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
			ElectronSwitches: []string{
				"--disable-http-cache",
			},
		},
		Windows: []*bootstrap.Window{{
			Homepage: "http://127.0.0.1:7123/index.html",
			Options: &astilectron.WindowOptions{
				Center: astikit.BoolPtr(true),
				Height: astikit.IntPtr(920),
				Width:  astikit.IntPtr(1600),
				WebPreferences: &astilectron.WebPreferences{
					EnableRemoteModule: astikit.BoolPtr(true),
					WebviewTag: astikit.BoolPtr(true),
				},
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
		panic(err)
	}
}
