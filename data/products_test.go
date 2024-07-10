package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Frapp",
		Price: 2,
		SKU:   "abd-sdf-seff",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
