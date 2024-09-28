package exec

import (
	"encoding/xml"
	"fmt"
)


type Plant struct {
	XMLName  xml.Name `xml:"plant"`
	Id 		 int 	  `xml:"id,attr"`
	Name 	 string   `xml:"name"`
}

func CreateXml(){
	coffe := &Plant{Id: 27, Name:"Coffe"}

	out, _ := xml.MarshalIndent(coffe, "", "  ")
	fmt.Println(string(out))
}
