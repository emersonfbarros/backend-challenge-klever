package router

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutesIntegration(t *testing.T) {
	utxoResSuccess := `[
		{
			"txid": "1c16ffaad93464a35af0501b95274fe08e2f68beeadc1599cda14f2fb612f1b6",
			"value": "450460",
			"confirmations": 72
		},
		{
			"txid": "c7770cd55fb0c0e35c216b6691206af0aeafc906223abaf225636711904be561",
			"value": "134087",
			"confirmations": 89
		},
		{
			"txid": "133b71c033e38e5ae6760e9af7abc42b8daea90d276709d7f109921267b53e19",
			"value": "120878",
			"confirmations": 192
		}
	]`

	addressResSuccess := `{
		"balance": "17454817",
		"totalReceived": "193498135",
		"totalSent": "176043318",
		"txs": 647
	}`

	txResSuccess := `{
		"vin": [
			{
				"addresses": ["bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"],
				"value": "484817655"
			}
		],
		"vout": [
			{
				"value": "623579",
				"addresses": ["36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6"]
			},
			{
				"value": "3283266",
				"addresses": ["bc1qe29ydjtwyjdmffxg4qwtd5wfwzdxvnap989glq"]
			},
			{
				"value": "90311",
				"addresses": ["bc1qanhueax8r4cn52r38f2h727mmgg6hm3xjlwd0x"]
			}
		],
		"blockHeight": 675674
	}`

	address := "19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n"
	tx := "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"
	invalidAddress := "1nv4lid_4ddre55"
	// invalidTx := "1nv4lid_7x"

	serverErrorMsg := `{"error":"Whatever not found"}`
	// badGatewayMds := `{"message":"Failed to request external resource"}`
	notFoundAddress := fmt.Sprintf(`{"message":"Address %s not found"}`, invalidAddress)
	// notFoundTx := fmt.Sprintf("Transaction %s not found", invalidTx)

	tests := []struct {
		name            string
		tx              string
		address         string
		route           string
		expectedBody    string
		expectedCode    int
		utxoRouteRes    string
		txRouteRes      string
		addressRouteRes string
	}{
		{
			name:         "Test balance route on success",
			address:      address,
			route:        "balance/",
			utxoRouteRes: utxoResSuccess,
			expectedBody: `{"confirmed":"705425","unconfirmed":"0"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Test balance route on not found",
			address:      invalidAddress,
			route:        "balance/",
			utxoRouteRes: serverErrorMsg,
			expectedBody: notFoundAddress,
			expectedCode: http.StatusNotFound,
		},
		{
			name:            "Test details route on success",
			address:         address,
			route:           "details/",
			utxoRouteRes:    utxoResSuccess,
			addressRouteRes: addressResSuccess,
			expectedBody:    `{"address":"19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n","balance":"17454817","totalTx":647,"balanceCalc":{"confirmed":"705425","unconfirmed":"0"},"total":{"sent":"176043318","received":"193498135"}}`,
			expectedCode:    http.StatusOK,
		},
		{
			name:            "Test details route on error",
			address:         invalidAddress,
			route:           "details/",
			utxoRouteRes:    serverErrorMsg,
			addressRouteRes: serverErrorMsg,
			expectedBody:    notFoundAddress,
			expectedCode:    http.StatusNotFound,
		},
		{
			name:         "Test tx route on success",
			tx:           tx,
			route:        "tx/",
			txRouteRes:   txResSuccess,
			expectedBody: `{"addresses":[{"address":"bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r","value":"484817655"},{"address":"36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6","value":"623579"},{"address":"bc1qe29ydjtwyjdmffxg4qwtd5wfwzdxvnap989glq","value":"3283266"},{"address":"bc1qanhueax8r4cn52r38f2h727mmgg6hm3xjlwd0x","value":"90311"}],"block":675674,"txID":"3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"}`,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// http server to mock exeternal api
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/utxo/" + tt.address:
					w.Write([]byte(tt.utxoRouteRes))
				case "/address/" + tt.address:
					w.Write([]byte(tt.addressRouteRes))
				case "/tx/" + tt.tx:
					w.Write([]byte(tt.txRouteRes))
				default:
					w.Write([]byte(serverErrorMsg))
				}
			}))
			// close test server when test finishes
			defer testServer.Close()

			os.Setenv("BASE_URL", testServer.URL)

			// initialize application server
			srv := Initialize(":8080")
			// shutdown server when test finishes
			defer srv.Shutdown(context.Background())

			var routeValue string
			if tt.route == "tx/" {
				routeValue = tt.tx
			} else {
				routeValue = tt.address
			}
			// creates request for balance route
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/"+tt.route+routeValue, nil)
			resp := httptest.NewRecorder()

			// handles request
			srv.Handler.(*gin.Engine).ServeHTTP(resp, req)

			// assertions
			assert.Equal(t, tt.expectedCode, resp.Code)
			assert.Equal(t, tt.expectedBody, resp.Body.String())
		})
	}
}
