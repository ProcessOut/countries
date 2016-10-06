package countries

import "testing"

func TestVerify(t *testing.T) {
	if ok := Alpha2("AF").Verify(); !ok {
		t.Error("Should have been ok")
	}
	if ok := Alpha2("Af").Verify(); !ok {
		t.Error("Should have been ok")
	}
	if ok := Alpha2("AA").Verify(); ok {
		t.Error("Shouldn't have been ok")
	}
}
