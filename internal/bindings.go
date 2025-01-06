package bindings

import (
	"linescroller/internal/physic"
	"linescroller/internal/worldmap"
	"time"

	webview "github.com/webview/webview_go"
)

type PlayerView struct {
	X, Y, Width, Height int
}

func (P *PlayerView) SetSize(w, h int) {
	P.Width = w
	P.Height = h
}
func (P *PlayerView) SetCoord(x, y int) {
	P.X = x
	P.Y = y
}

func parseUserInput(M map[string]any) physic.UserInput {
	out := physic.UserInput{}
	for k, it := range M {
		switch k {
		case "Down":
			out.Down = it.(bool)
		case "Left":
			out.Left = it.(bool)
		case "Right":
			out.Right = it.(bool)
		case "Up":
			out.Up = it.(bool)

		}
	}
	return out
}

func RegisterBindings(w webview.WebView) {
	w.Bind("setPlayerView", CurrentView.SetSize)
	w.Bind("requestPlayerCoord", requestPlayerCoord)
	w.Bind("requestLine", requestLine)
	w.Bind("setPhysic", physic.Const.Set)
	w.Bind("getPhysic", physic.Const.Get)
	w.Bind("setMap", worldmap.CurrentMap.Set)
	w.Bind("getMap", worldmap.CurrentMap.Get)

}

var CurrentView = PlayerView{X: 0, Y: 0, Width: 300, Height: 300}
var CurrentState = physic.State{Time: time.Now(), Obj: []physic.Object{physic.NewObject(200, 0, 28, 1)}}
var N int = 0

type PlayerCoord struct{ X, Y, XSpeed, YSpeed, XAbs, YAbs int }

func requestPlayerCoord(M map[string]any) PlayerCoord {
	Input := parseUserInput(M)
	for _, obj := range CurrentState.Obj {
		physic.PFD(&obj, worldmap.CurrentMap.Generate, Input, CurrentState.TimeElapsed())
		// physic.PFD(&obj, worldmap.SlopeFloor(.2, 300), Input, CurrentState.TimeElapsed())
	}
	// log.Println(CurrentState.TimeElapsed())
	CurrentState.Time = time.Now()
	player := CurrentState.Obj[0]
	CurrentView.SetCoord(int(player.Coord.X)-CurrentView.Width/2, int(player.Coord.Y)-CurrentView.Height/2)
	return PlayerCoord{X: int(player.Coord.X) - CurrentView.X, Y: int(player.Coord.Y) - CurrentView.Y, XSpeed: int(player.Speed.X), YSpeed: int(player.Speed.Y), XAbs: int(player.Coord.X), YAbs: int(player.Coord.Y)}

}

const LineBuffer = 20

func requestLine() map[int]int {
	var out = map[int]int{}
	// f := worldmap.SlopeFloor(.2, 300)
	// log.Println(CurrentView)
	for x := -LineBuffer; x < CurrentView.Width+LineBuffer; x++ {
		rx := x + CurrentView.X
		out[x] = int(worldmap.CurrentMap.Generate(float64(rx))) - CurrentView.Y
		// out[x] = int(f(float64(rx))) + CurrentView.Y
	}
	return out
}
