package countries

import "testing"

func TestVerify(t *testing.T) {
	t.Parallel()

	// Alpha2
	t.Run("Alpha2", func(t *testing.T) {
		t.Parallel()
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
	})

	// Alpha3
	t.Run("Alpha3", func(t *testing.T) {
		t.Parallel()
		for _, v := range CountryList() {
			if ok := v.Alpha3.Verify(); !ok {
				t.Error("Should have been ok")
			}
		}
		if ok := Alpha3("AfG").Verify(); !ok {
			t.Error("Should have been ok")
		}

		if ok := Alpha2("AAa").Verify(); ok {
			t.Error("Shouldn't have been ok")
		}
	})
}

func TestInformation(t *testing.T) {
	t.Parallel()

	// Alpha2
	t.Run("Alpha2", func(t *testing.T) {
		t.Parallel()
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
	})

	// Alpha3
	t.Run("Alpha3", func(t *testing.T) {
		t.Parallel()
		for k, v := range CountryList() {
			if k2, info := v.Alpha3.Information(); info == nil {
				t.Error("The impossible happenned.")
			} else if k != k2 || v != *info {
				t.Error("The informations should have been equal")
			}
		}

		if _, info := Alpha3("ad").Information(); info != nil {
			t.Error("Info should have been nil")
		}
	})
}

func TestString(t *testing.T) {
	// a2 := Alpha2("test")
}
