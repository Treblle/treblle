package app

import (
	"os"
	"os/user"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"gitub.com/treblle/treblle/pkg/config"
	"gitub.com/treblle/treblle/pkg/console"
	eventDispatcher "gitub.com/treblle/treblle/pkg/event/dispatcher"
	"gitub.com/treblle/treblle/pkg/http"
	queueDispatcher "gitub.com/treblle/treblle/pkg/queue/dispatcher"
	"gitub.com/treblle/treblle/pkg/storage"
)

// App is the primary struct for your application.
type App struct {
	Name       string
	Config     config.Config
	Console    *console.Registry
	Dispatcher *eventDispatcher.Dispatcher
	Http       *http.PendingRequest
	Path       string
	Queue      *queueDispatcher.Dispatcher
	Storage    storage.Storage
}

// New will create a new App using the users home directory for the path
func New(name string, path string, workers int) *App {
	usr, err := user.Current()
	if err != nil {
		log.Printf("Error:%v\n", err)

		os.Exit(1)
	}

	return &App{
		Name:       name,
		Config:     *config.New(),
		Console:    console.New(&cobra.Command{Use: name}),
		Dispatcher: eventDispatcher.New(),
		Http:       http.New(),
		Path:       usr.HomeDir + path,
		Queue:      queueDispatcher.New(workers),
		Storage:    *storage.New(),
	}
}

// Boot will ensure that there is a directory available for the App Path
func (app *App) Boot() {
	app.Storage.At(app.Path)

	if !app.Storage.Exists() {
		if err := app.Storage.CreateDir(0755); err != nil {
			log.Errorf("Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}
}
