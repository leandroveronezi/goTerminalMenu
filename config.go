package goTerminalMenu

import goTerminal "github.com/leandroveronezi/go-terminal"

type Options struct {
	Label     string
	Status    bool
	Immutable bool
}

type config struct {
	Pointer          rune
	Checked          rune
	Unchecked        rune
	StartParentheses rune
	EndParentheses   rune

	CheckedColor     goTerminal.SGR
	UncheckedColor   goTerminal.SGR
	PointColor       goTerminal.SGR
	LabelColor       goTerminal.SGR
	ParenthesesColor goTerminal.SGR
	ImmutableColor   goTerminal.SGR
}

func NewConfig() config {

	return config{
		'➤',
		'✔',
		'✘',
		'[',
		']',
		goTerminal.ForegroundGreen,
		goTerminal.ForegroundRed,
		goTerminal.ForegroundBlue,
		goTerminal.ForegroundCyan,
		goTerminal.ForegroundYellow,
		goTerminal.ForegroundLightGray,
	}

}

/*
	✓
	✔
	✘
	░
	▒
	▓
	╔══╗
	║  ║
	╠═╦╣
	╠═╬╣
	╚═╩╝
*/
