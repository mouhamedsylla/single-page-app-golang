package framego

import (
	"regexp"
	"syscall/js"
)

type Route struct {
	path string
	view string
}

type Router struct {
	routes []Route
}

func (r *Router) AddRoute(p string, v string) {
	myRoute := Route{
		path: p,
		view: v,
	}
	r.routes = append(r.routes, myRoute)
}

// func (r *Router) ResolvRoute() interface{} {
// 	URL := js.Global().Get("location").Get("href").String()

// 	element := js.Global().Get("document").Call("querySelector", "#App")
// 	match := MatchPath(URL)
// 	for i, route := range r.routes {
// 		if route.path == match {
// 			element.Set("innerHTML", route.view)
// 			js.Global().Get("history").Call("pushState", js.Null(), "", route.path)
// 			break
// 		 } else if i == len(r.routes)-1 {
// 			element.Set("innerHTML", "<h1>Error 404: Page Not Found")
// 		}
// 	}
// 	return nil
// }

// func (r *Router) ResolvRoute() {
	
// 	location := js.Global().Get("location")
// 	u := location.Get("href").String()
// 	parsedURL := js.Global().Get("URL").New(u)
// 	path := parsedURL.Get("pathname").String()

// 	element := js.Global().Get("document").Call("querySelector", "#App")

// 	match := false
// 	for _, route := range r.routes {
// 		if route.path == path {
// 			match = true
// 			r.renderView(element, route.view)
// 			js.Global().Get("history").Call("pushState", js.Null(), "", route.path)
// 			break
// 		}
// 	}

// 	if !match {
// 		r.renderView(element, "<h1>Erreur 404 : Page non trouvée</h1>")
// 		js.Global().Get("history").Call("pushState", js.Null(), "", "/404")
// 	}
// }

func (r *Router) ResolvRoute(URL string) interface{} {
	element := js.Global().Get("document").Call("querySelector", "#App")
	match := MatchPath(URL)
	for i, route := range r.routes {
		if route.path == match {
			element.Set("innerHTML", route.view)
			js.Global().Get("history").Call("pushState", js.Null(), "", route.path)
			break
		 } else if i == len(r.routes)-1 {
			element.Set("innerHTML", "<h1>Error 404: Page Not Found")
		}
	}
	return nil
}

// func (r *Router) renderView(element js.Value, view string) {
// 	element.Set("innerHTML", view)
// }

func MatchPath(url string) string {
	pathRegex := regexp.MustCompile(`http://[^/]+(/.*)?`)
	match := pathRegex.FindStringSubmatch(url)
	if match != nil {
		return match[1]
	}
	return ""
}
