package command

import "testing"

func TestCommands_String(t *testing.T) {
	c := NewCommands(
		NewCommand("查看帮助", "help", "?", "?", "帮助"),
		NewCommand("Ping", "ping", "p"),
	)

	t.Log("\n", c.String())
}
