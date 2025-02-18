package configuration

import (
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Settings struct {
		How_often_to_buy      time.Duration `yaml:"how_often_to_buy"`
		Maximum_spend_per_day int           `yaml:"maximum_spend_per_day"`
	} `yaml:"settings"`
	Assets struct {
		Shares struct {
			Percent float64 `yaml:"percent"`
		} `yaml:"shares"`
		Bonds struct {
			Percent                             float64  `yaml:"percent"`
			Buy_in_proportion_to_capitalization bool     `yaml:"buy_in_proportion_to_capitalization"`
			List                                []string `yaml:"list"`
		} `yaml:"bonds"`
		Etf struct {
			Percent float64 `yaml:"percent"`
			//List []struct
		} `yaml:"etf"`
		Metals struct {
			Percent float64 `yaml:"percent"`
		}
	} `yaml:"assets"`
}

func Load(input []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(input, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
