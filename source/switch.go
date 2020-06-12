package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 1; i <= 6; i++ {
		switch i {
		case 1:
			fmt.Println("satu")
		case 2:
			fmt.Println("dua")
		case 3:
			fmt.Println("tiga")
		case 4:
			fmt.Println("empat")
		case 5:
			fmt.Println("lima")
		default:
			fmt.Println("default")
		}
	}

	// mengambil sistem operasi
	sistemOperasi := runtime.GOOS
	switch sistemOperasi {
	case "darwin":
		fmt.Println("Mac")
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("unknown")
	}
	if sistemOperasi == "linux" {
		fmt.Println("Mac")
	} else {
		fmt.Println("gak tau")
	}

}
