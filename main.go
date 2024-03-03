package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// -w
	// --bsid
	// pathtofile
	output, err := exec.Command("/usr/local/bin/aircrack-ng", "-w", "test_passphrases.txt", "wpa_april.cap", "-b", "48:F8:B3:1F:49:F2").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
