package chain

import (
	"github.com/wdvxdr1123/ZeroBot/message"
	"strings"
)

// MessageChain 消息链
type MessageChain []message.MessageSegment

var lb = message.Text("\n")
var empty = message.Text(" ")

// Join 加入一条消息
func (c *MessageChain) Join(segment message.MessageSegment) MessageChain {
	*c = append(*c, segment)
	return *c
}

// Line 加入一行消息，在末尾换行
func (c *MessageChain) Line(segments ...message.MessageSegment) MessageChain {
	*c = append(*c, append(segments, lb)...)
	return *c
}

// Split 分割消息，每条消息间会以换行分割
func (c *MessageChain) Split(segments ...message.MessageSegment) MessageChain {
	return c.split(lb, segments...)
}

// SplitEmpty 分割消息，每条消息间会以空格分割
func (c *MessageChain) SplitEmpty(segments ...message.MessageSegment) MessageChain {
	return c.split(empty, segments...)
}

// split 分割消息
func (c *MessageChain) split(sep message.MessageSegment, segments ...message.MessageSegment) MessageChain {
	last := len(segments) - 1
	for idx, segment := range segments {
		if idx == last {
			*c = append(*c, segment)
			break
		}
		*c = append(*c, segment, sep)
	}
	return *c
}

func (c *MessageChain) String() string {
	var builder strings.Builder
	for _, segment := range *c {
		builder.WriteString(segment.String())
	}
	return builder.String()
}
