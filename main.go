package main

import (
	"embed"
	"wails-svelte-tailwind-shadcn-template/internal/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app.Run(assets)
}
