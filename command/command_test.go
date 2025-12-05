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
				CMD:  "plugins",
				Args: nil,
				Desc: "查看已加载插件",
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
