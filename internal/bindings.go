package bindings

import (
	"linewalker/internal/physic"
	"linewalker/internal/worldmap"
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

	w.Bind("requestObjectCoord", requestObjectCoord)
	w.Bind("requestLine", requestLine)

	w.Bind("setPhysic", physic.Const.Set)
	w.Bind("getPhysic", physic.Const.Get)

	w.Bind("setMap", worldmap.CurrentMap.Set)
	w.Bind("getMap", worldmap.CurrentMap.Get)

	w.Bind("setObject", physic.CurrentState.Set)
	w.Bind("getObject", physic.CurrentState.Get)

	w.Bind("addObject", physic.CurrentState.AddObject)

}

var CurrentView = PlayerView{X: 0, Y: 0, Width: 300, Height: 300}
var N int = 0

func requestObjectCoord(M map[string]any) []physic.Object {
	Input := parseUserInput(M)
	for k, obj := range physic.CurrentState.Obj {
		physic.PFD(&obj, worldmap.CurrentMap.Generate, Input, physic.CurrentState.TimeElapsed())
		if k == 0 {
			CurrentView.SetCoord(int(obj.Coord.X)-CurrentView.Width/2, int(obj.Coord.Y)-CurrentView.Height/2)
		}
		obj.SetScreenCoord(CurrentView.X, CurrentView.Y)
	}
	physic.CurrentState.Time = time.Now()
	return physic.CurrentState.Obj

}

const LineBuffer = 20

func requestLine() map[int]int {
	var out = map[int]int{}
	for x := -LineBuffer; x < CurrentView.Width+LineBuffer; x++ {
		rx := x + CurrentView.X
		out[x] = int(worldmap.CurrentMap.Generate(float64(rx))) - CurrentView.Y
	}
	return out
}
