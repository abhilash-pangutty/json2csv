package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Secrets struct {
	Line          string      `json:"line"`
	LineNumber    json.Number `json:"lineNumber"`
	Offender      string      `json:"offender"`
	Commit        string      `json:"commit"`
	Repo          string      `json:"repo"`
	RepoURL       string      `json:"repoURL"`
	LeakURL       string      `json:"leakURL"`
	Rule          string      `json:"rule"`
	CommitMessage string      `json:"commitMessage"`
	Author        string      `json:"author"`
	Email         string      `json:"email"`
	File          string      `json:"file"`
	Date          string      `json:"date"`
	Tags          string      `json:"tags"`
	FullURL       string      `json:"full_url"`
}

func main() {

	m := readCurrentDir()
	//**files := readCurrentDir()
	//total_files := len(files)
	//fmt.Println(total_files)
	csvFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	for file, org := range m {
		//fmt.Println("Total Files to Process: ", total_files-i)
		//fmt.Println("Processing File Number :", i)
		jsonData, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}

		var jdata []Secrets
		err = json.Unmarshal([]byte(jsonData), &jdata)
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		writer.Write([]string{"RepoName", "RULE", "OFFENDER", "Date", "URL"})

		for _, items := range jdata {
			var row []string
			row = append(row, items.Repo)
			row = append(row, items.Rule)
			row = append(row, items.Offender)
			row = append(row, items.Date)
			row = append(row, items.Commit)
			row = append(row, ("https://github.com/" + org + "/" + items.Repo + "/blob/" + items.Commit + "/" + items.File))
			writer.Write(row)
		}
		writer.Flush()

	}

}

func readCurrentDir() map[string]string {

	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("Failed Opening Directory: %s", err)
	}
	defer file.Close()
	fileList, _ := file.Readdir(0)
	repoOrg := make(map[string]string)
	for _, files := range fileList {
		if !(files.IsDir()) && (strings.HasSuffix(files.Name(), "json")) {
			org := strings.Split(files.Name(), "__")[0]
			repoOrg[files.Name()] = org
		}
	}
	return repoOrg

}
