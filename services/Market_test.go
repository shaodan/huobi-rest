package services

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/shaodan/go-huobi/models"
)

type TestEnv struct {
	Client *HuobiRestClient
}

var env TestEnv

func TestPlaceOrder(t *testing.T) {
	BeforeAll()

	account := env.Client.GetAccounts()

	fmt.Println("Account: ", account)

	if strings.Compare(account.Status, "ok") == 0 {
		accounts := account.Data

		if len(accounts) >= 1 {
			var account models.AccountsData

			for _, entry := range accounts {
				if entry.Type == "spot" {
					account = entry
					break
				}
			}

			if 0 != account.ID {
				var placeParams models.PlaceRequestParams
				placeParams.AccountID = strconv.Itoa(int(account.ID))
				placeParams.Amount = "1.0"
				placeParams.Price = "5721"
				placeParams.Source = "api"
				placeParams.Symbol = "btcusdt"
				placeParams.Type = "sell-limit"

				fmt.Println("Place order with: ", placeParams)
				placeReturn := env.Client.Place(placeParams)
				if placeReturn.Status == "ok" {
					fmt.Println("Place return: ", placeReturn.Data)
				} else {
					t.Errorf("Place error: %s", placeReturn.ErrMsg)
				}
			} else {
				t.Error("account is nil")
			}

		}

	} else {
		t.Error(account.ErrMsg)
	}
}

func Test_getSymbols(t *testing.T) {
	symbols := env.Client.GetSymbols()
	if symbols.Status != "ok" {
		t.Error("failed to get symbols")
	} else {
		t.Logf("toal symbols: %v", len(symbols.Data))
	}
}

func BeforeAll() {
	env.Client = &HuobiRestClient{
		// todo test config
	}
}
