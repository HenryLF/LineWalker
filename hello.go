package main

import (
	bindings "linewalker/internal"
	"path/filepath"
	"reflect"

	webview "github.com/webview/webview_go"
)

func ConvertMapToStruct(m map[string]interface{}, s interface{}) error {
	stValue := reflect.ValueOf(s).Elem()
	sType := stValue.Type()
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		if value, ok := m[field.Name]; ok {
			stValue.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(1080, 1080, webview.HintNone)
	bindings.RegisterBindings(w)
	url, _ := filepath.Abs("./static/index.html")
	url = filepath.Join("file://", url)
	w.Navigate(url)
	w.Run()
}
