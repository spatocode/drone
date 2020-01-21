package karid

type Renderer interface {
	Render(Scene)
	Draw(Window)
	SetSceneSize(int, int)
}
