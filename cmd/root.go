package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gencli [filename | dirname]",
		Short: "to make cli tool ",
		Long:  `This is long form of cli desc`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error { // here error is specified because we are using RunE imp to specify it return error
			if len(args) == 0 {
				return cmd.Help()
			}
			extfile := args[0]

			if _, err := os.Stat(extfile); os.IsNotExist(err) {
				if ext := filepath.Ext(extfile); ext != "" { // to check wheather args has filename with extension of the file or not
					file, err := os.Create(extfile) // if then create the file   or
					if err != nil {
						fmt.Println("Erorr creating file ", err)
					}
					defer file.Close()
					fmt.Printf(" %s has been created succesfully\n", extfile)
				} else {
					err := os.Mkdir(extfile, 0755) // ceate the dir but it has neither then
					if err != nil {
						fmt.Println("error creating dir", err)
					}
				}

			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(webCmd)
	// rootCmd.DisableFlagParsing = true // Add this line
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if err.Error() == "accepts 1 arg(s), received 0" {
			fmt.Println(rootCmd.Long)
			return
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
