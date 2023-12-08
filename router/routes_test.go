package router

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutesIntegration(t *testing.T) {
	utxoRes := `[
		{
			"txid": "1c16ffaad93464a35af0501b95274fe08e2f68beeadc1599cda14f2fb612f1b6",
			"vout": 119,
			"value": "450460",
			"height": 820184,
			"confirmations": 72
		},
		{
			"txid": "c7770cd55fb0c0e35c216b6691206af0aeafc906223abaf225636711904be561",
			"vout": 155,
			"value": "134087",
			"height": 820167,
			"confirmations": 89
		},
		{
			"txid": "133b71c033e38e5ae6760e9af7abc42b8daea90d276709d7f109921267b53e19",
			"vout": 25,
			"value": "120878",
			"height": 820064,
			"confirmations": 192
		}
	]`

	addressRes := `{
		"balance": "17454817",
		"totalReceived": "193498135",
		"totalSent": "176043318",
		"txs": 647
	}`

	t.Run("Test balance route o  n success", func(t *testing.T) {
		// http server to mock exeternal api
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(utxoRes)) // returns expected body
		}))

		defer testServer.Close()

		os.Setenv("BASE_URL", testServer.URL)

		// initiali application server
		srv := Initialize(":8080")
		// shutdown server when test finishes
		defer srv.Shutdown(context.Background())

		// creates requesto for balance route
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/balance/v9waeg80qeg", nil)
		resp := httptest.NewRecorder()

		// handles request
		srv.Handler.(*gin.Engine).ServeHTTP(resp, req)

		// assertions
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, `{"confirmed":"705425","unconfirmed":"0"}`, resp.Body.String())
	})

	t.Run("Test balance route on error", func(t *testing.T) {
		serverRes := `{ "message": "Not found" }`

		// http server to mock exeternal api
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(serverRes)) // returns expected body
		}))

		defer testServer.Close()

		os.Setenv("BASE_URL", testServer.URL)

		// initiali application server
		srv := Initialize(":8080")
		// shutdown server when test finishes
		defer srv.Shutdown(context.Background())

		// creates requesto for balance route
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/balance/not_found_address", nil)
		resp := httptest.NewRecorder()

		// handles request
		srv.Handler.(*gin.Engine).ServeHTTP(resp, req)

		// assertions
		assert.Equal(t, http.StatusBadGateway, resp.Code)
		assert.Equal(t, `{"message":"failed to request external resource"}`, resp.Body.String())
	})

	t.Run("Test details route on success", func(t *testing.T) {
		// http server to mock exeternal api
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/utxo/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n":
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(utxoRes))
			case "/address/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n":
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(addressRes))
			}
		}))
		defer testServer.Close()

		os.Setenv("BASE_URL", testServer.URL)

		// initiali application server
		srv := Initialize(":8080")
		// shutdown server when test finishes
		defer srv.Shutdown(context.Background())

		// creates requesto for balance route
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/details/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n", nil)
		resp := httptest.NewRecorder()

		// handles request
		srv.Handler.(*gin.Engine).ServeHTTP(resp, req)

		// assertions
		expectedBody := `{"address":"19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n","balance":"17454817","totalTx":647,"balanceCalc":{"confirmed":"705425","unconfirmed":"0"},"total":{"sent":"176043318","received":"193498135"}}`
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, expectedBody, resp.Body.String())
	})
}
