package flightSuggestion

import (
	"fmt"
	"testDataSource"
)

func ExampleFlightSuggestion_FindCheapestRoute() {
	var cheapestRoute Route
	var err error

	var ds = testDataSource.TestDataSource{}
	ds.AppendData("GRU", "BRC", 10)
	ds.AppendData("BRC", "SCL", 5)
	ds.AppendData("BRC", "CDG", 35)
	ds.AppendData("GRU", "CDG", 75)
	ds.AppendData("GRU", "SCL", 20)
	ds.AppendData("GRU", "ORL", 56)
	ds.AppendData("ORL", "CDG", 5)
	ds.AppendData("SCL", "ORL", 20)

	var rl = FlightSuggestion{}
	cheapestRoute, err = rl.FindCheapestRoute(&ds, "GRU", "CDG")
	if err != nil {
		panic(err)
	}

	for _, v := range cheapestRoute.GetRoute() {
		fmt.Printf("%v>%v: $%v\n", v.Origin, v.Destination, v.Price)
	}

	fmt.Printf("Total: $%v\n", cheapestRoute.GetPrice())

	// Output:
	// GRU>BRC: $10
	// BRC>SCL: $5
	// SCL>ORL: $20
	// ORL>CDG: $5
	// Total: $40

}
