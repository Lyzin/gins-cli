/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"embed"
	"gins_cli/actions"
	"gins_cli/cmd"
)

//go:embed assets/tpl
var Resource embed.FS

func main() {
	actions.Assets = Resource
	cmd.Execute()
}