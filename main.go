
package main

import (
	"context"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	
	"stock-scribe-fifo/services"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the inventory service
	inventoryService := services.NewInventoryService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "StockScribe FIFO",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			inventoryService.SetContext(ctx)
		},
		Bind: []interface{}{
			inventoryService,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// Add these to ensure window is visible
			WindowIsResizable:    true,
			Theme:                windows.SystemDefault,
			IsZoomControlEnabled: false,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
