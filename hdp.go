package main

import (
	"os"
	"os/user"
	"net"
	"strconv"
	"strings"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
	"github.com/olekukonko/tablewriter"
	"github.com/vincer/libhdplatinum"
)

func validate(ip string, port int, c *cli.Context) *cli.ExitError {
	if strings.TrimSpace(ip) == "" {
		return cli.NewExitError("IP address is required. `hdp -h` for usage info.", 1)
	}
	if net.ParseIP(ip) == nil {
		return cli.NewExitError("Invalid IP address", 1)
	}
	if port < 0 || port > 65535 {
		return cli.NewExitError("Invalid port", 2)
	}
	return nil
}

func getUserHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func main() {
	app := cli.NewApp()
	app.Name = "hdp"
	app.Usage = "Hunter Douglas Platinum CLI"
	app.Version = "0.0.1"
	app.HideHelp = true
	app.EnableBashCompletion = true

	var ip string
	var port int

	home := getUserHome()
	flags := []cli.Flag{
		altsrc.NewStringFlag(cli.StringFlag{
			Name: "ip",
			Usage: "ip address of the Hunter Douglas Platinum Gateway. Required.",
			Destination: &ip,
		}),
		cli.IntFlag{
			Name: "port",
			Value: 522,
			Usage: "port of the Hunter Douglas Platinum Gateway.",
			Destination: &port,
		},
		cli.StringFlag{
			Name: "config, c",
			Value: home + "/.hdp.yml",
			Usage: "YAML config file. Useful for saving IP address of gateway.",
		},
		cli.BoolFlag{
			Name: "help, h",
			Usage: "Show usage help.",
		},
	}

	app.Before = func(c *cli.Context) error {
		filePath := c.String("config")
		if _, err := os.Stat(filePath); err == nil {
			return altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))(c)
		}
		return nil
	}

	app.Flags = flags
	app.Commands = []cli.Command{
		{
			Name: "shades",
			Usage: "Shade commands",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Aliases: []string{"ls"},
					Usage: "List shades",
					Action: func(c *cli.Context) error {
						exitError := validate(ip, port, c)
						if (exitError != nil) {
							return exitError
						}
						shades := libhdplatinum.GetShades(ip, port)
						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"Name", "Height", "id"})
						for _, shade := range shades {
							table.Append([]string{shade.Name(), strconv.Itoa(shade.Height()), shade.Id()})
						}
						table.SetBorder(false)
						table.SetAutoFormatHeaders(false)
						table.SetCenterSeparator("-")
						table.Render()

						return nil
					},
				},
				{
					Name: "set",
					Usage: "set <shade id or name> <height 0-255>",
					Action: func(c *cli.Context) {
						shades := libhdplatinum.GetShades(ip, port)
						shadeIdOrName := c.Args().First()
						for _, shade := range shades {
							if shadeIdOrName == shade.Id() || shadeIdOrName == shade.Name() {
								height, _ := strconv.Atoi(c.Args().Get(1))
								shade.SetHeight(height)
							}
						}
					},
				},
			},
		},
		{
			Name: "rooms",
			Usage: "Room commands",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Aliases: []string{"ls"},
					Usage: "List rooms",
					Action: func(c *cli.Context) {
						rooms := libhdplatinum.GetRooms(ip, port)
						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"Name", "Number of Shades", "id"})
						for _, room := range rooms {
							table.Append([]string{room.Name(), strconv.Itoa(len(room.Shades())), room.Id()})
						}
						table.SetBorder(false)
						table.SetAutoFormatHeaders(false)
						table.SetCenterSeparator("-")
						table.Render()
					},
				},
				{
					Name: "set",
					Usage: "set <room id or name> <height 0-255>",
					Action: func(c *cli.Context) {
						rooms := libhdplatinum.GetRooms(ip, port)
						roomIdOrName := c.Args().First()
						for _, room := range rooms {
							if roomIdOrName == room.Id() || roomIdOrName == room.Name() {
								height, _ := strconv.Atoi(c.Args().Get(1))
								for _, shade := range room.Shades() {
									shade.SetHeight(height)
								}
							}
						}
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
