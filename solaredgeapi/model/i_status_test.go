package model

import "testing"

func TestIStatus_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		status  IStatus
		wantErr bool
	}{
		{"OFF", "OFF", 1, false},
		{"MPPT", "MPPT", 4, false},
		{"THROTTLED", "THROTTLED", 5, false},
		{"lowercase", "off", 0, true},
		{"unknown", "MYSTATUS", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.status.UnmarshalJSON([]byte(`"` + tt.data + `"`)); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
