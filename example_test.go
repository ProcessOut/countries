package countries

import "testing"

func TestExample(*testing.T) {
	// It is recommended you base you use ISO 3166-1 Alpha2 codes to reference countries,
	// because the map is based on Alpha2 codes. Though you can use what you want.
	myCountry := "US"

	// With Alpha2 type (implements Stringer interface)
	us := Alpha2(myCountry) // Simple cast to invoke methods
	ok := us.Verify()       // Verifies that the alpha2 code is valid
	ci := us.Information()  // CountryInformation about `us' (i.e. Alpha3, Numeric, FullName)

	// With Alpha3 (implements Stringer interface)
	usa := ci.Alpha3         // Retrieves Alpha3 code for `US', AKA `USA'
	ok2 := usa.Verify()      // Verifies that the alpha3 code is valid
	ci2 := usa.Information() // Same CountryInformation as for `us' (ci == ci2)

	// With Numeric
	num := ci2.Numeric
	a2, ci := FromNumeric(num) // Alpha2 and CountryInformation from USA numeric

	// With FullName (not recommended because consistency)
	name := "United States of America"
	a2, ci := FromFullName(name) // Alpha2 and CountryInformation from USA full name

	// Retrieving the list of countries (map[Alpha2]CountryInformation)
	list := List()
}
