package exec

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func GetConf() {
	conf  := Conf{
		Hits:11,
		Time: 43,
	}

	out, _ := yaml.Marshal(conf)
	fmt.Println(string(out))
}