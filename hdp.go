package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/vincer/libhdplatinum"
	"github.com/olekukonko/tablewriter"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "hdp"
	app.Usage = "Hunter Douglas Platinum CLI"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "shades",
			Usage: "Shade commands",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Aliases: []string{"ls"},
					Usage: "List shades",
					Action: func(c *cli.Context) {
						shades := libhdplatinum.GetShades()
						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"Name", "Height", "id"})
						for _, shade := range shades {
							table.Append([]string{shade.Name(), strconv.Itoa(shade.Height()), shade.Id()})
						}
						table.SetBorder(false)
						table.SetAutoFormatHeaders(false)
						table.SetCenterSeparator("-")
						table.Render()
					},
				},
				{
					Name: "set",
					Usage: "set <shade id or name> <height 0-255>",
					Action: func(c *cli.Context) {
						shades := libhdplatinum.GetShades()
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
						rooms := libhdplatinum.GetRooms()
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
						rooms := libhdplatinum.GetRooms()
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
