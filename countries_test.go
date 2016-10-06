package countries

import (
	"math/rand"
	"testing"
)

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
	t.Parallel()

	// Alpha2
	t.Run("Alpha2", func(t *testing.T) {
		t.Parallel()
		a := Alpha2("Af")
		if a.String() != "AF" {
			t.Error("Wrong answer: " + a.String())
		}
		a = Alpha2("afa")
		if a.String() != string(WrongAlpha2) {
			t.Error("Wrong answer: " + a.String())
		}
	})

	// Alpha3
	t.Run("Alpha3", func(t *testing.T) {
		t.Parallel()
		a := Alpha3("Afg")
		if a.String() != "AFG" {
			t.Error("Wrong answer: " + a.String())
		}
		a = Alpha3("afag")
		if a.String() != string(WrongAlpha3) {
			t.Error("Wrong answer: " + a.String())
		}
	})
}

func TestFrom(t *testing.T) {
	t.Parallel()

	tester := func(f func() (Alpha2, CountryInformation), ogK Alpha2, ogV CountryInformation) {
		a2, ci := f()
		if a2 != ogK {
			t.Error("Alpha2 codes don't match")
		} else if ci != ogV {
			t.Error("CountryInformation doesn't match")
		}
	}

	var c int
	rand := rand.Intn(len(countryList))
	for k, v := range countryList {
		if c != rand {
			c++
			continue
		}

		t.Run(k.String(), func(t *testing.T) {
			t.Parallel()

			// FromFullName
			t.Run("FullName", func(t *testing.T) {
				t.Parallel()
				tester(
					func() (Alpha2, CountryInformation) {
						return FromFullName(v.FullName)
					},
					k,
					v,
				)

				a2, _ := FromFullName("ProcessOut")
				if a2 != "" {
					t.Error("ProcessOut is not a country (yet)")
				}
			})

			// FromAlpha3
			t.Run("Alpha3", func(t *testing.T) {
				t.Parallel()
				tester(
					func() (Alpha2, CountryInformation) {
						return FromAlpha3(v.Alpha3)
					},
					k,
					v,
				)

				a2, _ := FromAlpha3("PRO")
				if a2 != "" {
					t.Error("ProcessOut is not a country (yet)")
				}
			})

			// FromNumericCode
			t.Run("NumericCode", func(t *testing.T) {
				t.Parallel()
				tester(
					func() (Alpha2, CountryInformation) {
						return FromNumericCode(v.Numeric)
					},
					k,
					v,
				)

				a2, _ := FromNumericCode(8406355007)
				if a2 != "" {
					t.Error("ProcessOot is not a country (yet)")
				}
			})
		})
		break
	}
}
