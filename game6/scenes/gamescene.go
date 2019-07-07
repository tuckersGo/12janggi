package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tuckersGo/12janggi/game6/global"
	"github.com/tuckersGo/12janggi/game6/scenemanager"
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

type GameScene struct {
	bgimg       *ebiten.Image
	gimulImges  [GimulMax]*ebiten.Image
	selectedImg *ebiten.Image
	board       [global.BoardWidth][global.BoardHeight]GimulType
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType
	gameover    bool
}

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

func (g *GameScene) Startup() {
	g.gameover = false
	g.currentTeam = TeamGreen

	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImges[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.gimulImges[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {
			g.board[i][j] = GimulNone
		}
	}

	g.board[0][0] = GimulGreenSang
	g.board[0][1] = GimulGreenWang
	g.board[0][2] = GimulGreenJang
	g.board[1][1] = GimulGreenJa

	g.board[3][0] = GimulRedSang
	g.board[3][1] = GimulRedWang
	g.board[3][2] = GimulRedJang
	g.board[2][1] = GimulRedJa
}

func (g *GameScene) Update(screen *ebiten.Image) error {

	screen.DrawImage(g.bgimg, nil)
	if g.gameover {
		return nil
	}

	// Input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/global.GridHeight, y/global.GridHeight

		if i >= 0 && i < global.GridWidth && j >= 0 && j < global.GridHeight {
			if !g.selected {
				if g.board[i][j] != GimulNone && g.currentTeam == GetTeamType(g.board[i][j]) {
					g.selected = true
					g.selectedX, g.selectedY = i, j
				}
			} else {
				if g.selectedX == i && g.selectedY == j {
					g.selected = false
				} else {
					// move
					g.moveGimul(g.selectedX, g.selectedY, i, j)
				}
			}
		}
	}

	// Draw gimuls
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {

			// The previous empty option struct
			opts := &ebiten.DrawImageOptions{}
			// Add the Translate effect to the option struct.
			opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*i), float64(global.GimulStartY+global.GridHeight*j))

			switch g.board[i][j] {
			case GimulGreenJa:
				screen.DrawImage(g.gimulImges[GimulGreenJa], opts)
			case GimulGreenJang:
				screen.DrawImage(g.gimulImges[GimulGreenJang], opts)
			case GimulGreenSang:
				screen.DrawImage(g.gimulImges[GimulGreenSang], opts)
			case GimulGreenWang:
				screen.DrawImage(g.gimulImges[GimulGreenWang], opts)
			case GimulRedJa:
				screen.DrawImage(g.gimulImges[GimulRedJa], opts)
			case GimulRedJang:
				screen.DrawImage(g.gimulImges[GimulRedJang], opts)
			case GimulRedSang:
				screen.DrawImage(g.gimulImges[GimulRedSang], opts)
			case GimulRedWang:
				screen.DrawImage(g.gimulImges[GimulRedWang], opts)
			}
		}
	}

	if g.selected {
		// The previous empty option struct
		opts := &ebiten.DrawImageOptions{}
		// Add the Translate effect to the option struct.
		opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*g.selectedX),
			float64(global.GimulStartY+global.GridHeight*g.selectedY))
		screen.DrawImage(g.selectedImg, opts)
	}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (g *GameScene) moveGimul(prevX, prevY, tarX, tarY int) {
	if g.isMovable(prevX, prevY, tarX, tarY) {
		g.OnDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulNone, g.board[prevX][prevY]
		g.selected = false
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}

func (g *GameScene) isMovable(prevX, prevY, tarX, tarY int) bool {
	if GetTeamType(g.board[prevX][prevY]) == GetTeamType(g.board[tarX][tarY]) {
		return false
	}
	switch g.board[prevX][prevY] {
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
func (g *GameScene) OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang ||
		gimulType == GimulRedWang {
		scenemanager.SetScene(&GameoverScene{})
	}
}
