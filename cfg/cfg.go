package main

import (
	"github.com/robfig/config"
)

func main() {
	c := config.NewDefault()
	//c.AddSection("Section")
	//c.AddOption("Section", "option", "value2")
	c.AddOption("", "option", "value2")
	c.WriteFile("c.cfg", 0644, "")
}
