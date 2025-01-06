package bindings

import (
	"linewalker/internal/physic"
	"linewalker/internal/worldmap"
	"time"

	webview "github.com/webview/webview_go"
)

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
	physic.CurrentState.UpdateState(Input, worldmap.CurrentMap.Generate, CurrentView.ScreenTransform)
	CurrentView.Center(physic.CurrentState.Obj[0])
	physic.CurrentState.ScreenCoordFromTransform(CurrentView.ScreenTransform)
	physic.CurrentState.Time = time.Now()
	return physic.CurrentState.Obj

}

const LineBuffer = 20

func requestLine() map[int]int {
	var out = map[int]int{}
	for x := -LineBuffer; x < CurrentView.Width+LineBuffer; x++ {
		out[x] = worldmap.CurrentMap.GenerateFromTransform(x, CurrentView.CoordTransformX, CurrentView.ScreenTransformY)
	}
	return out
}
