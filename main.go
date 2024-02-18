package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	// "runtime"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
		{
			Name: "help",
			Aliases: []string{"h"},
			Usage: "shows you help stuff.",
			Action: func (*cli.Context) error {
				c := color.New(color.FgGreen, color.Bold)
				fmt.Printf("Need help - ")
				c.Printf("https://to.abku.dev/cli-help")
				return nil
			},
		},
		{
			Name: "create",
			Aliases: []string{"c"},
			Usage: "create a new file with given file-type.",
			Action: func(c *cli.Context) error {
				filename := c.Args().First()
				if filename == "" {
					c := color.New(color.FgRed, color.Bold)
					c.Println("please give a filename");
					return nil
				}

				file, err := os.Create(filename)
				
				if err != nil {
					return err
				}

				defer file.Close()

				success := color.New(color.FgGreen, color.Bold)
				success.Printf("created file %s successfully.\n", filename)
				return nil
			},
		},
		{
			Name: "remove",
			Aliases: []string{"r"},
			Usage: "remove a existing file with file-name.",
			Action: func(c *cli.Context) error {
				filename := c.Args().First()
				if filename == "" {
					c := color.New(color.FgRed, color.Bold)
					c.Println("please give a filename");
					return nil
				}

				err := os.Remove(filename)
				
				if err != nil {
					return err
				}

				success := color.New(color.FgGreen, color.Bold)
				success.Printf("deleted file %s successfully.\n", filename)
				return nil
			},
		},
	},
}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// "C:\Users\Abhishek\abku-cli\abku.exe"