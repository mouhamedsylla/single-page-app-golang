package pages

type Posts struct {

}

func (p *Posts) Render() string {
	return ` <h1>Posts</h1>
	<p>You are viewing the posts!</p>`
}