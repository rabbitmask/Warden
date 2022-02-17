package Wash

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

type Address struct {
	City        string
	Subdivision string
	Country     string
	ISO         string
	//Coordinates []float64
}

func GetAddress(aip string) Address {
	db, err := geoip2.Open("Config/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(aip)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("City Name: %v\n", record.City.Names["en"])
	//if len(record.Subdivisions) > 0 {
	//	fmt.Printf("Subdivision Name: %v\n", record.Subdivisions[0].Names["en"])
	//}
	//fmt.Printf("Country Name: %v\n", record.Country.Names["en"])
	//fmt.Printf("ISO Country Code: %v\n", record.Country.IsoCode)
	//fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	var address Address
	if len(record.Subdivisions) > 0 {
		address = Address{City: record.City.Names["en"], Subdivision: record.Subdivisions[0].Names["en"], Country: record.Country.Names["en"], ISO: record.Country.IsoCode}
	} else {
		address = Address{City: record.City.Names["en"], Subdivision: "", Country: record.Country.Names["en"], ISO: record.Country.IsoCode}
	}

	return address
}
