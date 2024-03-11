package cmd

import (
	"encoding/json"
	"fmt"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	. "gae-cli/gsc/modernizing/coca/pkg/application/git"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"time"
)

type GitCmdConfig struct {
	Size        int
	ShowSummary bool
}

var (
	relatedConfig string
	gitCmdConfig  GitCmdConfig
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "analysis git commit history for revs count, summary and suggest",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		message := getCommitMessage()
		commitMessages := BuildMessageByInput(message)
		cModel, _ := json.MarshalIndent(commitMessages, "", "\t")
		cmd_util.WriteToCocaFile("commits.json", string(cModel))

		if gitCmdConfig.ShowSummary {
			ShowChangeLogSummary(commitMessages, output)
		}

		isFullMessage := cmd.Flag("full").Value.String() == "true"
		size := gitCmdConfig.Size

		table := cmd_util.NewOutput(output)

		if cmd.Flag("basic").Value.String() == "true" {
			basicSummary := BasicSummary(commitMessages)
			table.SetHeader([]string{"Statistic", "Number"})
			table.Append([]string{"Commits", strconv.Itoa(basicSummary.Commits)})
			table.Append([]string{"Entities", strconv.Itoa(basicSummary.Entities)})
			table.Append([]string{"Changes", strconv.Itoa(basicSummary.Changes)})
			table.Append([]string{"Authors", strconv.Itoa(basicSummary.Authors)})
			table.Render()
		}

		if cmd.Flag("team").Value.String() == "true" {
			teamSummary := GetTeamSummary(commitMessages)
			table.SetHeader([]string{"EntityName", "RevsCount", "AuthorCount"})

			if len(teamSummary) > size && isFullMessage {
				teamSummary = teamSummary[:size]
			}
			for _, v := range teamSummary {
				table.Append([]string{v.EntityName, strconv.Itoa(v.RevsCount), strconv.Itoa(v.AuthorCount)})
			}
			table.Render()
		}

		if cmd.Flag("age").Value.String() == "true" {
			ages := CalculateCodeAge(commitMessages)
			var agesDisplay []CodeAgeDisplay
			for _, info := range ages {
				const secondsOfOneMonth = 2600640
				month := time.Since(info.Age).Seconds() / secondsOfOneMonth
				displayMonth := strconv.FormatFloat(month, 'f', 2, 64)
				agesDisplay = append(agesDisplay, CodeAgeDisplay{EntityName: info.EntityName, Month: displayMonth})
			}

			table.SetHeader([]string{"EntityName", "Month"})

			if len(agesDisplay) > size && isFullMessage {
				agesDisplay = agesDisplay[:size]
			}
			for _, v := range agesDisplay {
				table.Append([]string{v.EntityName, v.Month})
			}
			table.Render()
		}

		if cmd.Flag("top").Value.String() == "true" {
			authors := GetTopAuthors(commitMessages)
			table.SetHeader([]string{"Author", "CommitCount", "LineCount"})

			if len(authors) > size && isFullMessage {
				authors = authors[:size]
			}
			for _, v := range authors {
				table.Append([]string{v.Name, strconv.Itoa(v.CommitCount), strconv.Itoa(v.LineCount)})
			}
			table.Render()
		}

		if relatedConfig != "" {
			config, err := ioutil.ReadFile(relatedConfig)
			if err != nil {
				_ = fmt.Errorf("lost related json %s", err)
				return
			}

			GetRelatedFiles(commitMessages, config)
		}
	},
}

func getCommitMessage() string {
	historyArgs := []string{"log", "--pretty=\"format:[%h] %aN %ad %s\"", "--date=short", "--numstat", "--reverse", "--summary"}
	cmd := exec.Command("git", historyArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalf("Cmd.Run() failed with %s\n", err)
	}

	return string(out)
}

func init() {
	rootCmd.AddCommand(gitCmd)

	gitCmd.PersistentFlags().BoolP("basic", "b", false, "Basic Summary")
	gitCmd.PersistentFlags().BoolP("team", "t", false, "Team Summary")
	gitCmd.PersistentFlags().BoolP("age", "a", false, "Code Age")
	gitCmd.PersistentFlags().BoolP("top", "o", false, "Top Authors")
	gitCmd.PersistentFlags().BoolP("full", "f", false, "full")
	gitCmd.PersistentFlags().BoolVarP(&gitCmdConfig.ShowSummary, "summary", "m", false, "full")
	gitCmd.PersistentFlags().IntVarP(&gitCmdConfig.Size, "size", "s", 20, "full")
	gitCmd.PersistentFlags().StringVarP(&relatedConfig, "related", "r", "", "related")
}
