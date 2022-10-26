package main

import (
    "encoding/xml"
    "fmt"
)

var Sample = `<?xml version="1.0"?>                                                                                                                                                                      
<items>                                                                                                                                                                                     
  <item id="10">                                                                                                                                                                            
    <owner id="50">John Smith</owner>                                                                                                                                                       
    <name>Box</name>                                                                                                                                                                        
    <location>Kitchen</location>                                                                                                                                                            
  </item>                                                                                                                                                                                   
</items>`

type Item struct {
    XMLName  xml.Name `xml:"item"`
    Id       int      `xml:"id,attr"`
	Owner struct {
		Id    int   `xml:"id,attr"`
		Value string `xml:",chardata"`
	} `xml:"owner"`  
    Name     string   `xml:"name"`
    Location string   `xml:"location"`
}

type Items struct {
    XMLName xml.Name `xml:"items"`
    Items   []Item   `xml:"item"`
}

func main() {
    items := Items{}
    data := []byte(Sample)
    xml.Unmarshal(data, &items)
    fmt.Println("items: ", items)
}