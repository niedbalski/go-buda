package buda

import (
	"testing"
	"io/ioutil"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)


func mockResponseFromFile(url string, filepath string) {
	httpmock.Activate()
	response, _ := ioutil.ReadFile(filepath)
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))
}

func TestAPIClient_GetMarkets(t *testing.T) {
	client, _ := NewAPIClient("test", "test")
	mockResponseFromFile(client.FormatResource(MarketsEndpoint), "fixtures/markets.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetMarkets()
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(MarketEndpoint, 1)), "fixtures/market.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetMarket(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetVolumeByMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(MarketVolumeEndpoint, "BTC-CLP")), "fixtures/market_volume.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetVolumeByMarket("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetTickerByMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(MarketTickerEndpoint, "BTC-CLP")), "fixtures/market_ticker.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetTickerByMarket("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetOrderBookByMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(MarketOrderBookEndpoint, "BTC-CLP")), "fixtures/market_order_book.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetOrderBookByMarket("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetTradesByMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(MarketTradesEndpoint, "BTC-CLP")), "fixtures/market_trades.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetTradesByMarket("BTC-CLP", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetBalances(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(BalancesEndpoint), "fixtures/balances.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetBalances()
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetBalance(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(BalancesEndpoint) + "/BTC", "fixtures/balance.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetBalanceByCurrency("BTC")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetOrdersByMarket(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(OrdersEndpoint, "BTC-CLP")), "fixtures/orders.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetOrdersByMarket("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetOrderById(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(OrderEndpoint, 1)), "fixtures/order.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetOrderById(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetDepositsByCurrency(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(DepositsEndpoint, "BTC-CLP")), "fixtures/deposits.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetDepositsByCurrency("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}

func TestAPIClient_GetWithdrawalsByCurrency(t *testing.T) {
	client, _ := NewAPIClient("", "")
	mockResponseFromFile(client.FormatResource(fmt.Sprintf(WithdrawalsEndpoint, "BTC-CLP")), "fixtures/withdrawals.json")
	defer httpmock.DeactivateAndReset()
	markets, err := client.GetWithdrawalsByCurrency("BTC-CLP")
	assert.NoError(t, err)
	assert.NotEmpty(t, markets)
}