package alpine

import (
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
)

const (
	Prevent  = "prevent"
	Stop     = "stop"
	Outside  = "outside"
	Window   = "window"
	Document = "document"
	Once     = "once"
	Debounce = "debounce"
	Throttle = "throttle"
	Self     = "self"
	Camel    = "camel"
	Dot      = "dot"
	Passive  = "passive"
	Capture  = ".capture"
)

const (
	Shift    = "shift"
	Enter    = "enter"
	Space    = "space"
	Ctrl     = "ctrl"
	Cmd      = "cmd"
	Meta     = "meta"
	Alt      = "alt"
	Up       = "up"
	Right    = "right"
	Down     = "down"
	Left     = "left"
	Escape   = "escape"
	Tab      = "tab"
	CapsLock = "caps-lock"
	Equal    = "equal"
	Period   = "period"
	Comma    = "comma"
	Slash    = "slash"
)

const (
	dot = "."
)

func CreateEvent(name string, script string, modifiers ...string) gox.Node {
	m := strings.Join(modifiers, ".")
	if len(m) > 0 {
		m = dot + m
	}
	return gox.CreateAttribute[string]("@" + name + m)(script)
}

func Click(script string, modifiers ...string) gox.Node {
	return CreateEvent("click", script, modifiers...)
}

func Input(script string, modifiers ...string) gox.Node {
	return CreateEvent("input", script, modifiers...)
}

func Submit(script string, modifiers ...string) gox.Node {
	return CreateEvent("submit", script, modifiers...)
}

func Scroll(script string, modifiers ...string) gox.Node {
	return CreateEvent("scroll", script, modifiers...)
}

func KeyUp(script string, modifiers ...string) gox.Node {
	return CreateEvent("keyup", script, modifiers...)
}

func MouseMove(script string, modifiers ...string) gox.Node {
	return CreateEvent("mousemove", script, modifiers...)
}
func MouseEnter(script string, modifiers ...string) gox.Node {
	return CreateEvent("mouseenter", script, modifiers...)
}

func MouseLeave(script string, modifiers ...string) gox.Node {
	return CreateEvent("mouseleave", script, modifiers...)
}
