package main

import "fmt"

func main() {
	gaji("Sammi", "rajin", 1000000)
	gaji("dev", "malas", 1000000)
	gaji("dandi", "malas", 1000000)
	gaji("gusnur", "rajin", 1000000)
}
func gaji(nama string, status string, gaji int) {

	statusType := [5]string{"rajin", "malas"}

	isBonusTrue := status == statusType[0]
	isBonusFalse := status == statusType[1]

	resultSalaryPlusBonus := gaji + (gaji / 2)
	resultSalary := gaji

	if isBonusTrue {
		fmt.Println(nama+" -> salary -> Rp ", resultSalaryPlusBonus)
	} else if isBonusFalse {
		fmt.Println(nama+" -> salary -> Rp ", resultSalary)
	} else {
		fmt.Println("tidak memiliki kriteria")
	}
}
