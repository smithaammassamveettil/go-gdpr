package vendorlist

import "testing"

func TestEagerlyParsedVendorList(t *testing.T) {
	AssertVendorlistCorrectness(t, func(data []byte) VendorList {
		vendorList, err := ParseEagerly(data)
		if err != nil {
			t.Errorf("ParseEagerly returned an unexpected error: %v", err)
		}
		return vendorList
	})
}
