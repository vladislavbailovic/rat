package main

import (
	"log"
	"os"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	KeyEsc   uint16 = 65307
	KeyUp    uint16 = 65362
	KeyDown  uint16 = 65364
	KeyRight uint16 = 65363
	KeyLeft  uint16 = 65361
)

func main() {
	update()
}

func update() {
	s := hook.Start()
	for {
		i := <-s

		if i.Kind != hook.KeyDown {
			continue
		}

		switch i.Rawcode {
		case KeyEsc:
			os.Exit(0)
		case KeyUp:
			log.Println("up")
			robotgo.MoveRelative(0, -10)
		case KeyDown:
			log.Println("down")
			robotgo.MoveRelative(0, 10)
		case KeyRight:
			log.Println("right")
			robotgo.MoveRelative(10, 0)
		case KeyLeft:
			log.Println("left")
			robotgo.MoveRelative(-10, 0)
		default:
			log.Printf("evt: %v\n", i)
		}
	}
}
