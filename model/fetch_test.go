package model

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	type args struct {
		route string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
		setup   func()
	}{
		{
			name: "Valid address in address route",
			args: args{route: "address",
				value: "v4al1d4DDr355",
			},
			want:    []byte(`OK`),
			wantErr: false,
		},
		{
			name: "Invalid address in address route",
			args: args{
				route: "address",
				value: "Inv4al1d4DDr355",
			},
			want:    []byte(`Not OK`),
			wantErr: false,
		},
		{
			name: "Valid address in utxo route",
			args: args{
				route: "utxo",
				value: "v4al1d4DDr355",
			},
			want:    []byte(`OK`),
			wantErr: false,
		},
		{
			name: "Invalid address in utxo route",
			args: args{
				route: "utxo",
				value: "Inv4al1d4DDr355",
			},
			want:    []byte(`Not OK`),
			wantErr: false,
		},
		{
			name: "Valid tx in tx route",
			args: args{
				route: "tx",
				value: "qefv987912g8y19hy891h61h098",
			},
			want:    []byte(`OK`),
			wantErr: false,
		},
		{
			name: "Invalid address in tx route",
			args: args{
				route: "tx",
				value: "qefv987912g8y19hy891h61h098",
			},
			want:    []byte(`Not OK`),
			wantErr: false,
		},
		{
			name: "Invalid route",
			args: args{
				route: "invalid",
				value: "v4al1d4DDr355",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock server
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Write([]byte(tt.want))
			}))
			// Close the server when test finishes
			defer server.Close()

			os.Setenv("BASE_URL", server.URL)

			// creates fetcher for tests
			fetcher := &Fetcher{
				BaseURL:  os.Getenv("BASE_URL"),
				Username: os.Getenv("USERNAME"),
				Password: os.Getenv("PASSWORD"),
			}

			got, err := fetcher.Fetch(tt.args.route, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
