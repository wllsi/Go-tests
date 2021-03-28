// IMPORT, MANIPULATE AND SEARCH DATA

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func shellSort(x []string) []string {
	n, k, gaps := len(x), 1, []int{1}

	for {
		a, b, c := k, 2, 1
		for a > 0 {
			if a&1 != 0 {
				c *= b
			}
			a >>= 1
			b *= b
		}
		gap := c + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			for j := i; j > 0; j -= gap {
				if x[j-gap] > x[j] {
					x[j-gap], x[j] = x[j], x[j-gap]
				}
			}
		}
	}
	return x
}

func binSearch(x string, y []string) int {
	high := len(y) - 1
	low := 0

	for high >= low {
		point := ((high - low) / 2) + low

		switch {
		case strings.EqualFold(x, y[point]):
			return point
		case strings.Compare(x, y[point]) == -1:
			high = point - 1
		case strings.Compare(x, y[point]) == 1:
			low = point + 1
		}
	}
	return -1
}

func main() {
	start := time.Now()
	data, err := ioutil.ReadFile("[FILE]")
	if err != nil {
		fmt.Println("File input ERROR:", err)
		return
	}
	// CONVERT DATA TO []string, TRIM "\n" CHARACTERS
	var data_str []string
	setup := ""
	for j := 0; j < len(data); j++ {
		if data[j] == 10 {
			data_str = append(data_str, setup)
			setup = ""
		} else {
			setup = setup + string(data[j])
		}
	}
	//fmt.Println("INPUT DATA:", data_str)
	searched := shellSort(data_str)
	fmt.Println("SORTED SLICE LIBRARY:", searched)
	sought := "maRSupial"
	m := binSearch(sought, searched)
	if m >= 0 {
		fmt.Println("SEARCH:", sought, "\nINDEX RESULT:", m, "\nSLICE LIBRARY MATCH:", searched[m])
	} else {
		fmt.Println(sought, "was not found in the slice's library!")
	}
	elapsed := time.Since(start)
	fmt.Println("ELAPSED TIME:", elapsed)
}