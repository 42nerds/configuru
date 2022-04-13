package configuru

import (
	"gitlab.com/42nerds/configuru/internal/app/answers"
	"gitlab.com/42nerds/configuru/internal/app/questions"
)

type Configuration struct {
	Source    string               `yaml:"source"`
	Version   string               `yaml:"version"`
	DocString string               `yaml:"docString"`
	Questions []questions.Question `yaml:"questions,omitempty"`
	Answers   []answers.Answer     `yaml:"answers,omitempty"`
}
