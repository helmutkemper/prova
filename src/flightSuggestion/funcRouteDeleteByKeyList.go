package flightSuggestion

import (
	"errors"
	"sort"
)

func (el *Route) deleteByKeyList(listKeysToDelete []int) (err error) {
	if len(el.route) == 0 {
		err = errors.New("Route is empty")
		return
	}

	for _, keyToDelete := range listKeysToDelete {
		if keyToDelete < 0 {
			err = errors.New("key must be a positive number")
			return
		}

		if keyToDelete >= len(el.route) {
			err = errors.New("key not found")
			return
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(listKeysToDelete)))
	for _, keyToDelete := range listKeysToDelete {
		err = el.deleteByKey(keyToDelete)
		if err != nil {
			return err
		}
	}

	el.recalculateTotalPrice()

	return
}
