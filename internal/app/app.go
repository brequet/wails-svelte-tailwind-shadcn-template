package app

import (
	"context"
	"io/fs"
	"wails-svelte-tailwind-shadcn-template/internal/usecase"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func Run(frontendAssets fs.FS) {
	helloUseCase := usecase.NewHelloUseCase()

	err := wails.Run(&options.App{
		Title:  "wails-svelte-tailwind-shadcn-template",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: frontendAssets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			helloUseCase.Start(ctx)
		},
		Bind: []interface{}{
			helloUseCase,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
