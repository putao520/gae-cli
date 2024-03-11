package cmd

import (
	"bufio"
	"encoding/json"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/cmd/config"
	"gae-cli/gsc/modernizing/coca/pkg/application/arch"
	"gae-cli/gsc/modernizing/coca/pkg/application/arch/tequila"
	"gae-cli/gsc/modernizing/coca/pkg/application/visual"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type ArchCmdConfig struct {
	DependencePath string
	IsMergePackage bool
	FilterString   string
	IsMergeHeader  bool
	WithVisual     bool
}

var (
	archCmdConfig ArchCmdConfig
)

var archCmd = &cobra.Command{
	Use:   "arch",
	Short: "project package visualization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		identifiers = cmd_util.LoadIdentify(apiCmdConfig.DependencePath)
		identifiersMap = core_domain.BuildIdentifierMap(identifiers)

		parsedDeps := cmd_util.GetDepsFromJson(archCmdConfig.DependencePath)
		archApp := arch.NewArchApp()
		result := archApp.Analysis(parsedDeps, identifiersMap)

		filter := strings.Split(archCmdConfig.FilterString, ",")
		var nodeFilter = func(key string) bool {
			for _, f := range filter {
				if strings.Contains(key, f) {
					return true
				}
			}
			return false
		}

		if archCmdConfig.WithVisual {
			output := visual.FromDeps(parsedDeps)
			out, _ := json.Marshal(output)
			cmd_util.WriteToCocaFile("visual.json", string(out))
		}

		if archCmdConfig.IsMergeHeader {
			result = result.MergeHeaderFile(tequila.MergeHeaderFunc)
		}

		if archCmdConfig.IsMergePackage {
			result = result.MergeHeaderFile(tequila.MergePackageFunc)
		}

		graph := result.ToMapDot(nodeFilter)
		f, _ := os.Create("coca_reporter/arch.dot")
		w := bufio.NewWriter(f)
		_, _ = w.WriteString("di" + graph.String())
		_ = w.Flush()

		cmd_util.ConvertToSvg("arch")
	},
}

func init() {
	rootCmd.AddCommand(archCmd)

	archCmd.PersistentFlags().StringVarP(&archCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	archCmd.PersistentFlags().BoolVarP(&archCmdConfig.IsMergePackage, "mergePackage", "P", false, "merge package")
	archCmd.PersistentFlags().BoolVarP(&archCmdConfig.IsMergeHeader, "mergeHeader", "H", false, "merge header")
	archCmd.PersistentFlags().BoolVarP(&archCmdConfig.WithVisual, "showVisual", "v", false, "build visual json")
	archCmd.PersistentFlags().StringVarP(&archCmdConfig.FilterString, "filter", "x", "", "filter -x com.phodal")
}
