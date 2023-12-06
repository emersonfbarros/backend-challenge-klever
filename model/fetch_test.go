package model

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type MockLogger struct{}

func (m *MockLogger) Infof(format string, args ...interface{}) {}

func (m *MockLogger) Errorf(format string, args ...interface{}) {}

func TestFetch(t *testing.T) {
	InitModel()
	originalLogger := logger // keeps orginal logger saved
	mockLog := &MockLogger{}
	logger = mockLog // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()

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
		{
			name: "Invalid base URL",
			args: args{
				route: "utxo",
				value: "v4al1d4DDr355",
			},
			wantErr: true,
			setup:   func() { os.Setenv("BASE_URL", "http://invalid-url") },
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

			// Use Server URL
			if tt.setup != nil {
				tt.setup()
			} else {
				os.Setenv("BASE_URL", server.URL)
			}
			got, err := Fetch(tt.args.route, tt.args.value)
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

func TestFetch_HttpError(t *testing.T) {
	originalLogger := logger // keeps orginal logger saved
	mockLog := &MockLogger{}
	logger = mockLog // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()
	tests := []struct {
		name    string
		route   string
		value   string
		wantErr bool
	}{
		{"HttpError", "address", "someValue", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))
			// closes mock server early for request to fail
			server.Close()

			os.Setenv("BASE_URL", server.URL)
			_, err := Fetch(tt.route, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
