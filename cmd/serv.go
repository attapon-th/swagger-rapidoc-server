/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/attapon-th/swagger-rapidoc/internal/route"
	"github.com/attapon-th/swagger-rapidoc/internal/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("serv called")
		fconf := setup.NewFiberConfig()
		app := fiber.New(fconf)
		route.New(app)
		setup.Listen(app)
	},
}

func init() {
	rootCmd.AddCommand(servCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
