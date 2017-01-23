// Package traveler is a package which regroups all countries under the ISO_3166,
// be it with Alpha2, Alpha3, FullName, or Numeric code. The map here is based
// on the Alpha2 for indexing. There are functions to search without giving an
// Alpha 2 however. See README for usage examples.
// Traveler also aggregates data related to those countries for easy access
// and referene, such as on which continent a country is, or what currency
// is used in a given country
// Note that what you can do with an Alpha2, you can also do with an Alpha3
package traveler

import "strings"

// CountryAlpha2 is the country ISO_3166-1 alpha 2 code
type CountryAlpha2 string

// CountryAlpha3 is the country ISO_3166-1 alpha 3 code
type CountryAlpha3 string

// CountryInformation contains all the information associated with a country,
// according to ISO_3166-1
type CountryInformation struct {
	FullName      string
	CountryAlpha3 CountryAlpha3
	Numeric       int
}

// All the following functions come in pair for CountryAlpha2 and CountryAlpha3

// WrongCountryAlpha2 represents what String() will print if
// CountryAlpha2 is wrong
var WrongCountryAlpha2 = CountryAlpha2("--")

// WrongCountryAlpha3 represents what String() will print if
// CountryAlpha3 is wrong
var WrongCountryAlpha3 = CountryAlpha3("---")

func (a CountryAlpha2) format() CountryAlpha2 {
	return CountryAlpha2(strings.ToUpper(string(a)))
}

func (a CountryAlpha3) format() CountryAlpha3 {
	return CountryAlpha3(strings.ToUpper(string(a)))
}

// Verify returns true if the CountryAlpha2 is valid
func (a CountryAlpha2) Verify() bool {
	_, ok := countryList[a.format()]
	return ok
}

// Verify returns true if the CountryAlpha3 is valid
func (a CountryAlpha3) Verify() bool {
	k, _ := FromCountryAlpha3(a)
	return k != ""
}

// Information returns the CountryInformation associated with the
// CountryAlpha2 code
func (a CountryAlpha2) Information() *CountryInformation {
	new, ok := countryList[a.format()]
	if !ok {
		return nil
	}
	return &new
}

// Information returns the CountryInformation associated with the
// CountryAlpha2 code
func (a CountryAlpha3) Information() (CountryAlpha2, *CountryInformation) {
	k, new := FromCountryAlpha3(a)
	if k == "" {
		return "", nil
	}
	return k, &new
}

// Continent returns the continent on which the country is
func (a CountryAlpha2) Continent() string {
	return countryToContinents[a]
}

// Continent returns the continent on which the country is
func (a CountryAlpha3) Continent() string {
	k, _ := FromCountryAlpha3(a)
	return countryToContinents[k]
}

// Currency returns the currency of the country
func (a CountryAlpha2) Currency() string {
	return countryToCurrency[a]
}

// Currency returns the currency of the country
func (a CountryAlpha3) Currency() string {
	k, _ := FromCountryAlpha3(a)
	return countryToCurrency[k]
}

// String for Stringer interface
func (a CountryAlpha2) String() string {
	if ok := a.Verify(); !ok {
		return string(WrongCountryAlpha2)
	}
	return string(a.format())
}

// String for Stringer interface
func (a CountryAlpha3) String() string {
	if ok := a.Verify(); !ok {
		return string(WrongCountryAlpha3)
	}
	return string(a.format())
}

// The following functions aren't quite as fast as the other ones,
// and not as safe to use (e.g. wiki standard for full names).
// These functions return "" as CountryAlpha2 if nothing is found.

// FromFullName returns the CountryAlpha2 code along with the CountryInformation
// based on the name provided.
func FromFullName(name string) (CountryAlpha2, CountryInformation) {
	for k, v := range countryList {
		if v.FullName == name {
			return k, v
		}
	}
	return "", CountryInformation{}
}

// FromCountryAlpha3 returns the CountryAlpha2 code along with the
// CountryInformation based on the CountryAlpha3 code provided
func FromCountryAlpha3(a CountryAlpha3) (CountryAlpha2, CountryInformation) {
	a = a.format()
	for k, v := range countryList {
		if v.CountryAlpha3 == a {
			return k, v
		}
	}
	return "", CountryInformation{}
}

// FromNumericCode returns the CountryAlpha2 code along with the
// CountryInformation based on the numeric code provided
func FromNumericCode(code int) (CountryAlpha2, CountryInformation) {
	for k, v := range countryList {
		if v.Numeric == code {
			return k, v
		}
	}
	return "", CountryInformation{}
}
