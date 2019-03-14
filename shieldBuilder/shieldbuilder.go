package shieldbuilder

import "strings"

type Shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

func (sh *Shield) String() string {
	s := "Shield status:\n"

	s += "Front: " + getShieldStatus(sh.front) + "\n"
	s += "Back: " + getShieldStatus(sh.back) + "\n"
	s += "Left: " + getShieldStatus(sh.left) + "\n"
	s += "Right: " + getShieldStatus(sh.right) + "\n"

	return s
}

func getShieldStatus(shield bool) string {
	if shield {
		return "On"
	}

	return "Off"
}

type ShieldBuilder struct {
	code string
}

func NewShieldBuilder() *ShieldBuilder {
	return new(ShieldBuilder)
}

func (sh *ShieldBuilder) RaiseFront() *ShieldBuilder {
	sh.code += "F"
	return sh
}

func (sh *ShieldBuilder) RaiseBack() *ShieldBuilder {
	sh.code += "B"
	return sh
}

func (sh *ShieldBuilder) RaiseLeft() *ShieldBuilder {
	sh.code += "L"
	return sh
}

func (sh *ShieldBuilder) RaiseRight() *ShieldBuilder {
	sh.code += "R"
	return sh
}

func (sh *ShieldBuilder) Build() *Shield {
	code := sh.code

	return &Shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		left:  strings.Contains(code, "L"),
		right: strings.Contains(code, "R"),
	}
}
