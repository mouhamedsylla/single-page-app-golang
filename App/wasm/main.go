package main

import (
	"fmt"
	"spaGo/App/wasm/pages"
	framego "spaGo/frameGo"
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

	router.ResolvRoute()

	// element := js.Global().Get("document").Call("getElementById", "link-2")
	// fmt.Println(element.String())
	// element.Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	args[0].Call("preventDefault")
	// 	//router.ResolvRoute()
	// 	return nil
	// }))
	// js.Global().Get("window").Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

	// 	return nil
	// }))

	// document := js.Global().Get("document")
	// document.Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	document.Call("addEvenListener", "click", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
	// 		if this.Call("target", "matches", "[data-link]").Bool() {
	// 			p[0].Call("preventDefault")
	// 		}

	// 		return nil
	// 	}))

	// 	return nil
	// }))

	//js.Global().Get("history").Call("pushState", nil, nil, p)
	//elements := js.Global().Get("document").Call("querySelectorAll", "a")

	// for i := 0; i < elements.Length(); i++ {
	// 	element := elements.Index(i)
	// 	element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 		args[0].Call("preventDefault")
	//
	// 		i = elements.Length()+1
	// 		return nil
	// 	}))
	// }

	// document.addEventListener("DOMContentLoaded", () => {
	// 	document.body.addEventListener("click", e => {
	// 		if (e.target.matches("[data-link]")) {
	// 			e.preventDefault();
	// 			navigateTo(e.target.href);
	// 		}
	// 	});

	// 	router();
	// });

	<-ch
}
