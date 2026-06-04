package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/thyamix/go-chess"
)

const GRIDSIZE int = 33

type Game struct {
	game     gochess.Game
	selected *[2]int
}

func NewGame() *Game {
	game, err := gochess.NewGame(gochess.BoardLayout(gochess.FENDefaultStart))
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		game: *game,
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if g.selected == nil {
			x, y := ebiten.CursorPosition()
			x /= GRIDSIZE
			y /= GRIDSIZE
			piece := g.game.State.Pieces[y*8+x]
			if piece != gochess.EMPTY {
				g.selected = &[2]int{x, y}
			}
		} else {
			g.selected = nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBoard(screen)
	for i, piece := range g.game.State.Pieces {
		x := i % 8
		y := i / 8
		if !(g.selected != nil && g.selected[0] == x && g.selected[1] == y) {
			DrawPiece(screen, x, y, false, piece)
		}
		if g.selected != nil {
			piece := g.game.State.Pieces[g.selected[1]*8+g.selected[0]]
			xpos, ypos := ebiten.CursorPosition()
			DrawPiece(screen, xpos, ypos, true, piece)
		}
	}
}

func DrawPiece(screen *ebiten.Image, x int, y int, selected bool, piece gochess.Piece) {
	var img *ebiten.Image
	switch piece.Type() {
	case gochess.PAWN:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_pawn.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_pawn.png")
		}
	case gochess.BISHOP:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_bishop.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_bishop.png")
		}
	case gochess.KING:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_king.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_king.png")
		}
	case gochess.KNIGHT:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_knight.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_knight.png")
		}
	case gochess.QUEEN:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_queen.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_queen.png")
		}
	case gochess.ROOK:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_rook.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_rook.png")
		}
	}
	if img == nil {
		return
	}
	geo := ebiten.GeoM{}
	geo.Scale(2, 2)
	if selected {
		geo.Translate(float64(x-GRIDSIZE/2), float64(y-GRIDSIZE/2))
	} else {
		geo.Translate(float64((GRIDSIZE*x)+2), float64((GRIDSIZE * y)))
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geo})
}

func DrawBoard(screen *ebiten.Image) {
	for x := range 8 {
		for y := range 8 {
			img := ebiten.NewImage(GRIDSIZE, GRIDSIZE)
			img.Fill(color.RGBA{255, 200, uint8((255) * ((y + x) % 2)), 255})
			geo := ebiten.GeoM{}
			geo.Translate(float64(GRIDSIZE*x), float64(GRIDSIZE*y))
			screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geo})
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 8 * GRIDSIZE, 8 * GRIDSIZE
}

func main() {

	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
