package task2

import _ "embed"

/*
 Функции в этом файле реализовывать и менять не нужно, они лишь показывают принцип применения парсера конфига
*/

//go:embed config.yaml
var configYaml []byte

func GetConfig() (*AppConfig, error) {
	conf := &AppConfig{}

	if err := ParseConfig(conf, configYaml); err != nil {
		return nil, err
	}

	return conf, nil
}
