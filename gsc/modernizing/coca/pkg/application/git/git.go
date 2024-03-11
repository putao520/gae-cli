package git

import (
	"bytes"
	"encoding/json"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/apriori"
	"log"
	"sort"
	"strings"
	"time"
)

var currentCommit CommitMessage
var currentFileChangeMap map[string]FileChange
var commits []CommitMessage
var currentFileChanges []FileChange

func BuildMessageByInput(inputStr string) []CommitMessage {
	currentFileChangeMap = make(map[string]FileChange)
	commits = nil
	currentFileChanges = nil

	splitStr := strings.Split(inputStr, "\n")
	for _, str := range splitStr {
		ParseLog(str)
	}

	return commits
}

func CalculateCodeAge(messages []CommitMessage) []ProjectInfo {
	infos := make(map[string]ProjectInfo)
	BuildCommitMessageMap(messages, infos)

	var agesArray []ProjectInfo
	for _, info := range infos {
		agesArray = append(agesArray, info)
	}

	sort.Slice(agesArray, func(i, j int) bool {
		return agesArray[i].Age.Before(agesArray[j].Age)
	})

	return agesArray
}

func GetTeamSummary(messages []CommitMessage) []TeamSummary {
	infos := make(map[string]ProjectInfo)
	infos = BuildCommitMessageMap(messages, infos)

	var sortInfos []TeamSummary
	for _, info := range infos {
		sortInfos = append(sortInfos, TeamSummary{info.EntityName, len(info.Authors), len(info.Revs)})
	}

	sort.Slice(sortInfos, func(i, j int) bool {
		return sortInfos[i].RevsCount > sortInfos[j].RevsCount
	})

	return sortInfos
}

func BuildCommitMessageMap(messages []CommitMessage, infos map[string]ProjectInfo) map[string]ProjectInfo {
	timeFormat := "2006-01-02"

	for _, commitMessage := range messages {
		for _, change := range commitMessage.Changes {
			fileName := change.File
			if complexMoveReg.MatchString(fileName) {
				infos, fileName = handleMoveInDirectory(infos, fileName)
			} else if basicMvReg.MatchString(fileName) {
				infos, fileName = handleMoveFullPath(infos, fileName)
			}

			if infos[fileName].EntityName == "" {
				authors := make(map[string]string)
				authors[commitMessage.Author] = commitMessage.Author
				revs := make(map[string]string)
				revs[commitMessage.Rev] = commitMessage.Rev

				date, _ := time.Parse(timeFormat, commitMessage.Date)
				infos[fileName] = ProjectInfo{fileName, authors, revs, date}
			} else {
				infos[fileName].Authors[commitMessage.Author] = commitMessage.Author
				infos[fileName].Revs[commitMessage.Rev] = commitMessage.Rev
			}

			if change.Mode == "delete" {
				delete(infos, fileName)
			}
		}
	}
	return infos
}

// 反向查询
func handleMoveInDirectory(infos map[string]ProjectInfo, changedFile string) (map[string]ProjectInfo, string) {
	changedFile, oldFileName, newFileName := UpdateMessageForChange(changedFile)

	if changedFile != oldFileName {
		infos = switchMapFile(infos, oldFileName, newFileName)
	}
	return infos, changedFile
}

func handleMoveFullPath(infos map[string]ProjectInfo, changedFile string) (map[string]ProjectInfo, string) {
	fileName := changedFile
	changed := basicMvReg.FindStringSubmatch(changedFile)

	if len(changed) == 3 {
		infos = switchMapFile(infos, changed[1], changed[2])
		fileName = changed[2]
	}

	return infos, fileName
}

func switchMapFile(infos map[string]ProjectInfo, oldFileName string, newFileName string) map[string]ProjectInfo {
	if _, ok := infos[oldFileName]; ok {
		oldInfo := infos[oldFileName]
		delete(infos, oldFileName)
		oldInfo.EntityName = newFileName
		infos[newFileName] = oldInfo
	}

	return infos
}

type TopAuthor struct {
	Name        string
	CommitCount int
	LineCount   int
}

var GIT_RELATED_MAX_SIZE = 10
var MIN_DATASET = 2

func GetRelatedFiles(commitMessages []CommitMessage, relatedConfig []byte) [][]string {
	var dataset [][]string
	for _, commitMessage := range commitMessages {
		var set []string
		if len(commitMessage.Changes) > GIT_RELATED_MAX_SIZE {
			continue
		}
		for _, change := range commitMessage.Changes {
			if strings.HasSuffix(change.File, ".java") && !strings.HasSuffix(change.File, "Test.java") {
				if strings.Contains(change.File, "core/main/java/") {
					split := strings.Split(change.File, "core/main/java/")
					change.File = strings.ReplaceAll(split[1], "/", ".")
				}

				set = append(set, change.File)
			}
		}

		if len(set) > MIN_DATASET {
			dataset = append(dataset, set)
		}
	}

	var newOptions = apriori.NewOptions(0.1, 0.9, 0, 0)

	decoder := json.NewDecoder(bytes.NewReader(relatedConfig))
	decoder.UseNumber()
	err := decoder.Decode(&newOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	newApriori := apriori.NewApriori(dataset)
	result := newApriori.Calculate(newOptions)

	var availableResults [][]string = nil
	for _, res := range result {
		items := res.GetSupportRecord().GetItems()
		if len(items) > MIN_DATASET {
			availableResults = append(availableResults, items)
		}
	}

	if len(availableResults) > 0 {
		return availableResults
	}

	return nil
}

func GetTopAuthors(commitMessages []CommitMessage) []TopAuthor {
	authors := make(map[string]*TopAuthor)
	for _, commitMessage := range commitMessages {
		if authors[commitMessage.Author] == nil {
			authors[commitMessage.Author] = &TopAuthor{commitMessage.Author, 0, 0}
		}
		authors[commitMessage.Author].CommitCount++
		for _, change := range commitMessage.Changes {
			authors[commitMessage.Author].LineCount = authors[commitMessage.Author].LineCount + change.Added
			authors[commitMessage.Author].LineCount -= change.Deleted
		}
	}

	var topAuthors []TopAuthor
	for _, info := range authors {
		topAuthors = append(topAuthors, TopAuthor{info.Name, info.CommitCount, info.LineCount})
	}

	sort.Slice(topAuthors, func(i, j int) bool {
		return topAuthors[i].CommitCount > topAuthors[j].CommitCount
	})

	return topAuthors
}

func BasicSummary(commitMessages []CommitMessage) *GitSummary {
	authors := make(map[string]string)
	entities := make(map[string]string)
	commits := len(commitMessages)
	changes := 0

	for _, commitMessage := range commitMessages {
		authors[commitMessage.Author] = commitMessage.Author
		for _, change := range commitMessage.Changes {
			entities[change.File] = change.File
			if change.Added > 0 {
				changes++
			}
			if change.Deleted > 0 {
				changes--
			}
		}
	}

	authorSummary := len(authors)
	entitySummary := len(entities)

	gitSummary := &GitSummary{commits, entitySummary, changes, authorSummary}
	return gitSummary
}
