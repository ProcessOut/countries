package countries

import "fmt"

func ExampleAlpha2_Verify() {
	us := Alpha2("us") // Simple cast to invoke methods
	fmt.Println(us)    // Implements stringer interface (prints uppercase always)
	// Output: US
}

func ExampleAlpha2_Information() {
	info := Alpha2("us").Information()
	fmt.Println(info)
	// Output: &{United States of America USA 840}
}

func ExampleFromNumericCode() {
	a2, info := FromNumericCode(10)
	fmt.Println(a2)   // Alpha2 code
	fmt.Println(info) // CountryInformation
	// Output:
	// AQ
	// {Antarctica ATA 10}
}

/*
func TestExample(*testing.T) {
	ok := us.Verify()      // Verifies that the alpha2 code is valid
	ci := us.Information() // CountryInformation about `us' (i.e. Alpha3, Numeric, FullName)
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
	a2, ci2 := FromNumericCode(num) // Alpha2 and CountryInformation from USA numeric
	fmt.Println(a2)

	// With FullName (not recommended because of consistency)
	name := "United States of America"
	a2, ci2 = FromFullName(name) // Alpha2 and CountryInformation from USA full name
	fmt.Println(a2)
	fmt.Println(ci2)

	// Retrieving the list of countries (map[Alpha2]CountryInformation)
	list := List()
	fmt.Println(list[us]) // Has to be uppercase, so watch out
}
*/
