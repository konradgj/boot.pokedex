package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		lResp := RespLocationAreas{}
		err := json.Unmarshal(val, &lResp)
		if err != nil {
			return RespLocationAreas{}, err
		}

		return lResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	lResp := RespLocationAreas{}
	err = json.Unmarshal(data, &lResp)
	if err != nil {
		return RespLocationAreas{}, err
	}

	c.cache.Add(url, data)
	return lResp, nil
}

func (c *Client) GetLocatioArea(locName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + locName

	if val, ok := c.cache.Get(url); ok {
		lResp := RespLocationArea{}
		err := json.Unmarshal(val, &lResp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return lResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	lResp := RespLocationArea{}
	err = json.Unmarshal(data, &lResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)
	return lResp, nil
}
