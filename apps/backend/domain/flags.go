package domain

import "flag"

var ShouldSeed bool

func SetupFlags() {
	flag.BoolVar(&ShouldSeed, "seed", false, "Activates database seed")

	flag.Parse()
}
