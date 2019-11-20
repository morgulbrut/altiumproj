package main

import (
	"github.com/morgulbrut/altiumproj/cmd"
	"github.com/morgulbrut/color256"
)

func main() {
	logo()
	cmd.Execute()
}

func logo() {
	color256.Init()
	color256.PrintRandom(" █████╗               ██╗      ██╗                                            ")
	color256.PrintRandom("██╔══██╗██╗  ████████╗██║██╗   ██║███╗   ███╗██████╗ ██████╗  ██████╗      ██╗")
	color256.PrintRandom("██║  ██║██║  ╚══██╔══╝██║██║   ██║████╗ ████║██╔══██╗██╔══██╗██╔═══██╗     ██║")
	color256.PrintRandom("███████║██║     ██║   ██║██║   ██║██╔████╔██║██████╔╝██████╔╝██║   ██║     ██║")
	color256.PrintRandom("██╔══██║██║     ██║   ██║██║   ██║██║╚██╔╝██║██╔═══╝ ██╔══██╗██║   ██║     ██║")
	color256.PrintRandom("██║  ██║███████╗██║   ██║╚██████╔╝██║ ╚═╝ ██║██║     ██║  ██║╚██████╔╝██   ██║")
	color256.PrintRandom("██║  ╚═╝╚══════╝██║   ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝     ██║  ╚═╝ ╚═════╝ ╚█████╔╝")
	color256.PrintRandom("╚═╝             ╚═╝                                  ╚═╝               ╚════╝ ")
}
