package main

import (
	"sort"
)

/*
 * Complete the 'sortIntersect' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY volcanic
 *  2. INTEGER_ARRAY nonVolcanic
 */

func sortIntersect(volcanic []int32, nonVolcanic []int32) []int32 {
	// sort both arrays
	// two pointers, while vol[0] != nonvol[i] && vol[0] > vol[i]
	var output []int32
	sort.Slice(volcanic, func(i, j int) bool { return volcanic[i] > volcanic[j] })
	sort.Slice(nonVolcanic, func(i, j int) bool { return nonVolcanic[i] > nonVolcanic[j] })
	vol := 0
	nonVol := 0
	lenVol := len(volcanic)
	lenNonVol := len(nonVolcanic)
	bigger := 0
	if lenVol > lenNonVol {
		bigger = lenVol
	} else {
		bigger = lenNonVol
	}
	for vol < bigger && nonVol < bigger {
		if volcanic[vol] == nonVolcanic[nonVol] {
			output = append(output, volcanic[vol])
			vol++
			nonVol++
		}
		if volcanic[vol] > nonVolcanic[nonVol] {
			vol++
		} else {
			nonVol++
		}
	}
	return output
}

//func main() {
//	fmt.Println(sortIntersect([]int32{7000, 7000,12000,13000,6900}, []int32{6910,7010,7000,12000,18000,15000,10450}))
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
