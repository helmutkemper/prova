package flightSuggestion

import (
	"errors"
	"sort"
)

// Apaga rotas baseados em uma lista
//   deleteList: lista de inteiro positivo
func (el *FlightSuggestion) deleteKeyByList(deleteList []int) (err error) {
	if len(el.list) == 0 {
		return
	}

	if len(deleteList) == 0 {
		return
	}

	for _, keyToDelete := range deleteList {
		if keyToDelete < 0 {
			err = errors.New("key must be a positive number")
			return
		}

		if keyToDelete >= len(el.list) {
			err = errors.New("key not found")
			return
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(deleteList)))
	for _, keyToDelete := range deleteList {
		err = el.deleteKey(keyToDelete)
		if err != nil {
			return
		}
	}

	return
}
