package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "This command is use to create vanilla ",
	RunE: func(cmd *cobra.Command, args []string) error {
		dirs := []string{"style", "script", "assest"}

		for _, dir := range dirs {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Printf("Error ceating dir %s: %v\n", dir, err)
				return err
			}
			fmt.Printf("Directory created succesfully: %s\n", dir)
		}

		return nil
	},
}
