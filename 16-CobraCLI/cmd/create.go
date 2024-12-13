/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"16-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  runCreate(categoryDb),
	}
}

// createCmd represents the create command
//var createCmd = &cobra.Command{
//	Use:   "create",
//	Short: "A brief description of your command",
//	Long:  `A longer description that spans multiple lines.`,
//
//	RunE: runCreate(*GetCategoryDB(GetDb())),
//}

func runCreate(categoryDb database.Category) RunFuncE {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(*GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	//	createCmd.PersistentFlags().String("name", "", "Nome da categoria")

	createCmd.Flags().StringP("name", "n", "", "Nome da categoria")
	createCmd.Flags().StringP("description", "d", "", "Descrição da categoria")

	//createCmd.MarkFlagRequired("name")
	//createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagsRequiredTogether("name", "description")

}
