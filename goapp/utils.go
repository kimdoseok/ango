package main

import (
	"fmt"
)

func getConditionStr(fstrs []string) (string, []string) {
	searchfields := []string{"code", "codetype", "description"}
	likestr := ""
	likeval := []string{}
	for i, fs := range fstrs {
		if i > 0 {
			likestr += " AND "
		}
		likestr += " ( "
		for j, sf := range searchfields {
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

