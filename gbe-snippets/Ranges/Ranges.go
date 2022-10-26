package main

import "fmt"

func main()  {

	for i, char := range "abc" {
		fmt.Printf("key %d : value %d \n",i,char)
	}

	arrInt := []int{11,21,31,41}
	sumArr := 0
	for _,num := range arrInt {
		sumArr += num
	}
	fmt.Println("Sum of array elements ", sumArr)

	slcInt := make([]int, 3)
	for i, elm := range slcInt {
		slcInt[i] = 10*i
		fmt.Printf("%d -> %d -> %d\n",i, elm, slcInt[i])
	}

	mapStrStr := map[string]string{"key1":"value1","key2":"value2"}
	mapStrStr["key3"] = "value3"
	for keyMap, valueMap := range mapStrStr {
		fmt.Printf("%s : %s\n",keyMap,valueMap)
	}
}