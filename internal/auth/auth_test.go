package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		want           string
		wantErr        bool
		expectedErrMsg string
	}{
		{
			name:           "No authorization header included",
			headers:        http.Header{},
			want:           "",
			wantErr:        false,
			expectedErrMsg: "no authorization header included",
		},
		{
			name:           "Malformed Authorization header",
			headers:        http.Header{"Authorization": []string{"Malformed"}},
			want:           "",
			wantErr:        true,
			expectedErrMsg: "malformed authorization header",
		},
		{
			name:    "Correctly formed Authorization header",
			headers: http.Header{"Authorization": []string{"ApiKey 12345"}},
			want:    "12345",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.expectedErrMsg {
				t.Errorf("GetAPIKey() error message = %v, expectedErrMsg %v", err.Error(), tt.expectedErrMsg)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
