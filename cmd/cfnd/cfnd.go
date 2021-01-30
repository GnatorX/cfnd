package main

import (
	"fmt"
	"log"
	"os"

	"cloudformation-error/pkg/ctl"

	"github.com/urfave/cli/v2"
)

func main() {
	var region string
	var stackName string
	app := &cli.App{
		Name:  "cfnd",
		Usage: "Find Cloudtrail events for your CloudFormation errors",
		Action: func(c *cli.Context) error {
			fmt.Println(region)
			fmt.Println(stackName)
			fmt.Println(ctl.Find(c.Context, stackName, region))
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "stackname",
				Usage:       "Name of the stack",
				Destination: &stackName,
				Aliases:     []string{"s"},
			},
			&cli.StringFlag{
				Name:        "region",
				Usage:       "AWS region for the search",
				Destination: &region,
				Aliases:     []string{"r"},
				DefaultText: "us-west-2",
				Value:       "us-west-2",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}