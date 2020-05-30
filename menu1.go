package goTerminalMenu

import (
	"github.com/eiannone/keyboard"
	goTerminal "github.com/leandroveronezi/go-terminal"
)

func Menu1(FirstLine int, Config config, Opcoes ...string) int {

	Col := 1
	line := FirstLine

	goTerminal.SetSGR(goTerminal.Reset)
	goTerminal.CursorLineColumn(line, Col)

	renderUnchecked := func(pos int) {

		goTerminal.SetSGR(Config.UncheckedColor)
		goTerminal.CursorLineColumn(line+(pos-1), Col+2)
		goTerminal.Print(string(Config.Unchecked))

	}

	renderChecked := func(pos int) {

		goTerminal.SetSGR(Config.CheckedColor)
		goTerminal.CursorLineColumn(line+(pos-1), Col+2)
		goTerminal.Print(string(Config.Checked))

	}

	for idx, str := range Opcoes {
		goTerminal.CursorLineColumn(line+idx, Col)

		goTerminal.SetSGR(Config.ParenthesesColor)
		goTerminal.Print(string(Config.StartParentheses) +
			"   " +
			string(Config.EndParentheses))

		goTerminal.SetSGR(Config.LabelColor)
		goTerminal.CursorLineColumn(line+idx, Col+6)
		goTerminal.Print(str)

		renderUnchecked(idx + 1)

	}

	selecionado := 1

	goTerminal.SaveCursorAttrs()

	renderChecked(selecionado)

	keysEvents, err := keyboard.GetKeys(Col)
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

		if event.Key == keyboard.KeyArrowUp {

			if selecionado == 1 {
				continue
			}

			renderUnchecked(selecionado)
			selecionado--
			renderChecked(selecionado)

		} else if event.Key == keyboard.KeyArrowDown || event.Key == keyboard.KeyTab {

			if selecionado >= len(Opcoes) {
				continue
			}

			renderUnchecked(selecionado)
			selecionado++
			renderChecked(selecionado)

		}

		if event.Key == keyboard.KeyEnter {
			break
		}
	}

	goTerminal.RestoreCursorAttrs()
	goTerminal.CursorDown(1)
	goTerminal.CursorColumn(1)
	goTerminal.SetSGR(goTerminal.Reset)

	return selecionado
}
