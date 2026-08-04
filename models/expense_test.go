package models_test

import (
	"testing"

	"github.com/FelipeLLeite/expense-tracker/models"
)

func TestExpenseDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"full RFC3339 datetime", []byte(`"2026-07-03T14:30:00Z"`), false},
		{"date only, defaults to midnight", []byte(`"2026-07-03"`), false},
		{"empty string", []byte(`""`), false},
		{"invalid format", []byte(`"07/03/2026"`), true},
		{"invalid data", []byte(`{}`), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var d models.ExpenseDate
			gotErr := d.UnmarshalJSON(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
			}
		})
	}
}
