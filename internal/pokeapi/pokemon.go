package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pResp := Pokemon{}
		err := json.Unmarshal(val, &pResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pResp := Pokemon{}
	err = json.Unmarshal(data, &pResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pResp, nil
}
