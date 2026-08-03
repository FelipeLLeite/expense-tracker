package models

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

const dateOnlyLayout = "2006-01-02"

type ExpenseDate struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *ExpenseDate) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		logger.Error("expenseDate: expected a JSON string", "error", err)
		return fmt.Errorf("expenseDate: expected a JSON string: %w", err)
		// return logger.
	}

	raw = strings.TrimSpace(raw)
	if raw == "" {
		d.Time = time.Time{}
		return nil
	}

	if t, err := time.Parse(time.RFC3339, raw); err == nil {
		d.Time = t
		return nil
	}

	if t, err := time.Parse(dateOnlyLayout, raw); err == nil {
		// Unspecified time fields from this layout default to
		// 00:00:00, and the location defaults to UTC.
		d.Time = t
		return nil
	}

	return fmt.Errorf("expenseDate: unrecognized date format %q (want RFC3339 or YYYY-MM-DD)", raw)
}

// MarshalJSON implements json.Marshaler.
func (d ExpenseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(time.RFC3339))
}

type Expense struct {
	ID          int64       `json:"id"`
	ExpenseDate ExpenseDate `json:"expenseDate"`
	Amount      float64     `json:"amount"`
	Currency    string      `json:"currency"`

	// Category/Subcategory replace the single "type" field from the
	// original schema, e.g. Category="Food", Subcategory="Groceries".
	Category    string `json:"category"`
	Subcategory string `json:"subcategory,omitempty"`

	Description   string `json:"description,omitempty"`
	PaymentMethod string `json:"paymentMethod,omitempty"`

	IsRecurring       bool   `json:"isRecurring,omitempty"`
	RecurringInterval string `json:"recurringInterval,omitempty"`

	// CreatedAt/UpdatedAt track record bookkeeping (when the entry was
	// added/edited), distinct from ExpenseDate (when the expense
	// actually occurred). These are always full timestamps.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
