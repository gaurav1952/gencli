package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "This command is use to create vanilla web struct",
	RunE: func(cmd *cobra.Command, args []string) error {
		dirs := []string{"styles", "scripts", "assests"}
		files := map[string]string{
			"index.html":        "index.html",
			"styles/style.css":  "styles/style.css",
			"scripts/script.js": "scripts/script.js",
		}

		for _, dir := range dirs {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Printf("Error ceating dir %s: %v\n", dir, err)
				return err
			}
			fmt.Printf("Directory created succesfully: %s\n", dir)
		}

		for fileName, filePath := range files {
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error crate file %s: %v\n ", fileName, err)
			}
			defer file.Close()
		}

		return nil
	},
}
