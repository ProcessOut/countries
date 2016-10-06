package countries

import "fmt"

func ExampleAlpha2_Verify() {
	us := Alpha2("us") // Simple cast to invoke methods
	ok := us.Verify()  // Verifies it is a correct ISO 3166 Alpha2 code
	fmt.Println(us)    // Implements stringer interface (prints uppercase always)
	fmt.Println(ok)    // Bool that says if "us" is a correct Alpha2 code
	// Output:
	// US
	// true
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

func ExampleFromFullName() {
	a2, info := FromFullName("Yemen")
	fmt.Println(a2)   // Alpha2 code
	fmt.Println(info) // CountryInformation
	// Output:
	// YE
	// {Yemen YEM 887}
}

func ExampleList() {
	list := List()
	fmt.Println(list["US"])
	// Output: {United States of America USA 840}
}
