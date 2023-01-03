package commit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"regexp"
)

type Commit struct {
	Hash    string
	Date    time.Time
	Subject string
	Author  string
}

func GroupCommitByLabel() map[string]int {
	groupedCommit := map[string]int{}

	commits := getCommits()
	for _, commit := range commits {
		label := regexp.MustCompile(`:.*?:`).FindString(commit.Subject)
		groupedCommit[label]++
	}

	return groupedCommit
}

func getCommits() (commits []Commit) {
	boundary := "=git-commit-stats-boundary="
	format := "%h" + boundary + "%cd" + boundary + "%s" + boundary + "%cn"
	output, err := exec.Command("git", "log", "--no-merges", "--format="+format).Output()
	if err != nil {
		fmt.Println("failed to get git commits!")
		os.Exit(1)
	}

	commitStrings := strings.Split(strings.TrimRight(string(output), "\n"), "\n")
	for i := 0; i < len(commitStrings); i++ {
		commitMetaData := strings.Split(commitStrings[i], boundary)
		date, _ := time.Parse(commitMetaData[1], "2006-01-02T15:04:05+07:00")
		commit := Commit{
			Hash:    commitMetaData[0],
			Date:    date,
			Subject: commitMetaData[2],
			Author:  commitMetaData[3],
		}
		commits = append(commits, commit)
	}

	return
}
