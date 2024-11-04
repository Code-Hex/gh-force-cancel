package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name         string
		urlstr       string
		wantHostname string
		wantRunID    string
		wantRepo     string
		wantErr      bool
	}{
		{
			name:         "valid URL",
			urlstr:       "https://github.com/owner/repo/actions/runs/123456789",
			wantHostname: "github.com",
			wantRunID:    "123456789",
			wantRepo:     "owner/repo",
			wantErr:      false,
		},
		{
			name:         "invalid URL",
			urlstr:       "https://github.com/owner/actions/runs/123456789",
			wantHostname: "",
			wantRunID:    "",
			wantRepo:     "",
			wantErr:      true,
		},
		{
			name:         "ignore query parameters",
			urlstr:       "https://github.com/owner/repo/actions/runs/123456789?pr=9876",
			wantHostname: "github.com",
			wantRunID:    "123456789",
			wantRepo:     "owner/repo",
			wantErr:      false,
		},
		{
			name:         "valid GitHub Enterprise URL",
			urlstr:       "https://git.example.com/owner/repo/actions/runs/123456789",
			wantHostname: "git.example.com",
			wantRunID:    "123456789",
			wantRepo:     "owner/repo",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHostname, gotRunID, gotRepo, err := parseURL(tt.urlstr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHostname != tt.wantHostname {
				t.Errorf("parseURL() got hostname = %v, want %v", gotHostname, tt.wantHostname)
			}
			if gotRunID != tt.wantRunID {
				t.Errorf("parseURL() got runID = %v, want %v", gotRunID, tt.wantRunID)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("parseURL() got repo = %v, want %v", gotRepo, tt.wantRepo)
			}
		})
	}
}
