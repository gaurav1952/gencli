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
		RunE: func(cmd *cobra.Command, args []string) error { // here error is specified because we are using RunE imp to specify it return error
			if len(args) == 0 {
				return cmd.Help()
			}
			extfile := args[0]
			for _, subCmd := range cmd.Commands() {
				if subCmd.Name() == extfile {
					return subCmd.Execute()
				}
			}
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
			return nil
			////

		},
	}
)

func init() {
	rootCmd.AddCommand(webCmd)
}

// Execute executes the root command.
func Execute() {

	// rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
