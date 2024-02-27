package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"slices"

	//"io"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	//"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

type coords struct {
	x int
	y int
}

const (
	scrWidth int = 640
	scrHeight int = 480

	gridWidth int = 60
	gridHeight int = 40
)

var (
	snake = []coords{
		{
			x: 10,
			y: 11,
		},
		{
			x: 10,
			y: 10,
		},
		{
			x: 10,
			y: 9,
		},
	}
	apple = coords{
		x: 30,
		y: 20,
	}

	tickCounter int = 0
	delay int = 30


	up = coords{x: 0, y: -1}
	down = coords{x: 0, y: 1}
	left = coords{x: -1, y: 0}
	right = coords{x: 1, y: 0}
	direction coords = down

	gameOn bool = true
	score int = 0

	key []ebiten.Key
)

func (g *Game) Update() error {

	if gameOn {
		key = inpututil.AppendPressedKeys(key[:0])
		log.Println(key)
		if len(key) > 0 {
			if key[0] == ebiten.KeyArrowUp {
				direction = up
			} else if key[0] == ebiten.KeyArrowDown {
				direction = down
			} else if key[0] ==  ebiten.KeyArrowLeft {
				direction = left
			} else if key[0] == ebiten.KeyArrowRight {
				direction = right
			}
		}
		if tickCounter >= delay {
			temp := []coords{{x: direction.x + snake[0].x, y: direction.y + snake[0].y}}
			snake = append(temp, snake...)
			if snake[0].x != apple.x || snake[0].y != apple.y {
				snake = snake[:len(snake) - 1]
			} else {
				score += 1
				for slices.Contains(snake, apple){
					apple = coords{x: rand.Intn(gridWidth), y: rand.Intn(gridHeight)}
					if delay > 6 {
						delay--
					}
				}
			}
			tickCounter = 0
		}
		tickCounter++
		if snake[0].y > gridHeight - 1 || snake[0].x > gridWidth - 1 || snake[0].y < 0 || snake[0].x < 0 {
			gameOn = false
		}
		log.Println(gameOn, snake[0])
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	//img := ebiten.NewImage(10, 10)
	screen.Fill(color.Gray{100})
	text.Draw(screen, fmt.Sprintf("%03d", score), gameFont, 560, 40, color.White) 
	text.Draw(screen, "SNAKE", gameFont, 260, 40, color.White)

	vector.DrawFilledRect(screen, 20, 60, 600, 400, color.Black, false)
	vector.StrokeRect(screen, 19, 59, 602, 402, 2, color.Gray{150}, false)

	for i, bodyPart := range snake {
		//log.Println(i,bodyPart.x*10 + 1, bodyPart.y*10 + 1, float32(bodyPart.x*10 + 1), float32(bodyPart.y*10 + 1))
		if i == 0 {
			vector.DrawFilledCircle(screen, float32(20 + bodyPart.x*10 + 5), float32(60 + bodyPart.y*10 + 5), 4, color.White, false)
		} else {
			vector.DrawFilledRect(screen, float32(20 + bodyPart.x*10 + 1), float32(60 + bodyPart.y*10 + 1), 8, 8, color.White, false)
		}
	}
	vector.DrawFilledCircle(screen, float32(20 + apple.x*10 + 4), float32(60 + apple.y*10 + 4), 4, color.RGBA{0xff, 0, 0, 0xff}, false)

	// vector.DrawFilledRect(screen, 20 + 1, 60 + 1, 8, 8, color.White, false)
	// vector.DrawFilledRect(screen, 620 + 1, 420 + 1, 8, 8, color.White, false)
	// vector.DrawFilledCircle(screen, 120 + 5, 100 + 5, 6, color.White, false)
	// vector.DrawFilledCircle(screen, 220 + 4, 200 + 4, 4, color.RGBA{0xff, 0, 0, 0xff}, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
