package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	chess "github.com/thyamix/go-chess"
)

const GRIDSIZE int = 33

type Game struct {
	board    chess.Game
	selected *[2]int
}

func NewGame() *Game {
	return &Game{
		board: chess.NewGame(chess.NewBoard()),
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if g.selected == nil {
			x, y := ebiten.CursorPosition()
			x = (8*GRIDSIZE - x) / GRIDSIZE
			y = (8*GRIDSIZE - y) / GRIDSIZE
			piece, _ := g.board.GetPiece(x, y)
			if piece.Type() != chess.EMPTY {
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
	for x := range 8 {
		for y := range 8 {
			piece, err := g.board.GetPiece(x, y)
			if err != nil || piece.Type() == chess.EMPTY {
				continue
			}
			if !(g.selected != nil && g.selected[0] == x && g.selected[1] == y) {
				DrawPiece(screen, x, y, false, piece)
			}
		}
	}
	if g.selected != nil {
		piece, _ := g.board.GetPiece(g.selected[0], g.selected[1])
		xpos, ypos := ebiten.CursorPosition()
		DrawPiece(screen, xpos, ypos, true, piece)
	}
}

func DrawPiece(screen *ebiten.Image, x int, y int, selected bool, piece chess.Piece) {
	var img *ebiten.Image
	switch piece.Type() {
	case chess.PAWN:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_pawn.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_pawn.png")
		}
	case chess.BISHOP:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_bishop.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_bishop.png")
		}
	case chess.KING:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_king.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_king.png")
		}
	case chess.KNIGHT:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_knight.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_knight.png")
		}
	case chess.QUEEN:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_queen.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_queen.png")
		}
	case chess.ROOK:
		if piece.IsBlack() {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/black_rook.png")
		} else {
			img, _, _ = ebitenutil.NewImageFromFile("assets/pieces/white_rook.png")
		}
	}
	geo := ebiten.GeoM{}
	geo.Scale(2, 2)
	if selected {
		geo.Translate(float64(x-GRIDSIZE/4), float64(y-GRIDSIZE/4))
	} else {
		geo.Translate(float64((7*GRIDSIZE)-(GRIDSIZE*x)+2), float64((7*GRIDSIZE)-(GRIDSIZE*y)))
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
