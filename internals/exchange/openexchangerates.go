package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenexchangeratesExchange struct {
	appID        string
	baseCurrency string
	baseURL      string
}

var _ Exchange = (*OpenexchangeratesExchange)(nil)

func NewOpenexchangeratesExchange(appID string) Exchange {
	return &OpenexchangeratesExchange{
		appID:        appID,
		baseCurrency: "USD", // Free plan of openexchangerates.org only offer USD as a base currency.
		baseURL:      "https://openexchangerates.org",
	}
}

func (e *OpenexchangeratesExchange) latest() (map[string]json.Number, error) {
	url := fmt.Sprintf("%s/api/latest.json?app_id=%s", e.baseURL, e.appID)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: cannot GET %s; %s\n", url, err)
		return nil, err
	}
	defer res.Body.Close()

	type Message struct {
		Disclaimer string                 `json:"disclaimer"`
		License    string                 `json:"license"`
		Timestamp  int                    `json:"timestamp"`
		Base       string                 `json:"base"`
		Rates      map[string]json.Number `json:"rates"`
	}

	var m Message
	dec := json.NewDecoder(res.Body)
	dec.UseNumber()
	err = dec.Decode(&m)
	if err != nil {
		return nil, err
	}

	return m.Rates, nil
}

func (e *OpenexchangeratesExchange) Convert(amt float64, from, to string) float64 {
	fmt.Println("Convert")
	rates, err := e.latest()
	if err != nil {
		fmt.Printf("error: cannot convert; %s\n", err.Error())
		return 0.0
	}
	// for targetCurrency, rate := range rates {
	// 	fmt.Printf("%s -> %s: %s\n", e.baseCurrency, targetCurrency, rate.String())
	// }

	rateFrom, err := rates[from].Float64()
	if err != nil {
		panic(fmt.Sprintf("error: unable to extract and convert rate for currency \"%s\"; %s\n", from, err))
	}
	rateTo, err := rates[to].Float64()
	if err != nil {
		panic(fmt.Sprintf("error: unable to extract and convert rate for currency \"%s\"; %s\n", to, err))
	}

	return amt * (rateTo / rateFrom)
}
