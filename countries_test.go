package countries

import "testing"

func TestVerify(t *testing.T) {
	for k := range CountryList() {
		if ok := k.Verify(); !ok {
			t.Error("Should have been ok")
		}
	}
	if ok := Alpha2("Af").Verify(); !ok {
		t.Error("Should have been ok")
	}

	if ok := Alpha2("AA").Verify(); ok {
		t.Error("Shouldn't have been ok")
	}
}

func TestInformation(t *testing.T) {
	for k, v := range CountryList() {
		if info := k.Information(); info == nil {
			t.Error("The impossible happenned.")
		} else if v != *info {
			t.Error("The informations should have been equal")
		}
	}

	if info := Alpha2("bad").Information(); info != nil {
		t.Error("Info should have been nil")
	}
}

func TestString(t *testing.T) {
	// a2 := Alpha2("test")
}
