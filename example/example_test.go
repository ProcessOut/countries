package countries

import (
	"fmt"
	"testing"

	"github.com/processout/countries"
)

func TestExample(*testing.T) {
	// It is recommended you base you use ISO 3166-1 Alpha2 codes to reference countries,
	// because the map is based on Alpha2 codes. Though you can use what you want.
	myCountry := "US"

	// With Alpha2 type (implements Stringer interface)
	us := countries.Alpha2(myCountry) // Simple cast to invoke methods
	ok := us.Verify()                 // Verifies that the alpha2 code is valid
	ci := us.Information()            // CountryInformation about `us' (i.e. Alpha3, Numeric, FullName)
	fmt.Println(us)
	fmt.Println(ok)

	// With Alpha3 (implements Stringer interface)
	usa := ci.Alpha3            // Retrieves Alpha3 code for `US', AKA `USA'
	ok2 := usa.Verify()         // Verifies that the alpha3 code is valid
	a2, ci := usa.Information() // Same CountryInformation as before
	fmt.Println(a2)
	fmt.Println(ok2)

	// With Numeric
	num := ci.Numeric
	a2, ci2 := countries.FromNumericCode(num) // Alpha2 and CountryInformation from USA numeric
	fmt.Println(a2)

	// With FullName (not recommended because of consistency)
	name := "United States of America"
	a2, ci2 = countries.FromFullName(name) // Alpha2 and CountryInformation from USA full name
	fmt.Println(a2)
	fmt.Println(ci2)

	// Retrieving the list of countries (map[Alpha2]CountryInformation)
	list := countries.List()
	fmt.Println(list[us]) // Has to be uppercase, so watch out
}
