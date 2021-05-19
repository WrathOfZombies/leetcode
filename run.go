package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
)

type Problem struct {
	Name       string
	Path       string
	Solution   string
	Validation string
}

func isError(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
	panic(err)
}

const PROBLEMS_DIR = "./problems"

func constructProblem(file fs.FileInfo) Problem {
	var name = file.Name()
	var basePath = path.Join(PROBLEMS_DIR, name)

	return Problem{
		Name:       name,
		Path:       basePath,
		Solution:   path.Join(basePath, "solution.go"),
		Validation: path.Join(basePath, "validation.yml"),
	}
}

func loadFiles() []Problem {
	results := []Problem{}

	folder, err := os.Open(PROBLEMS_DIR)
	isError(err)

	files, err := folder.Readdir(-1)
	folder.Close()
	isError(err)

	for _, file := range files {
		results = append(results, constructProblem(file))
	}

	return results
}

func promptSearch() {
	files := loadFiles()

	templates := promptui.SelectTemplates{
		Active:   `🍕 {{ .Name | cyan | bold }}`,
		Inactive: `   {{ .Name | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ "Problem" | bold }}: {{ .Name | cyan }}`,
	}

	list := promptui.Select{
		Label:     "Problem",
		Items:     files,
		Templates: &templates,
		Searcher: func(input string, idx int) bool {
			problem := files[idx]
			title := strings.ToLower(problem.Name)

			if strings.Contains(title, input) {
				return true
			}

			return false
		},
	}

	idx, data, err := list.Run()
	isError(err)

	fmt.Println(data, files[idx].Name)
}

func main() {
	promptSearch()
}
