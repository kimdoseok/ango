package main

import (
	"fmt"
)

func getConditionStr(fstrs []string) (string, []string) {
	likestr := ""
	likeval := []string{}
	flds := []string{}
	for i, fs := range fstrs {
		if i > 0 {
			likestr += " AND "
		}
		likestr += " ( "
		for j, sf := range flds {
			if j > 0 {
				likestr += " OR "
			}
			likestr += fmt.Sprintf(" %s like ? ", sf)
			likeval = append(likeval, fmt.Sprintf("%%%s%%", fs))
		}
		likestr += " ) "
	}
	//fmt.Println(likestr, likeval)
	return likestr, likeval
}
