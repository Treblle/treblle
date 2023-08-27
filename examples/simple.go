package examples

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"gitub.com/treblle/treblle/pkg/app"
	"gitub.com/treblle/treblle/pkg/console"
)

var (
	AppPath = "/.examples"
)

func main() {
	app := app.New("example", AppPath, 12)
	app.Boot()
}

func RegisterCommands(app *app.App) {
	app.Console.Register(
		console.
			Build("test").
			Description("This is a test command").
			Action(func(cmd *cobra.Command, args []string) {
				log.Info("Output from the test command\n\n")
			},
			),
	)
}
