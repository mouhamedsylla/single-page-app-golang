package main

import (
	"fmt"
	"spaGo/App/wasm/pages"
	framego "spaGo/frameGo"
	"syscall/js"
)

func main() {
	ch := make(chan chan struct{})
	fmt.Println("Hello Browser from Go.")

	router := &framego.Router{}
	routes := []string{
		"/",
		"/posts",
		"/settings",
	}
	pages := []framego.Component{
		&pages.Dashboard{},
		&pages.Posts{},
		&pages.Settings{},
	}
	for i, r := range routes {
		router.AddRoute(r, pages[i].Render())
	}
	home := js.Global().Get("location").Get("href").String()
	router.ResolvRoute(home)
	elements := js.Global().Get("document").Call("querySelectorAll", "a")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i)
		element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
			event := p[0]
			event.Call("preventDefault")
			url := element.Get("href").String()
			router.ResolvRoute(url)
			return nil
		}))
	}
	

	<-ch
}
