package cloud

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Precision interface {
	Precision() uint64
}

type Price struct{}

func (p Price) Precision() uint64 {
	return 2 //nolint:mnd //constant
}

type Quantity struct{}

func (q Quantity) Precision() uint64 {
	return 3 //nolint:mnd //constant
}

type Percent struct{}

func (d Percent) Precision() uint64 {
	return 5 //nolint:mnd //constant
}

// Decimal represents a fixed‑point decimal stored as a scaled int64 using the provided precision.
type Decimal[T Precision] struct {
	p T
	v int64
}

// MarshalJSON converts the Decimal to a JSON string.
func (d Decimal[T]) MarshalJSON() ([]byte, error) {
	precision := d.p.Precision()
	divider := math.Pow10(int(precision)) //nolint:gosec //acceptable

	high := d.v / int64(divider)
	low := d.v % int64(divider)

	return fmt.Appendf(nil, "%d.%0"+strconv.FormatUint(precision, 10)+"d", high, low), nil
}

// UnmarshalJSON parses a JSON string into a Decimal.
func (d *Decimal[T]) UnmarshalJSON(data []byte) error {
	precision := d.p.Precision()
	multiplier := int64(math.Pow10(int(precision))) //nolint:gosec //acceptable

	s := string(data)
	dotIndex := strings.Index(s, ".")
	if dotIndex == -1 {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse decimal: %w", err)
		}

		*d = Decimal[T]{p: d.p, v: v * multiplier}
		return nil
	}

	high, err := strconv.ParseInt(s[:dotIndex], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse decimal: %w", err)
	}

	low, err := strconv.ParseInt(s[dotIndex+1:], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse decimal: %w", err)
	}

	*d = Decimal[T]{p: d.p, v: high*multiplier + low}
	return nil
}
