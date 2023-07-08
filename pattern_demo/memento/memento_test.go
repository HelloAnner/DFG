package memento

import "testing"

//
// @author Anner on 2021/10/6

func TestMemento(t *testing.T) {
	game := Game{
		hp: 10,
		mp: 20,
	}
	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()

	// Current HP:10,MP:20
	// Current HP:7,MP:18
	// Current HP:10,MP:20
}
