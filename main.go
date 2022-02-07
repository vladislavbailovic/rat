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

	KeyShift  uint16 = 65505
	KeyRshift uint16 = 65506
	KeyCtrl   uint16 = 65507
	KeyRctrl  uint16 = 65508
	KeySuper  uint16 = 65515
	KeyAlt    uint16 = 65513
	KeyRalt   uint16 = 65514
)

type pressed bool

const (
	notPressed pressed = false
	isPressed  pressed = true
)

type modifierState struct {
	shift pressed
	ctrl  pressed
	alt   pressed
	super pressed
}

func (m *modifierState) update(in hook.Event) {
	if in.Kind != hook.KeyDown && in.Kind != hook.KeyUp {
		return
	}
	state := notPressed
	if in.Kind == hook.KeyDown {
		state = isPressed
	}

	switch in.Rawcode {
	case KeyShift, KeyRshift:
		m.shift = state
	case KeyCtrl, KeyRctrl:
		m.ctrl = state
	case KeyAlt, KeyRalt:
		m.alt = state
	case KeySuper:
		m.super = state
	default:
		log.Printf("")
	}
}

func main() {
	s := hook.Start()
	modifiers := modifierState{}
	baseSpeed := 10

	for {
		i := <-s
		modifiers.update(i)
		speed := baseSpeed

		if modifiers.shift == isPressed {
			speed *= 25
		}
		if modifiers.ctrl == isPressed {
			speed /= 10
		}

		if i.Kind == hook.KeyDown {

			switch i.Rawcode {
			case KeyEsc:
				os.Exit(0)
			case KeyUp:
				log.Println("up")
				robotgo.MoveRelative(0, -1*speed)
			case KeyDown:
				log.Println("down")
				robotgo.MoveRelative(0, speed)
			case KeyRight:
				log.Println("right")
				robotgo.MoveRelative(speed, 0)
			case KeyLeft:
				log.Println("left")
				robotgo.MoveRelative(-1*speed, 0)
			default:
				if modifiers.alt == isPressed && i.Rawcode < 255 {
					switch rune(i.Rawcode) {
					case 'z', 'a', 'q':
						log.Println("left click")
						robotgo.Click("left")
					case 'x', 's', 'w':
						log.Println("mid click")
						robotgo.Click("center")
					case 'c', 'd', 'e':
						log.Println("right click")
						robotgo.Click("right")
					}
				}
				log.Printf("evt: %v [%v]\n", i, modifiers)
			}
		}
	}
}
