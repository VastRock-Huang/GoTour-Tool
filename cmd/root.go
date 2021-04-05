package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}

//执行命令行参数解析
func Execute() error {
	return rootCmd.Execute()
}
