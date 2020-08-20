package flightSuggestion

import "sort"

func (el *FlightSuggestion) deleteKeyByList(deleteList []int) {
	if len(el.list) == 0 {
		return
	}

	if len(deleteList) == 0 {
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(deleteList)))
	for _, deleteKey := range deleteList {
		el.deleteKey(deleteKey)
	}
}
