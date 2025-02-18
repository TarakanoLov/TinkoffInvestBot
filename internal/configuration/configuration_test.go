package configuration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	config, err := Load([]byte(`
settings:
  how_often_to_buy: "24h"
  maximum_spend_per_day: 1000
assets:
  shares:
    percent: 0
  bonds:
    percent: 33.3
    buy_in_proportion_to_capitalization: true
    list:
      - any
  etf:
    percent: 33.4
    list:
      - TMOS:
          percent: 33.3
      - SBMX:
          percent: 33.3
      - EQMX:
          percent: 33.4
  metals:
    percent: 33.3
    list:
      - GLDRUB_TOM:
          percent: 50
      - GOLD:
          percent: 50`))
	if err != nil {
		t.Fatalf("Unexpected error in Load: %v", err)
	}
	if config == nil {
		t.Fatalf("Config not parsed")
	}

	assert.Equal(t, config.Settings.How_often_to_buy, time.Hour*24)
	assert.Equal(t, config.Settings.Maximum_spend_per_day, 1000)

	assert.Equal(t, config.Assets.Shares.Percent, 0.0)

	assert.Equal(t, config.Assets.Bonds.Percent, 33.3)
	assert.Equal(t, config.Assets.Bonds.Buy_in_proportion_to_capitalization, true)
	assert.Equal(t, config.Assets.Bonds.List, []string{"any"})

    assert.Equal(t, config.Assets.Etf.Percent, 33.4)

}

func TestLoadErrorIfPercentNotEqualTo100(t *testing.T) {
	config, err := Load([]byte(`
settings:
  how_often_to_buy: "24h"
  maximum_spend_per_day: 1000
assets:
  shares:
    percent: 1
  bonds:
    percent: 33.3
    buy_in_proportion_to_capitalization: true
    list:
      - any
  etf:
    percent: 33.4
    list:
      - TMOS:
          percent: 33.3
      - SBMX:
          percent: 33.3
      - EQMX:
          percent: 33.4
  metals:
    percent: 33.3
    list:
      - GLDRUB_TOM:
          percent: 50
      - GOLD:
          percent: 50`))
	assert.Nil(config)
	assert.Nil(err)
}
/*
etf:
    percent: 33.4
    list:
      - TMOS:
          percent: 33.3
      - SBMX:
          percent: 33.3
      - EQMX:
          percent: 33.4
  metals:
    percent: 33.3
    list:
      - GLDRUB_TOM:
          percent: 50
      - GOLD:
          percent: 50`))*/