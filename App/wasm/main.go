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
	routes := []string{"/", "/posts", "/settings"}
	pages := []framego.Component{&pages.Dashboard{}, &pages.Posts{}, &pages.Settings{}}

	for i, r := range routes {
		router.AddRoute(r, pages[i].Render())
	}

	router.ResolvRoute()

	<-ch
}
