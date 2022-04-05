package mock

import (
	"fmt"
	"github.com/Silly-Goose-Software/goose/pkg/mock_text_generator"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "mock [string]",
	Short: "Prints out a string AlL FuNNy LiKe",
	Long:  `Don't judge me, I do this for fun, alright?'`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toConvert := args[0]
		converted := mock_text_generator.ToMockText(toConvert)
		fmt.Println(converted)
	},
}

func init() {

}
