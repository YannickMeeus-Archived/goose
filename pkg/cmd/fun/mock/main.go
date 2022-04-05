package fun

import (
	"github.com/spf13/cobra"
)

var FunSubCommand = &cobra.Command{
	Use:   "fun",
	Short: "A collection of useless commands.",
	Long:  `Sometimes also contains sandboxes useful commands that need a bit more polish.`,
}

func init() {

}
