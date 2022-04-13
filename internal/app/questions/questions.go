package questions

import (
	"errors"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
	"gitlab.com/42nerds/configuru/internal/app/answers"
)

type Question struct {
	Name        string              `yaml:"name"`
	Message     string              `yaml:"message"`
	Type        QuestionType        `yaml:"type"`
	Selectables []string            `yaml:"selectables,omitempty"`
	When        []QuestionCondition `yaml:"when"`
}

type QuestionCondition struct {
	Name       string `yaml:"name"`
	Is         bool   `yaml:"is"`
	IsNot      bool   `yaml:"isNot"`
	IsEqual    string `yaml:"isEqual"`
	IsNotEqual string `yaml:"isNotEqual"`
}

type QuestionType string

const (
	QuestionTypeString  QuestionType = "string"
	QuestionTypeBoolean QuestionType = "boolean"
	QuestionTypeSelect  QuestionType = "select"
)

func PromptQuestions(questions []Question) ([]answers.Answer, error) {
	var answers []answers.Answer
	for _, question := range questions {
		if question.CheckCondtion(answers) {
			answer, err := question.PromptQuestion()
			if err != nil {
				return nil, err
			}
			answers = append(answers, answer)

		}
	}
	return answers, nil
}

var floatValidator = func(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Expected Float. Cannot convert input: " + input)
	}
	return nil
}

func (question *Question) PromptQuestion() (answers.Answer, error) {
	switch question.Type {
	case QuestionTypeString:
		prompt := promptui.Prompt{
			Label: question.Message,
		}

		result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		return answers.Answer{Name: question.Name, Value: result}, nil
	case QuestionTypeSelect:
		prompt := promptui.Select{
			Label:             question.Message,
			Items:             question.Selectables,
			StartInSearchMode: true,
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		return answers.Answer{Name: question.Name, Value: result}, nil
	case QuestionTypeBoolean:
		prompt := promptui.Select{
			Label: question.Message,
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		if result == "Yes" {
			return answers.Answer{Name: question.Name, Value: true}, nil
		} else {
			return answers.Answer{Name: question.Name, Value: false}, nil
		}
	}
	return answers.Answer{}, nil
}

func (question *Question) CheckCondtion(previousAnswers []answers.Answer) bool {
	if len(question.When) == 0 {
		return true
	}
	met := false
	for _, condition := range question.When {
		if condition.Is == getAnswerValue(condition.Name, previousAnswers) {
			met = true
			continue
		} else if condition.IsNot != getAnswerValue(condition.Name, previousAnswers) {
			met = true
			continue
		} else if condition.IsEqual == getAnswerValue(condition.Name, previousAnswers) && len(condition.IsNotEqual) > 0 {
			met = true
			continue
		} else if condition.IsNotEqual != getAnswerValue(condition.Name, previousAnswers) && len(condition.IsNotEqual) > 0 {
			log.Println(condition.IsNotEqual)
			log.Println(getAnswerValue(condition.Name, previousAnswers))
			met = true
			continue
		}
	}
	return met
}

func getAnswerValue(questionName string, previousAnswers []answers.Answer) interface{} {
	for _, answer := range previousAnswers {
		if answer.Name == questionName {
			return answer.Value
		}
	}
	return nil
}
