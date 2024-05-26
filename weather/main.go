package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color" // imports color for CLI messages
)

type Weather struct {
	Location struct {
		City  string `json:"name"`
		State string `json:"region"`
	} `json:"location"`
	Current struct {
		TempF     float64 `json:"temp_f"`
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Condition string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempF     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
} // access each JSON item needed from API call

func main() {
	q := "11220"
	if len(os.Args) > 1 {
		q = os.Args[1]
	} // check args; default is 11220 zip code

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=6b143505c7144abdb8a164128242605&q=" + q + "&days=1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	} // error handling for api key
	body, err := io.ReadAll(res.Body) // read response from api call
	if err != nil {
		panic(err)
	}
	var weather Weather // declare weather var type Weather
	err = json.Unmarshal(body, &weather) // parse JSON string -> data structure
	if err != nil {
		panic(err)
	}
	name, country, temp, condition, hours :=
		weather.Location.City,
		weather.Location.State,
		weather.Current.TempF,
		weather.Current.Condition.Condition,
		weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s: %.0fF, %s\n", name, country, temp, condition) // print current weather

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0) // begins time at 00:00

		if date.Before(time.Now()) {
			continue // skips all the time that has already passed
		}

		message := fmt.Sprintf("%s - %.0fF, %.0f, %s\n",
			date.Format("03:04PM"), // for military time: date.Format("15:00")
			hour.TempF,
			hour.ChanceOfRain,
			hour.Condition.Text,
		) // print forecast

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		} // handles hours where chance of rain is high. highlights it red instead of default black
	}
}
