package main

import "fmt"

type Color string

func (self Color) Text(text string) Color {
	self += Color(text)
	return self
}

func (self Color) Reset() Color {
	self += "\u001b[0m"
	return self
}

func (self Color) Cyan(text string) Color {
	self = Color(fmt.Sprintf("%s\u001b[36m%s%s", self, text, self.Reset()))
	return self
}

func (self Color) Red(text string) Color {
	self = Color(fmt.Sprintf("%s\u001b[31m%s%s", self, text, self.Reset()))
	return self
}
