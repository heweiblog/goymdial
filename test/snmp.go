package main

import (
	"github.com/alouca/gosnmp"
	"log"
	//"strconv"
)

func main() {
	s, err := gosnmp.NewGoSNMP("192.168.127.1", "yamu.com", gosnmp.Version2c, 5)
	if err != nil {
		log.Fatal(err)
	}
	resps, err := s.Walk(".1.3.6.1.2.1.2.2.1.16")
	if err == nil {
		for _, v := range resps {
			switch v.Type {
			case gosnmp.OctetString:
				log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
			case gosnmp.Integer:
				log.Printf("Response: %s : %d : %s \n", v.Name, v.Value.(int), v.Type.String())
			case gosnmp.Counter32:
				log.Printf("Response: %s : %d : %s \n", v.Name, v.Value.(uint64), v.Type.String())
			}
		}
	}

	resp, err := s.Get(".1.3.6.1.2.1.1.5.0")
	if err == nil {
		for _, v := range resp.Variables {
			switch v.Type {
			case gosnmp.OctetString:
				log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
			}
		}
	}
}
