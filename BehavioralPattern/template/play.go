package template

type Game interface {
	Init() error
	Play()
	Stop(time int)
}

func Play(game Game) {
	err := game.Init()
	if err != nil {
		return
	}
	game.Play()
	game.Stop(10)
}
