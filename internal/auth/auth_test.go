package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApiKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{
			name: "valid api key",
			header: http.Header{
				"Authorization": []string{"Bearer 1234567890"},
			},
			want:    "1234567890",
			wantErr: false,
		},
		{
			name:    "no api key",
			header:  http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid api key",
			header: http.Header{
				"Authorization": []string{"Bearer1234567890"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetApiKey(test.header)
			if (err != nil) != test.wantErr {
				t.Errorf("GetApiKey() error = %v, wantErr %v", err, test.wantErr)
			} else {
				assert.Equal(t, test.want, got)
			}
		})
	}
}
