package cmd

import (
	"encoding/json"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/cmd/config"
	"gae-cli/gsc/modernizing/coca/pkg/application/suggest"
	"github.com/spf13/cobra"
	"log"
)

var (
	suggestConfig ApiCmdConfig
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "find usable Design Patterns from code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		parsedDeps = nil
		depFile := cmd_util.ReadFile(apiCmdConfig.DependencePath)
		if depFile == nil {
			log.Fatal("lost deps")
		}

		_ = json.Unmarshal(depFile, &parsedDeps)

		app := suggest.NewSuggestApp()
		results := app.AnalysisPath(parsedDeps)

		table := cmd_util.NewOutput(output)
		table.SetHeader([]string{"Class", "Pattern", "Reason"})

		for _, result := range results {
			table.Append([]string{result.Class, result.Pattern, result.Reason})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)

	suggestCmd.PersistentFlags().StringVarP(&suggestConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence D")
}
