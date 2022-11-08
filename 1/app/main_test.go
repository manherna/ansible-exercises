package main

import "testing"

func TestGetWeatherFail(t *testing.T) {
	want := "There was a problem fetching results. HTTP Staus code: 401"
	if got := getWeather("badapikey", "Honolulu"); got != want {
		t.Errorf("getWeather expected %s\n, got %s", want, got)
	}
}
