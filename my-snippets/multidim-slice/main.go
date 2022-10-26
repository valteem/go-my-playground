package main

import "fmt"

func main() {

	mds := [][][]string{{
		                 	{"e000","e001","e002"},
	                     	{"e010","e011","e012"},
						 },
						{
							{"e100","e101", "e102"},
							{"e110", "e111","e112"},
						},
						{
							{"e200", "e201", "e202"},
							{"e210", "e211", "e212"},

						},
					}

	fmt.Printf("%v\n", mds)
	fmt.Println("len(mds) = ", len(mds)) // this returns a number of 'rows'

	for i := 0; i < len(mds); i++ {
		for j := 0; j < len(mds[i]); j++ {
			for k := 0; k < len(mds[j]); k++ {
				fmt.Println(i, j, k, mds[i][j][k])
			}
		}
	}
}
