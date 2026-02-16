package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string
	Description string
	Change      int
}

var currencySymbols = map[string]string{
	"USD": "$",
	"EUR": "€",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	symbol, ok := currencySymbols[currency]
	if !ok {
		return "", errors.New("invalid currency")
	}

	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("invalid locale")
	}

	sorted := make([]Entry, len(entries))
	copy(sorted, entries)

	for _, e := range sorted {
		if _, err := time.Parse("2006-01-02", e.Date); err != nil {
			return "", err
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Date != sorted[j].Date {
			return sorted[i].Date < sorted[j].Date
		}
		if sorted[i].Description != sorted[j].Description {
			return sorted[i].Description < sorted[j].Description
		}
		return sorted[i].Change < sorted[j].Change
	})

	var b strings.Builder

	if locale == "en-US" {
		b.WriteString("Date       | Description               | Change\n")
	} else {
		b.WriteString("Datum      | Omschrijving              | Verandering\n")
	}

	for _, e := range sorted {
		date := formatDate(e.Date, locale)
		desc := truncateDescription(e.Description)
		change := formatChange(e.Change, symbol, locale)
		b.WriteString(fmt.Sprintf("%-10s | %-25s | %13s\n", date, desc, change))
	}

	return b.String(), nil
}

func formatDate(date string, locale string) string {
	t, _ := time.Parse("2006-01-02", date)
	if locale == "en-US" {
		return t.Format("01/02/2006")
	}
	return t.Format("02-01-2006")
}

func truncateDescription(desc string) string {
	if len(desc) > 25 {
		return desc[:22] + "..."
	}
	return desc
}

func formatChange(cents int, symbol string, locale string) string {
	negative := cents < 0
	abs := cents
	if negative {
		abs = -abs
	}
	units := abs / 100
	frac := abs % 100

	var thousandsSep, decimalSep string
	if locale == "en-US" {
		thousandsSep = ","
		decimalSep = "."
	} else {
		thousandsSep = "."
		decimalSep = ","
	}

	unitsStr := formatWithThousands(units, thousandsSep)
	amount := fmt.Sprintf("%s%s%02d", unitsStr, decimalSep, frac)

	if locale == "en-US" {
		if negative {
			return "(" + symbol + amount + ")"
		}
		return symbol + amount + " "
	}
	if negative {
		return symbol + " " + amount + "-"
	}
	return symbol + " " + amount + " "
}

func formatWithThousands(n int, sep string) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var result strings.Builder
	remainder := len(s) % 3
	if remainder > 0 {
		result.WriteString(s[:remainder])
	}
	for i := remainder; i < len(s); i += 3 {
		if result.Len() > 0 {
			result.WriteString(sep)
		}
		result.WriteString(s[i : i+3])
	}
	return result.String()
}
