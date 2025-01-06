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

func (P *PlayerView) Center(Obj physic.Object) {
	P.X = int(Obj.Coord.X) - P.Width/2
	P.Y = int(Obj.Coord.Y) - P.Height/2
}
func parseUserInput(M map[string]bool) physic.UserInput {
	out := physic.UserInput{}
	for k, it := range M {
		switch k {
		case "Down":
			out.Down = it
		case "Left":
			out.Left = it
		case "Right":
			out.Right = it
		case "Up":
			out.Up = it

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

func requestObjectCoord(M map[string]bool) []physic.Object {
	Input := parseUserInput(M)
	Colision := physic.CurrentState.ColisionMap()
	for k, obj := range physic.CurrentState.Obj {
		if k == 0 {
			physic.PFD(&obj, worldmap.CurrentMap.Generate, Input, Colision[k], physic.CurrentState.TimeElapsed())
			CurrentView.Center(obj)
		} else {
			physic.PFD(&obj, worldmap.CurrentMap.Generate, physic.UserInput{}, Colision[k], physic.CurrentState.TimeElapsed())
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
