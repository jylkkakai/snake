package main

import (
	"fmt"
	"image/color"
	"log"
	//"io"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	//"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)


var gameFont font.Face

func init()  {
	f, err  := opentype.Parse(goregular.TTF)
	//f, err  := opentype.ParseCollectionReaderAt(io.ByteReader("AtariClassic-gry3.ttf"))
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size: 32,
		DPI: 72,

	})
	
	gameFont = face
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	//img := ebiten.NewImage(10, 10)
	screen.Fill(color.Gray{100})
	vector.DrawFilledRect(screen, 20, 60, 600, 400, color.Black, false)
	vector.StrokeRect(screen, 19, 59, 602, 402, 2, color.Gray{150}, false)
	vector.DrawFilledRect(screen, 100 + 1, 100 + 1, 8, 8, color.White, false)
	vector.DrawFilledRect(screen, 110 + 1, 100 + 1, 8, 8, color.White, false)
	vector.DrawFilledCircle(screen, 120 + 5, 100 + 5, 6, color.White, false)
	vector.DrawFilledCircle(screen, 220 + 4, 200 + 4, 4, color.RGBA{0xff, 0, 0, 0xff}, false)
	text.Draw(screen, fmt.Sprintf("%03d", 20), gameFont, 560, 40, color.White)
	text.Draw(screen, "SNAKE", gameFont, 260, 40, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
