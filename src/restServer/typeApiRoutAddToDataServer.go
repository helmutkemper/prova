package restServer

import "commonData"

type apiRoutAddToDataServer struct {
	Api         apiRestful
	Origin      string
	Destination string
	Price       commonData.Price
}
