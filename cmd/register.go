/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Ewan-Greer09/st-cli/client"
	"github.com/spf13/cobra"
)

type RegisterAgentResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new 'agent' under an email address",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")

		c := client.New("https://api.spacetraders.io/v2")
		resp, err := c.RegisterAgent(cmd.Flag("faction").Value.String(), cmd.Flag("symbol").Value.String(), cmd.Flag("email").Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		var content RegisterAgentResponse
		err = json.Unmarshal(resp.Body(), &content)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(content.Data.Token)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("faction", "f", "", "faction your agent will be a part of")
	registerCmd.Flags().StringP("symbol", "s", "", "name of your agent")
	registerCmd.Flags().StringP("email", "e", "", "email address associated with the agent")

	registerCmd.MarkFlagRequired("faction")
	registerCmd.MarkFlagRequired("symbol")
	registerCmd.MarkFlagRequired("email")
}
