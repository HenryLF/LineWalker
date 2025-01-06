package render

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type Entity struct {
	spriteSheet                                     image.Image
	spriteWidth, spriteHeight, maxN, delayN, dx, dy int
}

func (E Entity) RenderSprite(im *image.RGBA, x, y, N int) {
	N /= E.delayN
	N %= E.maxN
	dst := image.Rect(x-E.dx, y-E.dy, x+E.dx, y+E.dy)
	sp := image.Point{E.spriteWidth * N, 0}

	draw.Draw(im, dst, E.spriteSheet, sp, draw.Src)
}

func newEntity(path string, N, delayN int) *Entity {
	out := new(Entity)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("No asset", path)
	}
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal("Bad Image", path)
	}
	out.spriteSheet = img
	out.spriteWidth = img.Bounds().Dx() / N
	out.spriteHeight = img.Bounds().Dy()
	out.dx = out.spriteWidth / 2
	out.dy = out.spriteHeight / 2
	out.maxN = N
	out.delayN = delayN

	return out
}

var PlayerIdle = newEntity("./internal/render/assets/playerIdle.png", 4, 20)
