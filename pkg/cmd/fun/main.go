package fun

import (
	"github.com/Silly-Goose-Software/goose/pkg/cmd/fun/mock"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "fun",
	Short: "A collection of useless commands.",
	Long:  `Sometimes also contains sandboxes useful commands that need a bit more polish.`,
}

func init() {
	Cmd.AddCommand(mock.Cmd)
}
