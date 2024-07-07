package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Frapp",
		Price: 1.2,
		SKU:   "abds-sdf-seff",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
