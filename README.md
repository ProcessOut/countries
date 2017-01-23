# traveler

[![GoDoc](https://godoc.org/github.com/ProcessOut/traveler?status.svg)](https://godoc.org/github.com/ProcessOut/traveler)

Traveler is a helper package designed to help deal with countries programatically.
It can be called with Alpha2, Alpha3, and numeric codes. By country name too, but that's
not very reliable. Based on the famous ISO 3166-1 standard
Traveler also aggregates data related to those countries for easy access
and referene, such as on which continent a country is, or what currency
is used in a given country

## Installation

```bash
go get gopkg.in/processout/traveler
```

## Usage

- All ISO 4217 countries are supported, and they can be found in `countries-list.go`
- Error handling isn't included here for brevity's sake
- If you call the `String()` method for `Alpha2` or `Alpha3` with an invalid code, it will reply with "`--`" or "`---`" respectively.

`import "github.com/processout/traveler"`

```go
// It is recommended you base you use ISO 3166-1 Alpha2 codes to reference countries,
// because the map is based on Alpha2 codes. Though you can use what you want.
myCountry := "US"

// With Alpha2 type (implements Stringer interface)
us := traveler.CountryAlpha2(myCountry) // Simple cast to invoke methods
ok := us.Verify()                       // Verifies that the alpha2 code is valid
ci := us.Information()                  // CountryInformation about `us' (i.e. Alpha3, Numeric, FullName)
fmt.Println(us)
fmt.Println(ok)
/**/

// With CountryAlpha3 (implements Stringer interface)
usa := ci.CountryAlpha3     // Retrieves Alpha3 code for `US', AKA `USA'
ok2 := usa.Verify()         // Verifies that the alpha3 code is valid
a2, ci := usa.Information() // Same CountryInformation as before
fmt.Println(a2)
fmt.Println(ok2)
/**/

// With Numeric
num := ci.Numeric
a2, ci2 := traveler.FromNumericCode(num) // Alpha2 and CountryInformation from USA numeric
fmt.Println(a2)
/**/

// With FullName (not recommended because of consistency)
name := "United States of America"
a2, ci2 = traveler.FromFullName(name) // Alpha2 and CountryInformation from USA full name
fmt.Println(a2)
fmt.Println(ci2)
/**/

// Retrieving the list of countries (map[Alpha2]CountryInformation)
list := traveler.Countries()
fmt.Println(list[us]) // Has to be uppercase, so watch out
/**/
```

## Notes

 Do not hesitate with issues or pull requests!

##### MIT license!
