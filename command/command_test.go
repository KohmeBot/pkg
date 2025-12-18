package command

import (
	"fmt"
	"testing"
)

func TestCommands_String(t *testing.T) {
	help := HelpTemplate{
		PluginName: "core",
		PluginDesc: "kohme 核心插件",
		Commands: []Command{
			{
				CMD: "help",
				Args: []Arg{
					{
						Name: "插件名称",
					},
				},
				Desc: "查看对应插件帮助",
			},
			{
				CMD:  "ping",
				Desc: "ping一下",
			},
			{
				CMD:  "plugin",
				Desc: "查看所有插件",
			},
			{
				CMD: "toggle",
				Args: []Arg{
					{
						Name: "插件名称",
					},
				},
				Desc: "开启/关闭插件",
			},
		},
	}

	fmt.Println(help.String())

}
