package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	db, err := geoip2.Open("D:/GeoLite2-City_20220329/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if len(os.Args) != 0 {
		p := os.Args[1]
		ip := net.ParseIP(p)
		record, err := db.City(ip)
		if err != nil {
			log.Fatal(err)
		}
		color.Set(color.FgMagenta, color.Bold)
		defer color.Unset()
		// color.Set(color.FgYellow)
		// red := color.New(color.FgRed)
		// boldRed := red.Add(color.Bold)
		fmt.Printf("ip: %v\n", p)
		fmt.Printf("位于: %v\n", record.Continent.Names["zh-CN"])
		fmt.Printf("国家: %v\n", record.Country.Names["zh-CN"])
		if record.Subdivisions != nil {
			for _, val := range record.Subdivisions {
				fmt.Printf("省份: %v\n", val.Names["zh-CN"])
			}
		}
		if record.City.Names != nil {
			fmt.Printf("城市: %v\n", record.City.Names["zh-CN"])
		}
		fmt.Printf("代码: %v\n", record.Country.IsoCode)
		fmt.Printf("时区: %v\n", record.Location.TimeZone)
		if record.Postal.Code != "" {
			fmt.Printf("邮编: %v\n", record.Postal.Code)
		}
		fmt.Printf("经纬度: %v,%v\n", record.Location.Latitude, record.Location.Longitude)
		//color.Unset()
	}
}
