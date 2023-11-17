package pages

type Settings struct {}

func (s *Settings) Render() string{
	return ` <h1>Settings</h1>
	<p>Manage your privacy and configuration.</p>`
}