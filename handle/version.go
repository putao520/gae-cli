package handle

import "fmt"

var GlobalVer = "0.0.3"

func ArgVersion() {
	fmt.Printf("GAE Version:%s", GlobalVer)
}
