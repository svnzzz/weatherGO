package main

import (
	"net/url"
	"strings"
)

func buildWeatherURL(city, apiKey string) (string, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.openweathermap.org",
		Path:   "/data/2.5/weather",
	}
	q := url.Values{}
	q.Set("q", strings.TrimSpace(city))
	q.Set("appid", apiKey)
	q.Set("units", "metric")
	q.Set("lang", "it")
	u.RawQuery = q.Encode()
	return u.String(), nil
}
