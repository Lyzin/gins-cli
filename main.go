/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"embed"
	"github.com/Lyzin/gins-cli/actions"
	"github.com/Lyzin/gins-cli/cmd"
)

//go:embed assets/tpl
var Resource embed.FS

func main() {
	actions.Assets = Resource
	cmd.Execute()
}