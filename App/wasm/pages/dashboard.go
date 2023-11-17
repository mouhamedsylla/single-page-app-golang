package pages

import (
	"bytes"
	"text/template"
)

type Dashboard struct {
	Name string
	Job  string
	Languages map[string][]string
}

func (d *Dashboard) Render() string {
	me := Dashboard{
		Name: "Mouhamed Sylla",
		Job:  "Developper",
		Languages: map[string][]string{
			"Language": {"Java", "Python", "JavaScript", "C", "Golang"},
		},
	}


	component := `
	<div class="profile">
		<div class="css-after-gradiant">
  			<img src="profile.png" />
		</div>
		<h1>Hi, my name is {{.Name}}</h1>
	</div>
	<p>
		<h3> I am a {{.Job}} </h3>
		<h4>My favorite programming languages</h4>
	</p>
	<ul>
		{{range .Languages.Language}}<li>{{.}}</li>{{end}}
	</ul>
	<p>
		<a href="/posts" data-link>View recent posts</a>.
	</p>`

	var resultBuffer bytes.Buffer
	tpml, _ := template.New("template").Parse(component)
	err := tpml.Execute(&resultBuffer, me)
	if err != nil {
		panic(err)
	}
	return resultBuffer.String()
}
