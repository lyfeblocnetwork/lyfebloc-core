package main

import (
	_ "embed"

	"github/lyfeblocnetwork/lyfebloc-core/command/root"
	"github/lyfeblocnetwork/lyfebloc-core/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
