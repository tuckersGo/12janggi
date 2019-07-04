package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tuckersGo/12janggi/game5/scenemanager"
	"github.com/tuckersGo/12janggi/game5/scenes"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 362
	BoardWidth   = 4
	BoardHeight  = 3
	GimulStartX  = 20
	GimulStartY  = 23
	GridWidth    = 116
	GridHeight   = 116
)

type GimulType int

const (
	GimulNone GimulType = iota - 1
	GimulGreenJa
	GimulGreenJang
	GimulGreenSang
	GimulGreenWang
	GimulRedJa
	GimulRedJang
	GimulRedSang
	GimulRedWang
	GimulMax
)

type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

var (
	bgimg       *ebiten.Image
	gimulImges  [GimulMax]*ebiten.Image
	selectedImg *ebiten.Image
	board       [BoardWidth][BoardHeight]GimulType
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType = TeamGreen
	gameover    bool
)

func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulGreenJa ||
		gimulType == GimulGreenJang ||
		gimulType == GimulGreenSang ||
		gimulType == GimulGreenWang {
		return TeamGreen
	}
	if gimulType == GimulRedJa ||
		gimulType == GimulRedJang ||
		gimulType == GimulRedSang ||
		gimulType == GimulRedWang {
		return TeamRed
	}
	return TeamNone
}

func update(screen *ebiten.Image) error {

	screen.DrawImage(bgimg, nil)
	if gameover {
		return nil
	}

	// Input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridHeight, y/GridHeight

		if i >= 0 && i < GridWidth && j >= 0 && j < GridHeight {
			if !selected {
				if board[i][j] != GimulNone && currentTeam == GetTeamType(board[i][j]) {
					selected = true
					selectedX, selectedY = i, j
				}
			} else {
				if selectedX == i && selectedY == j {
					selected = false
				} else {
					// move
					moveGimul(selectedX, selectedY, i, j)
				}
			}
		}
	}

	// Draw gimuls
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			// The previous empty option struct
			opts := &ebiten.DrawImageOptions{}
			// Add the Translate effect to the option struct.
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))

			switch board[i][j] {
			case GimulGreenJa:
				screen.DrawImage(gimulImges[GimulGreenJa], opts)
			case GimulGreenJang:
				screen.DrawImage(gimulImges[GimulGreenJang], opts)
			case GimulGreenSang:
				screen.DrawImage(gimulImges[GimulGreenSang], opts)
			case GimulGreenWang:
				screen.DrawImage(gimulImges[GimulGreenWang], opts)
			case GimulRedJa:
				screen.DrawImage(gimulImges[GimulRedJa], opts)
			case GimulRedJang:
				screen.DrawImage(gimulImges[GimulRedJang], opts)
			case GimulRedSang:
				screen.DrawImage(gimulImges[GimulRedSang], opts)
			case GimulRedWang:
				screen.DrawImage(gimulImges[GimulRedWang], opts)
			}
		}
	}

	if selected {
		// The previous empty option struct
		opts := &ebiten.DrawImageOptions{}
		// Add the Translate effect to the option struct.
		opts.GeoM.Translate(float64(GimulStartX+GridWidth*selectedX), float64(GimulStartY+GridHeight*selectedY))
		screen.DrawImage(selectedImg, opts)
	}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func moveGimul(prevX, prevY, tarX, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		OnDie(board[tarX][tarY])
		board[prevX][prevY], board[tarX][tarY] = GimulNone, board[prevX][prevY]
		selected = false
		if currentTeam == TeamGreen {
			currentTeam = TeamRed
		} else {
			currentTeam = TeamGreen
		}
	}
}

func isMovable(prevX, prevY, tarX, tarY int) bool {
	if GetTeamType(board[prevX][prevY]) == GetTeamType(board[tarX][tarY]) {
		return false
	}
	switch board[prevX][prevY] {
	case GimulGreenJa:
		return prevY == tarY && prevX+1 == tarX
	case GimulRedJa:
		return prevY == tarY && prevX-1 == tarX
	case GimulGreenJang, GimulRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulGreenSang, GimulRedSang:
		return (abs(prevX-tarX) == 1 && abs(prevY-tarY) == 1)
	case GimulGreenWang, GimulRedWang:
		return (abs(prevX-tarX) == 1 || abs(prevY-tarY) == 1)
	}
	return true
}

// OnDie calls when gimul is died
func OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang ||
		gimulType == GimulRedWang {
		gameover = true
	}
}

func main() {
	var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	gimulImges[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulNone
		}
	}

	board[0][0] = GimulGreenSang
	board[0][1] = GimulGreenWang
	board[0][2] = GimulGreenJang
	board[1][1] = GimulGreenJa

	board[3][0] = GimulRedSang
	board[3][1] = GimulRedWang
	board[3][2] = GimulRedJang
	board[2][1] = GimulRedJa

	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, ScreenWidth, ScreenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
