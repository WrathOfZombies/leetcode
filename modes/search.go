package modes

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/wrathofzombies/leetcode/utils"
)

// Lists a prompt that allows searching for problems, either by number of by name
func Search() {
	files := utils.LoadProblems()

	templates := promptui.SelectTemplates{
		Active:   `{{ .Title | cyan | bold }}`,
		Inactive: `{{ .Title | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ "Problem" | bold }}: {{ .Title | cyan }}`,
	}

	list := promptui.Select{
		Label:     "Select a problem to run",
		Items:     files,
		Templates: &templates,
		Searcher: func(input string, idx int) bool {
			problem := files[idx]
			title := strings.ToLower(problem.Title)

			if strings.Contains(title, input) {
				return true
			}

			return false
		},
	}

	idx, data, err := list.Run()
	utils.IsError(err)

	fmt.Println(data, files[idx].Name)
}
