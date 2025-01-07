package bindings

import (
	"linewalker/internal/physic"
	"linewalker/internal/worldmap"
	"time"

	webview "github.com/webview/webview_go"
)

var CurrentState = physic.State{Time: time.Now(), Obj: []physic.Object{physic.NewObjectSide(500, 0, 5, 4500)}}

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

	w.Bind("setObject", CurrentState.Set)
	w.Bind("getObject", CurrentState.Get)

	w.Bind("addObject", CurrentState.AddObject)

	w.Bind("setGlobals", CurrentView.Set)
	w.Bind("getGlobals", CurrentView.Get)

	w.Bind("setPlayer", CurrentState.SetPlayer)
	w.Bind("getPlayer", CurrentState.GetPlayer)

}

func requestObjectCoord(M map[string]bool) []physic.Object {
	Input := parseUserInput(M)
	CurrentState.UpdateState(Input, worldmap.CurrentMap.Generate, CurrentView.ScreenTransform)
	CurrentView.Center(CurrentState.Obj[0])
	CurrentState.ScreenCoordFromTransform(CurrentView.ScreenTransform)
	CurrentState.Time = time.Now()
	return CurrentState.Obj

}

const LineBuffer = 20

func requestLine() map[int]int {
	var out = map[int]int{}
	for x := -LineBuffer; x < CurrentView.Width+LineBuffer; x++ {
		out[x] = worldmap.CurrentMap.GenerateFromTransform(x, CurrentView.CoordTransformX, CurrentView.ScreenTransformY)
	}
	return out
}
