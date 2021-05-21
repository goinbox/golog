package golog

import (
	"github.com/goinbox/color"
)

type colorFunc func(msg []byte) []byte

type consoleFormater struct {
	f Formater

	levelColorFuncs map[int]colorFunc
}

func NewConsoleFormater(f Formater) *consoleFormater {
	c := &consoleFormater{
		f: f,

		levelColorFuncs: map[int]colorFunc{
			LevelDebug:     color.Yellow,
			LevelInfo:      color.Blue,
			LevelNotice:    color.Cyan,
			LevelWarning:   color.Maganta,
			LevelError:     color.Red,
			LevelCritical:  color.Black,
			LevelAlert:     color.White,
			LevelEmergency: color.Green,
		},
	}

	return c
}

func (c *consoleFormater) SetColor(level int, cf colorFunc) *consoleFormater {
	c.levelColorFuncs[level] = cf

	return c
}

func (c *consoleFormater) Format(level int, msg []byte) []byte {
	return c.levelColorFuncs[level](c.f.Format(level, msg))
}
