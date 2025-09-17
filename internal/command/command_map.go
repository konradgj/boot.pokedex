package command

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	lResp, err := cfg.PokeapiClient.GetLocationArea(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = lResp.Next
	cfg.prevLocationUrl = lResp.Previous

	for _, loc := range lResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	lResp, err := cfg.PokeapiClient.GetLocationArea(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = lResp.Next
	cfg.prevLocationUrl = lResp.Previous

	for _, loc := range lResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
