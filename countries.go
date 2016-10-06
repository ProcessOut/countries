// Package countries is a package which regroups all countries under the ISO_3166,
// be it with Alpha2, Alpha3, FullName, or Numeric code. The map here is based
// on the Alpha2 for indexing. There are functions to search without giving an
// Alpha 2 however. See README for usage examples.
package countries

import "strings"

// Alpha2 is the ISO_3166-1 alpha 2 code
type Alpha2 string

// Alpha3 is the ISO_3166-1 alpha 3 code
type Alpha3 string

// CountryInformation contains all the information associated with a country,
// according to ISO_3166-1
type CountryInformation struct {
	FullName string
	Alpha3   Alpha3
	Numeric  int
}

func (a Alpha2) format() Alpha2 {
	return Alpha2(strings.ToUpper(a.String()))
}

func (a Alpha3) format() Alpha3 {
	return Alpha3(strings.ToUpper(a.String()))
}

// Verify returns true if the Alpha2 is valid
func (a Alpha2) Verify() bool {
	_, ok := countryList[a.format()]
	return ok
}

// Information returns the CountryInformation associated with the Alpha2 code
func (a Alpha2) Information() *CountryInformation {
	new, ok := countryList[a.format()]
	if !ok {
		return nil
	}
	return &new
}

// String for Stringer interface
func (a Alpha2) String() string {
	return string(a.format())
}

// String for Stringer interface
func (a Alpha3) String() string {
	return string(a.format())
}

// The following functions aren't quite as fast as the other ones,
// and not as safe to use (e.g. wiki standard for full names).
// These functions return "" as alpha2 if nothing is found.

// FromFullName returns the Alpha2 code along with the CountryInformation
// based on the name provided
func FromFullName(name string) (Alpha2, CountryInformation) {
	for k, v := range countryList {
		if v.FullName == name {
			return k, v
		}
	}
	return "", CountryInformation{}
}

// FromAlpha3 returns the Alpha2 code along with the CountryInformation
// based on the Alpha3 code provided
func FromAlpha3(a Alpha3) (Alpha2, CountryInformation) {
	a = a.format()
	for k, v := range countryList {
		if v.Alpha3 == a {
			return k, v
		}
	}
	return "", CountryInformation{}
}

// FromNumericCode returns the Alpha2 code along with the CountryInformation
// based on the numeric code provided
func FromNumericCode(code int) (Alpha2, CountryInformation) {
	for k, v := range countryList {
		if v.Numeric == code {
			return k, v
		}
	}
	return "", CountryInformation{}
}
