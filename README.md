<div align="center">
  <img src="https://treblle-github.s3.amazonaws.com/header.png"/>
</div>
<div align="center">

# Treblle CLI

<a href="https://docs.treblle.com/en/integrations" target="_blank">Integrations</a>
<span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
<a href="http://treblle.com/" target="_blank">Website</a>
<span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
<a href="https://docs.treblle.com" target="_blank">Docs</a>
<span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
<a href="https://blog.treblle.com" target="_blank">Blog</a>
<span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
<a href="https://twitter.com/treblleapi" target="_blank">Twitter</a>
<span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
<a href="https://treblle.com/chat" target="_blank">Discord</a>
<br />

  <hr />
</div>

A framework for building CLI applications fluently in Go.

> This is currently unreleased, and is a work in progress. Breaking Changes are expected.

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

## Examples

- [Simple Example](./examples/simple.go)


## Community 💙

First and foremost: **Star and watch this repository** to stay up-to-date.

Also, follow our [Blog](https://blog.treblle.com), and on [Twitter](https://twitter.com/treblleapi).

You can chat with the team and other members on [Discord](https://treblle.com/chat) and follow our tutorials and other video material at [YouTube](https://youtube.com/@treblle).

[![Treblle Discord](https://img.shields.io/badge/Treblle%20Discord-Join%20our%20Discord-F3F5FC?labelColor=7289DA&style=for-the-badge&logo=discord&logoColor=F3F5FC&link=https://treblle.com/chat)](https://treblle.com/chat)

[![Treblle YouTube](https://img.shields.io/badge/Treblle%20YouTube-Subscribe%20on%20YouTube-F3F5FC?labelColor=c4302b&style=for-the-badge&logo=YouTube&logoColor=F3F5FC&link=https://youtube.com/@treblle)](https://youtube.com/@treblle)

[![Treblle on Twitter](https://img.shields.io/badge/Treblle%20on%20Twitter-Follow%20Us-F3F5FC?labelColor=1DA1F2&style=for-the-badge&logo=Twitter&logoColor=F3F5FC&link=https://twitter.com/treblleapi)](https://twitter.com/treblleapi)

### How to contribute

Here are some ways of contributing to making Treblle better:

- **[Try out Treblle](https://docs.treblle.com/en/introduction#getting-started)**, and let us know ways to make Treblle better for you. Let us know here on [Discord](https://treblle.com/chat).
- Join our [Discord](https://treblle.com/chat) and connect with other members to share and learn from.
- Send a pull request to any of our [open source repositories](https://github.com/Treblle) on Github. Check the contribution guide on the repo you want to contribute to for more details about how to contribute. We're looking forward to your contribution!

### Contributors
<a href="https://github.com/Treblle/treblle/graphs/contributors">
  <p align="center">
    <img  src="https://contrib.rocks/image?repo=Treblle/treblle" alt="A table of avatars from the project's contributors" />
  </p>
</a>
