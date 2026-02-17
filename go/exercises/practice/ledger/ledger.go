package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string
	Description string
	Change      int
}

func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	var symbol string
	switch currency {
	case "USD":
		symbol = "$"
	case "EUR":
		symbol = "€"
	default:
		return "", errors.New("invalid currency")
	}

	var header string
	switch locale {
	case "en-US":
		header = "Date       | Description               | Change\n"
	case "nl-NL":
		header = "Datum      | Omschrijving              | Verandering\n"
	default:
		return "", errors.New("invalid locale")
	}

	for _, e := range entries {
		if _, err := time.Parse("2006-01-02", e.Date); err != nil {
			return "", err
		}
	}

	sorted := make([]Entry, len(entries))
	copy(sorted, entries)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Date != sorted[j].Date {
			return sorted[i].Date < sorted[j].Date
		}
		if sorted[i].Description != sorted[j].Description {
			return sorted[i].Description < sorted[j].Description
		}
		return sorted[i].Change < sorted[j].Change
	})

	var sb strings.Builder
	sb.WriteString(header)

	for _, e := range sorted {
		t, _ := time.Parse("2006-01-02", e.Date)

		var date string
		switch locale {
		case "en-US":
			date = fmt.Sprintf("%02d/%02d/%04d", t.Month(), t.Day(), t.Year())
		case "nl-NL":
			date = fmt.Sprintf("%02d-%02d-%04d", t.Day(), t.Month(), t.Year())
		}

		desc := e.Description
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		} else {
			desc = fmt.Sprintf("%-25s", desc)
		}

		amount := formatAmount(locale, symbol, e.Change)

		sb.WriteString(date + " | " + desc + " | " + amount + "\n")
	}

	return sb.String(), nil
}

func formatAmount(locale, symbol string, cents int) string {
	negative := cents < 0
	if negative {
		cents = -cents
	}

	whole := cents / 100
	remainder := cents % 100

	wholeStr := strconv.Itoa(whole)

	var thousandSep, decimalSep string
	switch locale {
	case "en-US":
		thousandSep = ","
		decimalSep = "."
	case "nl-NL":
		thousandSep = "."
		decimalSep = ","
	}

	if len(wholeStr) > 3 {
		var parts []string
		for len(wholeStr) > 3 {
			parts = append([]string{wholeStr[len(wholeStr)-3:]}, parts...)
			wholeStr = wholeStr[:len(wholeStr)-3]
		}
		parts = append([]string{wholeStr}, parts...)
		wholeStr = strings.Join(parts, thousandSep)
	}

	amountStr := wholeStr + decimalSep + fmt.Sprintf("%02d", remainder)

	var formatted string
	switch locale {
	case "en-US":
		if negative {
			formatted = "(" + symbol + amountStr + ")"
		} else {
			formatted = symbol + amountStr + " "
		}
	case "nl-NL":
		if negative {
			formatted = symbol + " " + amountStr + "-"
		} else {
			formatted = symbol + " " + amountStr + " "
		}
	}

	return fmt.Sprintf("%13s", formatted)
}
