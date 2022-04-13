package configuru

import (
	"io/ioutil"
	"os"

	"gitlab.com/42nerds/configuru/internal/app/files"
	"gitlab.com/42nerds/configuru/internal/app/questions"
	"gopkg.in/yaml.v2"
)

func MarshalconfiguruFile(path string, configuration Configuration) error {
	targetDir, err := files.GetDir(path)
	if err != nil {
		return err
	}

	configuration.dropQuestions()

	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return err
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err := os.Mkdir(targetDir, 0755)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(targetDir+"configuru.yaml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (configuration *Configuration) dropQuestions() {
	configuration.Questions = make([]questions.Question, 0)
}
