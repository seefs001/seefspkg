package xutil

import (
	"github.com/seefs001/seefslib-go/xconvertor"
	"strings"
)

func IdsStrToIdsIntGroup(keys string) []int {
	IDS := make([]int, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := xconvertor.StringToInt(ids[i])
		IDS = append(IDS, ID)
	}
	return IDS
}
