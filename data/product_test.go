package data

import "testing"

func TestAbc(t *testing.T) {
	prod := &Product{
		Name:  "nadir",
		Price: 1,
		SKU:   "1bc-abc-abc",
	}

	if prod.Validation() != nil {
		t.Error(prod.Validation())
	} else {
		t.Log("pass the test")
	}
}
