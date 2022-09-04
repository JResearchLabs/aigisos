package main

import (
	_ "embed"

	"github.com/JResearchLabs/Flutechain/command/root"
	"github.com/JResearchLabs/Flutechain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
