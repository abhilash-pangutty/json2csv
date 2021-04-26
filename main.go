package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Secrets struct {
	Line          string `json:"line"`
	LineNumber    string `json:"lineNumber"`
	Offender      string `json:"offender"`
	Commit        string `json:"commit"`
	Repo          string `json:"repo"`
	RepoURL       string `json:"repoURL"`
	LeakURL       string `json:"leakURL"`
	Rule          string `json:"rule"`
	CommitMessage string `json:"commitMessage"`
	Author        string `json:"author"`
	Email         string `json:"email"`
	File          string `json:"file"`
	Date          string `json:"date"`
	Tags          string `json:"tags"`
	FullURL       string `json:"full_url"`
}

func main() {

	var fileName string
	var csvFileName string

	fileName = os.Args[1]
	csvFileName = os.Args[2]

	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var jdata []Secrets
	err = json.Unmarshal([]byte(jsonData), &jdata)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create(csvFileName)

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, usance := range jdata {
		var row []string
		row = append(row, usance.Line)
		row = append(row, usance.LineNumber)
		row = append(row, usance.Offender)
		row = append(row, usance.Commit)
		row = append(row, usance.Repo)
		row = append(row, usance.RepoURL)
		row = append(row, usance.LeakURL)
		row = append(row, usance.Rule)
		row = append(row, usance.CommitMessage)
		row = append(row, usance.Author)
		row = append(row, usance.Email)
		row = append(row, usance.File)
		row = append(row, usance.Date)
		row = append(row, usance.Tags)
		row = append(row, usance.FullURL)
		writer.Write(row)
	}

	writer.Flush()

}
