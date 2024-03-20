package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 1.23,
		SKU:   "abs-abc-def",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
