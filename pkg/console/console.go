package console

import "github.com/spf13/cobra"

// Registry maintains a list of commands and provides a fluent interface for registration.
type Registry struct {
	RootCmd *cobra.Command
	SubCmds []*cobra.Command
}

// Builder helps in fluently creating and registering commands.
type Builder struct {
	Command *cobra.Command
}

func Build(use string) *Builder {
	return &Builder{
		Command: &cobra.Command{Use: use},
	}
}

func (b *Builder) Description(desc string) *Builder {
	b.Command.Short = desc
	return b
}

func (b *Builder) Action(action func(cmd *cobra.Command, args []string)) *Builder {
	b.Command.Run = action
	return b
}

func New(rootCmd *cobra.Command) *Registry {
	return &Registry{RootCmd: rootCmd}
}

func (r *Registry) Register(builder *Builder) *Registry {
	r.RootCmd.AddCommand(builder.Command)
	return r
}
