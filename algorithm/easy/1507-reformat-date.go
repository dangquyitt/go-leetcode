package main

import (
	"fmt"
)

func reformatDate(date string) string {
	result := ""
	monthMap := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	if len(date) == 12 {
		date = "0" + date
	}

	year := date[9:]
	keyMonth := date[5:8]
	monthVal := monthMap[keyMonth]
	day := date[:2]
	result = year + "-" + monthVal + "-" + day
	return result
}

func main() {
	fmt.Println(reformatDate("6th Jun 1933"))
	fmt.Println(reformatDate("26th May 1960"))
}
