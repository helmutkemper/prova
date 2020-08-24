package commonData

import "strconv"

func (el Price) String() (price string) {
	return "$" + strconv.Itoa(int(el))
}
