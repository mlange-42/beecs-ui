package cli

import (
	"fmt"
	"os"

	game "github.com/mlange-42/beecs-ui/internal"
	"github.com/mlange-42/beecs/params"
	"github.com/spf13/cobra"
)

func Run() {
	cobra.MousetrapHelpText = ""
	if err := rootCommand().Execute(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		fmt.Print("\nRun `beecs-ui -h` for help!\n\n")
		os.Exit(1)
	}
}

// rootCommand sets up the CLI
func rootCommand() *cobra.Command {
	var paramFile string
	var layout string

	root := cobra.Command{
		Use:           "beecs-ui",
		Short:         "beecs-ui provides a graphical user interface for the beecs model.",
		Long:          `beecs-ui provides a graphical user interface for the beecs model.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			p := params.CustomParams{
				Parameters: params.Default(),
			}
			if paramFile != "" {
				err := p.FromJSONFile(paramFile)
				if err != nil {
					return err
				}
			}

			game.Run(layout, paramFile, 4)

			return nil
		},
	}

	root.Flags().StringVarP(&paramFile, "parameters", "p", "",
		"Parameter file")

	root.Flags().StringVarP(&layout, "layout", "l", "default",
		"Layout or layout file")

	root.Flags().SortFlags = false

	return &root
}
