package weatherClient

import (
	"encoding/json"
	"io"
	"net/http"
	"weatherAPI/internal/model"
)

const key = "09733c70a5010225ffd555ea2e26a7b1"

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

func (c Client) GetWeather(city string) (model.Weather, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + key + "&units=metric"

	resp, err := c.client.Get(url)
	if err != nil {
		return model.Weather{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Weather{}, err
	}

	var w model.Weather

	err = json.Unmarshal(body, &w)

	return w, nil
}
