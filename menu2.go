package goTerminalMenu

import (
	"github.com/eiannone/keyboard"
	goTerminal "github.com/leandroveronezi/go-terminal"
)

func Menu2(FirstLine int, Config config, Opcoes map[int]Options) map[int]Options {

	Col := 2
	line := FirstLine

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.CursorLineColumn(line, Col)

	renderStatus := func(pos int) {

		goTerminal.CursorLineColumn(line+(pos-1), Col+2)

		if Opcoes[pos].Immutable {
			goTerminal.SetSGR(Config.ImmutableColor)
		} else {
			goTerminal.SetSGR(Config.UncheckedColor)
		}

		goTerminal.Print(string(Config.Unchecked))

		if Opcoes[pos].Status {
			goTerminal.CursorLineColumn(line+(pos-1), Col+2)

			if Opcoes[pos].Immutable {
				goTerminal.SetSGR(Config.ImmutableColor)
			} else {
				goTerminal.SetSGR(Config.CheckedColor)
			}

			goTerminal.Print(string(Config.Checked))

		}

	}

	for idx, str := range Opcoes {

		goTerminal.CursorLineColumn(line+idx-1, Col)

		goTerminal.SetSGR(Config.ParenthesesColor)
		goTerminal.Print(string(Config.StartParentheses) +
			"   " +
			string(Config.EndParentheses))

		goTerminal.SetSGR(Config.LabelColor)
		goTerminal.CursorLineColumn(line+idx-1, Col+6)
		goTerminal.Print(str.Label)

		renderStatus(idx)

	}

	selecionado := 1

	goTerminal.CursorLineColumn(line+len(Opcoes), Col)
	goTerminal.SaveCursorAttrs()

	renderUnpoint := func() {

		goTerminal.CursorLineColumn(line+(selecionado-1), Col-1)
		goTerminal.Print(" ")

	}

	renderPoint := func() {

		goTerminal.CursorLineColumn(line+(selecionado-1), Col-1)
		goTerminal.SetSGR(Config.PointColor)
		goTerminal.Print(string(Config.Pointer))

	}

	renderUnpoint()
	renderPoint()

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeySpace {

			if Opcoes[selecionado].Immutable {
				continue
			}

			aux := Opcoes[selecionado]
			aux.Status = !aux.Status
			Opcoes[selecionado] = aux
			renderStatus(selecionado)

		} else if event.Key == keyboard.KeyArrowUp {

			if selecionado == 1 {
				continue
			}

			renderUnpoint()
			selecionado--
			renderPoint()

		} else if event.Key == keyboard.KeyArrowDown || event.Key == keyboard.KeyTab {

			if selecionado >= len(Opcoes) {
				continue
			}

			renderUnpoint()
			selecionado++
			renderPoint()

		}

		if event.Key == keyboard.KeyEnter {
			break
		}
	}

	goTerminal.RestoreCursorAttrs()
	goTerminal.CursorDown(1)
	goTerminal.CursorColumn(1)
	goTerminal.SetSGR(goTerminal.Reset)

	return Opcoes
}
