package cmd

import (
	"github.com/SCKelemen/chexmix/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans directory or files for errors",
	Long: `🔍 scan:
accepts a file or directory and will scan the files 
for errors according to the activators for the extension of the files.`,
	Run: scanFunc,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scanFunc(cmd *cobra.Command, args []string) {
	scan := scanner.New("")
	scan.Scan()
}
