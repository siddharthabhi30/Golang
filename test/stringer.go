package main

import "fmt"

type IPAddr [4]byte

func main() {

	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	var aa IPAddr
	aa[0] = 12
	fmt.Println(aa)

}
