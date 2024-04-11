package main

import (
	"context"
	"embed"
	"wails-svelte-tailwind-shadcn-template/pkg/app/hello"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	helloService := hello.NewHelloService()

	err := wails.Run(&options.App{
		Title:  "wails-svelte-tailwind-shadcn-template",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			helloService.Start(ctx)
		},
		Bind: []interface{}{
			helloService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
