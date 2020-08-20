package flightSuggestion

import (
	"errors"
	"sort"
)

func (el *Route) deleteByKeyList(listKeysToDelete []int) (err error) {
	if len(el.route) == 0 {
		err = errors.New("route is empty")
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(listKeysToDelete)))
	for _, keyToDelete := range listKeysToDelete {
		if keyToDelete >= len(el.route) {
			err = errors.New("key not found")
			return
		}

		el.route = append(el.route[:keyToDelete], el.route[keyToDelete+1:]...)
	}

	return
}
