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

func (r *Router) ResolvRoute() {
	URL := js.Global().Get("location").Get("href").String()
	
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
}

func MatchPath(url string) string {
	pathRegex := regexp.MustCompile(`http://[^/]+(/.*)?`)
	match := pathRegex.FindStringSubmatch(url)
	if match != nil {
		return match[1]
	}
	return ""
}
