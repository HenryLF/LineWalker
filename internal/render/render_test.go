package render

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
)

func printTestImgRGBA(im *image.RGBA, n int, t *testing.T) {
	out, err := os.Create(fmt.Sprintf("./tests/out_%v.png", n))
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	if err := png.Encode(out, im); err != nil {
		t.Fatal(err)
	}
}
func printTestImg(im *image.Image, n int, t *testing.T) {
	out, err := os.Create(fmt.Sprintf("./tests/out_%v.png", n))
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	if err := png.Encode(out, *im); err != nil {
		t.Fatal(err)
	}
}

func logToFile(logs []string) {
	out, err := os.Create("./tests/log.txt")
	if err != nil {
		return
	}
	defer out.Close()
	for _, log := range logs {
		out.WriteString(log)
		out.WriteString("\n")
	}

}

func TestDraw(t *testing.T) {
	P := PlayerView{0, 0, 500, 500}
	im := NewImage(P)

	DrawFloor(im, P, func(f float64) float64 { return 300 })
	printTestImgRGBA(im, 0, t)
}

func TestEntity(t *testing.T) {
	printTestImg(&PlayerIdle.spriteSheet, 0, t)
}

func TestRenderSprite(t *testing.T) {
	P := PlayerView{0, 0, 500, 500}
	for n := range PlayerIdle.maxN {
		im := NewImage(P)
		PlayerIdle.RenderSprite(im, 50, 50, n)
		printTestImgRGBA(im, n, t)
	}
}
