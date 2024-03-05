package main

import (
	_ "embed"

	"github.com/JResearchLabs/aigisos/command/root"
	"github.com/JResearchLabs/aigisos/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
