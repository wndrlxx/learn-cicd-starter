package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr error
	}{
		{"missing header", "", "", ErrNoAuthHeaderIncluded},
		{"malformed - no space", "ApiKeyabc123", "", errors.New("malformed authorization header")},
		{"malformed - wrong prefix", "Bearer abc123", "", errors.New("malformed authorization header")},
		{"valid", "ApiKey abc123", "abc123", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(headers)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if (err == nil) != (tt.wantErr == nil) {
				t.Errorf("got err %v, want %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("got err %q, want %q", err.Error(), tt.wantErr.Error())
			}
		})
	}
}
