package shortener_test

import (
	"testing"
	"git.sr.ht/~will-clarke/url-shortener-golang/shortener"
)

func TestURL_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       shortener.URL
		wantErr bool
	}{
		{
			name:    "valid URL",
			u:       shortener.URL("https://www.example.com/lolz"),
			wantErr: false,
		},
		{
			name:    "invalid URL",
			u:       shortener.URL("this isn't a real URL guys"),
			wantErr: true,
		},
		{
			name:    "no URL",
			u:       shortener.URL(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("URL.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
