package main

import (
	`fmt`
	`os`
	`runtime`

	"github.com/urfave/cli/v2"

	`goctl-auth-api/action`
)

var (
	version  = "20210706"
	commands = []*cli.Command{
		{
			Name:   "goctl-auth-api",
			Usage:  "generates goctl-auth-api",
			Action: action.AuthApi,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "package",
					Usage: "the package of goctl-auth-api",
				},
			},
		},
	}
)


func main()  {
	app := cli.NewApp()
	app.Usage = "generates goctl-auth-api"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf(" goctl-auth-api: %+v\n", err)
	}
}
