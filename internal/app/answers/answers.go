package answers

type Answer struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
	Key   int         `yaml:"key,omitempty"`
}
