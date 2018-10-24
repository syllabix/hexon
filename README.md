# Hexon
A simple client for working with a limited portion of the hexon api.

## Intended usage
This project is set up to be managed via dep [https://golang.github.io/dep/](https://golang.github.io/dep/).

```
dep ensure -add github.com/syllabix/hexon
```

```
//Example:
package main

import (
	"log"
	"time"

	"github.com/syllabix/hexon"
)

const (
	siteCode = "foobar123"
)

func handleResponse(res *hexon.APIResponse, err error) {
	if err == hexon.ErrorHexonAPI {
		for _, msg := range res.Errors {
			log.Println(msg)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	creds := hexon.Credentials{
		Username: "username",
		Password: "nice-password",
	}

	hexonClient, err := hexon.NewClient(creds)
	if err != nil {
		log.Fatal(err)
	}

	car := hexon.Vehicle{
		Vin:                   "BBBB23ERTASDF90764",
		LicenseNumber:         "5XDZ78",
		LocationCode:          "lp.10.8",
		Make:                  "Volkswagen",
		Model:                 "Golf",
		BodyStyle:             "Hatchback",
		HexonCategory:         "car",
		Currency:              "EUR",
		ExpectedDateAvailable: time.Date(2018, time.November, 23, 0, 0, 0, 0, time.UTC),
		PriceIncludingVat:     11043.29,
		IsNewCar:              false,
		ExpectedMileage:       119273,
		MileageUnit:           "kilometres",
		ExteriorColor:         "white",
		FuelType:              "diesel",
		IsVatDeductible:       true,
	}

	//Create Vehicle
	res, err := hexonClient.CreateVehicle(car)
	log.Println("Handling Create")
	handleResponse(res, err)

	res, err = hexonClient.PublishVehicle(car.Vin, siteCode)
	log.Println("Handling Publish")
	handleResponse(res, err)

	res, err = hexonClient.UpdateVehicle(car)
	log.Println("Handling Update")
	handleResponse(res, err)
}

```
