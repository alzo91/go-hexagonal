/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alzo91/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "This command allows you to interact with the application using the command line",
	Long: `This command allows you to interact with the application using the command line for handling subjects about the entity product.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Action to be executed: create, enable, disable or get")
	cliCmd.Flags().StringVarP(&productId, "productId", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "productName", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "productPrice", "p", 0.0, "Product Price")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
