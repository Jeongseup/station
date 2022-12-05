package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/stationlove-ventures/station/app"
	"github.com/stationlove-ventures/station/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"station",
		"cosmos",
		app.DefaultNodeHome,
		"station-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
