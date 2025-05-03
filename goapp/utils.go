package main

import (
	"fmt"
	"regexp"
	"strings"
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

func checkOrigin(origin string) bool {
	var result = false
	hostname := ""
	re, _ := regexp.Compile("([-a-zA-Z0-9])+")
	matched := re.FindAllString(origin, 2)
	if len(matched) > 1 {
		hostname = matched[1]
	} else if len(matched) == 1 {
		hostname = matched[1]
	}

	//fmt.Println("Hostname: ", hostname)

	for _, a := range sitesallowed {
		if strings.EqualFold(hostname, a) {
			//fmt.Println("origin==a", hostname, a)
			result = true
			break
		}
	}
	//fmt.Println("checkOrigin result:", result)
	return result
}
