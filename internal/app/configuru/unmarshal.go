package configuru

import (
	"io/ioutil"

	"gitlab.com/42nerds/configuru/internal/app/files"
	"gopkg.in/yaml.v2"
)

func UnmarshalconfiguruFile(path string) (Configuration, error) {
	workdir, err := files.GetDir(path)
	if err != nil {
		return Configuration{}, err
	}

	f, err := ioutil.ReadFile(workdir + "configuru.yaml")
	if err != nil {
		return Configuration{}, err
	}

	configuration := Configuration{Source: workdir}
	err = yaml.Unmarshal(f, &configuration)
	if err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}
