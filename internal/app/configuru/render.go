package configuru

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/42nerds/configuru/internal/app/answers"
	"gitlab.com/42nerds/configuru/internal/app/files"
)

type RenderConfiguration struct {
	Files     []RenderConfigurationFiles
	Answers   map[string]answers.Answer
	DocString string
}

type RenderConfigurationFiles struct {
	Name         string
	TemplatePath string
}

func (renderConfig *RenderConfiguration) Fill(configuration Configuration) error {
	answers := make(map[string]answers.Answer)
	for _, answer := range configuration.Answers {
		answers[answer.Name] = answer
	}
	renderConfig.Answers = answers
	renderConfig.DocString = configuration.DocString
	return nil
}

func (renderConfig *RenderConfiguration) AddTemplates(path string) error {
	templateDir, err := files.GetDir(path)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name() == "configuru.yaml" {
			continue
		}
		renderConfig.Files = append(renderConfig.Files, RenderConfigurationFiles{Name: file.Name(), TemplatePath: templateDir + file.Name()})
	}
	return nil
}

func (renderConfig *RenderConfiguration) Execute(path string) error {
	targetDir, err := files.GetDir(path)
	if err != nil {
		return err
	}
	for _, file := range renderConfig.Files {
		t, err := template.ParseFiles(file.TemplatePath)
		if err != nil {
			return err
		}
		fileHandler, err := os.OpenFile(targetDir+file.Name, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer fileHandler.Close()
		// log.Println(renderConfig.Answers)
		comment := createComment(file.Name, renderConfig.DocString)
		if len(comment) > 0 {
			_, err = fileHandler.WriteString(comment)
			if err != nil {
				return err
			}
		}
		err = t.Execute(fileHandler, renderConfig)
		if err != nil {
			return err
		}

	}
	return nil
}

func createComment(file string, message string) string {
	if len(message) == 0 {
		return ""
	}
	if strings.HasSuffix(file, ".py") {
		return "# " + message + "\n"
	}
	if strings.HasSuffix(file, ".yaml") {
		return "# " + message + "\n"
	}
	if strings.HasSuffix(file, ".json") {
		return "// " + message + "\n"
	}
	if strings.HasSuffix(file, ".xml") {
		return "<!-- " + message + " -->\n"
	}
	return ""
}
