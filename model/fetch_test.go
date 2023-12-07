package model

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestFetchSuccess(t *testing.T) {
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

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}

func TestFetchHttpRequestError(t *testing.T) {
	InitModel()
	originalLogger := logger // keeps orginal logger saved
	mockLogger := new(MockLogger)
	logger = mockLogger // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}))
	// closes mock server early for request to fail
	server.Close()

	os.Setenv("BASE_URL", server.URL)

	// creates fetcher for tests
	fetcher := &Fetcher{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()

	got, err := fetcher.Fetch("utxo", "v4al1d4DDr355")
	if err != nil != true {
		t.Errorf("Fetch() error = %v, wantErr %v", err, true)
		return
	}
	if !reflect.DeepEqual(len(got), 0) {
		t.Errorf("Fetch() = %v, want %v", len(got), 0)
	}
	mockLogger.AssertCalled(t, "Errorf", "failed to make http request %v", mock.Anything)
}
