# Treblle CLI

A framework for building CLI applications fluently in Go.

## Usage

To start you need to create your `main.go` and create a new application

```go

const (
	AppDir = "/.your-app"
)

func main() {
    app := app.New(
        "CLI Name", // The Name of your CLI Application
        AppDir, // The path within ~/.config to create CLI configuration
        1000, // The number of parallel workers you want to be available for dispatching tasks
    )

    app.Boot() // This will boot the application, setting the storage path for this run.

    // If you need to load your CLI config from a file you can do that using a fluent interface
    app.Config.LoadFromFile(app.Storage.Path + "/config.go")
}
```

Once you have a basic application set up, you can move onto registering your commands using the fluent builder interface

```go
func main() {
    ... other code goes here

    RegisterCommands(app) // Register all the commands in your application

    app.Console.RootCmd.Execute() // Execute your main command
}

func RegisterCommands(app *app.App) {
    app.Console.Register(
		console.
			Build("test"). // The name of the command
			Description("Test command"). // the description (short) for the command
			Action(func(cmd *cobra.Command, args []string) {
				log.Info("Test Command Called") // what you want this command to do
			}
        ),
	)
}
```

## Using the HTTP Client

```go
func main() {
    ... other code goes here

    request := app.Http.Base("https://api.treblle.com/").WithToken("123123")

	var message Message

	response, err := request.Get("/")
	if err != nil {
		log.Errorf("Request error: %v\n", err)
		os.Exit(1)
	}

	log.Infof("Success: %v", response.JSON(&message))
}
```

## Using the Applications Queue

You can start your applications queue runner, which will listen on a wait group for Jobs to be passed into it

```go
func main() {
    ... other code goes here

    runner := app.Queue

	runner.Run()
}

func SomeOtherMethod() {
    var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			job := &ExampleJob{Message: fmt.Sprintf("Hello, I am job number %d", i)}
			dispatcher.JobQueue <- job
		}(i)
	}
	wg.Wait()
}
```

## Using the Event Dispatcher

```go
func main() {
    ... other code goes here

    // Register listeners
	app.Dispatcher.Event("UserRegistered").
		Listen(func(data event.EventData) {
			userData, _ := data.(map[string]string)
			println("Send welcome email to:", userData["email"])
		}).
		Listen(func(data event.EventData) {
			println("Log user registration to audit trail.")
		})
}

func SomeOtherMethod() {
    // Dispatch event
	app.Dispatcher.Event("UserRegistered").Dispatch(map[string]string{
		"email": "user@example.com",
		"name":  "John Doe",
	})
}
```

