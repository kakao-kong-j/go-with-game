package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type GimulType int

const (
	GimulTypeNone GimulType = -1 + iota
	GimulTypeGreenWang
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax

	BoardWidth  = 4
	BoardHeight = 3

	ScreenWidth  = 500
	ScreenHeight = 400

	GimulStartX = 20
	GimulStartY = 23
	GridWidth   = 116
	GridHeight  = 116
)

var (
	board      [BoardWidth][BoardHeight]GimulType
	bgimg      *ebiten.Image
	gimulImges [GimulTypeMax]*ebiten.Image
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))
			switch board[i][j] {
			case GimulTypeGreenJa:
				screen.DrawImage(gimulImges[GimulTypeGreenJa], opts)
			case GimulTypeGreenJang:
				screen.DrawImage(gimulImges[GimulTypeGreenJang], opts)
			case GimulTypeGreenWang:
				screen.DrawImage(gimulImges[GimulTypeGreenWang], opts)
			case GimulTypeGreenSang:
				screen.DrawImage(gimulImges[GimulTypeGreenSang], opts)
			case GimulTypeRedJa:
				screen.DrawImage(gimulImges[GimulTypeRedJa], opts)
			case GimulTypeRedJang:
				screen.DrawImage(gimulImges[GimulTypeRedJang], opts)
			case GimulTypeRedWang:
				screen.DrawImage(gimulImges[GimulTypeRedWang], opts)
			case GimulTypeRedSang:
				screen.DrawImage(gimulImges[GimulTypeRedSang], opts)
			}
		}
	}
	return nil
}

func main() {
	var err error

	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}
	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	board[3][0] = GimulTypeRedSang
	board[3][1] = GimulTypeRedWang
	board[3][2] = GimulTypeRedJang
	board[2][1] = GimulTypeRedJa

	err = ebiten.Run(update, ScreenWidth, ScreenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
