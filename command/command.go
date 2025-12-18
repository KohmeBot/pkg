package command

import (
	"fmt"
	"strings"
)

// ======================= Arg =======================

type Arg struct {
	Name     string
	Optional bool
	Desc     string
}

func (a Arg) String() string {
	if a.Optional {
		return "[" + a.Name + "]"
	}
	return "<" + a.Name + ">"
}

func (a Arg) Description() string {
	return fmt.Sprintf("  %s  %s", a.String(), a.Desc)
}

// ======================= Command =======================

type Command struct {
	CMD  string
	Args []Arg
	Desc string
}

func (c Command) String() string {
	var b strings.Builder

	// 第一行：命令 + 参数
	b.WriteString(c.CMD)

	for _, arg := range c.Args {
		b.WriteByte(' ')
		b.WriteString(arg.String())
	}

	// 命令描述
	if c.Desc != "" {
		b.WriteString(" — ")
		b.WriteString(c.Desc)
	}

	// 参数说明
	for i, arg := range c.Args {
		if arg.Desc != "" {
			if i == 0 {
				b.WriteByte('\n')
			}
			b.WriteString(arg.Description())
			b.WriteByte('\n')
		}

	}

	return b.String()
}

// ======================= HelpTemplate =======================

type HelpTemplate struct {
	PluginName string
	PluginDesc string
	Commands   []Command
}

func (h HelpTemplate) wrap(text string, w int) string {
	if w <= 0 {
		return text
	}

	var out strings.Builder
	var line strings.Builder

	for _, r := range text {
		line.WriteRune(r)
		if line.Len() >= w {
			out.WriteString(line.String())
			out.WriteByte('\n')
			line.Reset()
		}
	}

	if line.Len() > 0 {
		out.WriteString(line.String())
	}

	return out.String()
}

func (h HelpTemplate) String() string {
	var b strings.Builder

	// 标题（短行，适合 QQ）
	b.WriteString("【")
	b.WriteString(h.PluginName)
	b.WriteString("】\n")

	if h.PluginDesc != "" {
		b.WriteString(h.PluginDesc)
		b.WriteString("\n")
	}

	b.WriteString("命令列表：\n")

	for _, c := range h.Commands {
		b.WriteString(c.String())
		b.WriteByte('\n')
	}

	return strings.TrimRight(b.String(), "\n")
}
