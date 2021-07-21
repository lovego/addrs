package addrs

import "fmt"

func ExampleIsLocalhost() {
	fmt.Println(IsLocalhost("127.0.0.1"))
	fmt.Println(IsLocalhost("127.1.0.1"))
	addrs, err := IPsOfLocalhost()
	// fmt.Println(addrs)
	fmt.Println(err)
	last := addrs[len(addrs)-1]
	fmt.Println(IsLocalhost(last))
	fmt.Println(IsLocalhost("1.1.1.1"))
	// Output:
	// true <nil>
	// true <nil>
	// <nil>
	// true <nil>
	// false <nil>
}
