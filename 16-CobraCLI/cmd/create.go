/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		db := GetDb()
		defer db.Close()

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		category := GetCategoryDB(db)
		category.Create(name, description)
	},
}

func init() {
	categoryCmd.AddCommand(createCmd)
	//	createCmd.PersistentFlags().String("name", "", "Nome da categoria")

	createCmd.Flags().StringP("name", "n", "", "Nome da categoria")
	createCmd.Flags().StringP("description", "d", "", "Descrição da categoria")

	//createCmd.MarkFlagRequired("name")
	//createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
