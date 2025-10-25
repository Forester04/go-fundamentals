package api

import (
	"backend/go/crypto/datatypes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRateCurrency(currency string) (*datatypes.Rate, error) {
	if currency == "" {
		return nil, errors.New("currency is empty")
	}
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	rate := datatypes.Rate{}
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyByte, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(bodyByte, &response)
		rate = datatypes.Rate{Currency: currency, Price: response.Bid}
		rate.Price = response.Bid
		rate.Currency = currency

	} else {
		return nil, fmt.Errorf("status code error is: %v", res.StatusCode)
	}
	return &rate, nil
}
