package traveler

import (
	"math/rand"
	"testing"
)

func TestVerify(t *testing.T) {
	t.Parallel()

	// CountryAlpha2
	t.Run("CountryAlpha2", func(t *testing.T) {
		t.Parallel()
		for k := range Countries() {
			if ok := k.Verify(); !ok {
				t.Error("Should have been ok")
			}
		}
		if ok := CountryAlpha2("Af").Verify(); !ok {
			t.Error("Should have been ok")
		}

		if ok := CountryAlpha2("AA").Verify(); ok {
			t.Error("Shouldn't have been ok")
		}
	})

	// CountryAlpha3
	t.Run("CountryAlpha3", func(t *testing.T) {
		t.Parallel()
		for _, v := range Countries() {
			if ok := v.CountryAlpha3.Verify(); !ok {
				t.Error("Should have been ok")
			}
		}
		if ok := CountryAlpha3("AfG").Verify(); !ok {
			t.Error("Should have been ok")
		}

		if ok := CountryAlpha3("AAa").Verify(); ok {
			t.Error("Shouldn't have been ok")
		}
	})
}

func TestInformation(t *testing.T) {
	t.Parallel()

	// CountryAlpha2
	t.Run("CountryAlpha2", func(t *testing.T) {
		t.Parallel()
		for k, v := range Countries() {
			if info := k.Information(); info == nil {
				t.Error("The impossible happenned.")
			} else if v != *info {
				t.Error("The informations should have been equal")
			}
		}

		if info := CountryAlpha2("bad").Information(); info != nil {
			t.Error("Info should have been nil")
		}
	})

	// CountryAlpha3
	t.Run("CountryAlpha3", func(t *testing.T) {
		t.Parallel()
		for k, v := range Countries() {
			if k2, info := v.CountryAlpha3.Information(); info == nil {
				t.Error("The impossible happenned.")
			} else if k != k2 || v != *info {
				t.Error("The informations should have been equal")
			}
		}

		if _, info := CountryAlpha3("ad").Information(); info != nil {
			t.Error("Info should have been nil")
		}
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	// CountryAlpha2
	t.Run("CountryAlpha2", func(t *testing.T) {
		t.Parallel()
		a := CountryAlpha2("Af")
		if a.String() != "AF" {
			t.Error("Wrong answer: " + a.String())
		}
		a = CountryAlpha2("afa")
		if a.String() != string(WrongCountryAlpha2) {
			t.Error("Wrong answer: " + a.String())
		}
	})

	// CountryAlpha3
	t.Run("CountryAlpha3", func(t *testing.T) {
		t.Parallel()
		a := CountryAlpha3("Afg")
		if a.String() != "AFG" {
			t.Error("Wrong answer: " + a.String())
		}
		a = CountryAlpha3("afag")
		if a.String() != string(WrongCountryAlpha3) {
			t.Error("Wrong answer: " + a.String())
		}
	})
}

func TestFrom(t *testing.T) {
	t.Parallel()

	tester := func(f func() (CountryAlpha2, CountryInformation),
		ogK CountryAlpha2, ogV CountryInformation) {

		a2, ci := f()
		if a2 != ogK {
			t.Error("CountryAlpha2 codes don't match")
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
					func() (CountryAlpha2, CountryInformation) {
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

			// FromCountryAlpha3
			t.Run("CountryAlpha3", func(t *testing.T) {
				t.Parallel()
				tester(
					func() (CountryAlpha2, CountryInformation) {
						return FromCountryAlpha3(v.CountryAlpha3)
					},
					k,
					v,
				)

				a2, _ := FromCountryAlpha3("PRO")
				if a2 != "" {
					t.Error("ProcessOut is not a country (yet)")
				}
			})

			// FromNumericCode
			t.Run("NumericCode", func(t *testing.T) {
				t.Parallel()
				tester(
					func() (CountryAlpha2, CountryInformation) {
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
