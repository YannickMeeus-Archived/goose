package cmd

import (
	"github.com/Silly-Goose-Software/goose/pkg/cmd/fun"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: `A collection of automation steps just to make the author's life a bit easier.
I would be surprised if this is going to be useful to a lot more people than just myself,
but we'll see. This CLI will help me create a bunch of ad hoc things that help me in my day to day
development tasks.`,
	Example: `goose github personal-access-token create
goose s3 bucket create`,
	Version: "v0.0.3", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	// Your code here
	//
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		pcli.CheckForUpdates()
		os.Exit(1)
	}

	pcli.CheckForUpdates()
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("Silly-Goose-Software/goose")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()
	rootCmd.AddCommand(fun.Cmd)

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)

	rand.Seed(time.Now().UnixNano())
}
