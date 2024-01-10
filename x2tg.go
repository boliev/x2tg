package x2tg

import "fmt"

type App struct {
}

func (a App) Run() {
	fmt.Println("I'm the service! I'm working!")
}
