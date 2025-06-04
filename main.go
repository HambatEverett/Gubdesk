package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	g "github.com/AllenDang/giu"
)

//go:embed images/gubby.png
var gubbyBytes []byte

var (
	// define variables
	affection int32 = 0
	open      bool  = false

	// define rgba image
	rgba *image.RGBA
)

func loop() {
	disp := fmt.Sprintf("Gubby Love: %d", affection)
	g.SingleWindow().Layout(
		g.Align(g.AlignCenter).To(
			g.Label("Gubby Friend!"),
			g.ImageWithRgba(rgba).OnClick(func() {
				affection = affection + 1
			}).Size(300, 300),
		),
		g.Label(disp),
		g.Button("Toys").Size(100, 50).OnClick(func() {
			open = true
			fmt.Printf("%v\n", open)
		}),
	)
}

func main() {
	// Load embedded image
	img, _, _ := image.Decode(bytes.NewReader(gubbyBytes))
	rgba = g.ImageToRgba(img)

	wnd := g.NewMasterWindow("My Friend Gubby", 500, 500, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
