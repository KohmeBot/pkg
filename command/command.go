package command

import (
	"fmt"
	"strings"
)

// Command 命令，仅作为描述
type Command struct {
	CommandGroup []string
	Desc         string
}

func (c Command) String() string {
	var builder strings.Builder
	latest := len(c.CommandGroup) - 1
	for idx, oneCommand := range c.CommandGroup {
		if idx == latest {
			builder.WriteString(fmt.Sprintf("<%s>", oneCommand))
			break
		}
		builder.WriteString(fmt.Sprintf("<%s> ", oneCommand))
	}
	builder.WriteString(fmt.Sprintf(": %s", c.Desc))
	return builder.String()
}

// NewCommand 创建一个命令
//
//	Desc: 描述
//	cms: 命令
func NewCommand(desc string, cms ...string) Command {
	return Command{
		CommandGroup: cms,
		Desc:         desc,
	}
}

type Commands []Command

// NewCommands 创建一个命令列表
func NewCommands(cms ...Command) Commands {
	var c Commands
	c = append(c, cms...)
	return c
}

func (c Commands) String() string {
	var builder strings.Builder
	for _, command := range c {
		builder.WriteString(command.String())
		builder.WriteByte('\n')
	}
	return builder.String()
}
