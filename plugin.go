package main

type Plugin interface {
	Name() string
	Help() string
	Execute(args []string) string
}

var Plugins = make([]Plugin, 0)
