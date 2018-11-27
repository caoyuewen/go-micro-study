package main

import (
	"encoding/json"
	"fmt"
	pb "shippy/consignment-service/proto/consignment"
	"testing"
)

func TestParseFile(t *testing.T) {

	jsonStr := `
{
  "description": "This is a test consignment",
  "weight": "550",
  "containers": [
    {
      "customer_id": "cust001",
      "user_id": "user001",
      "origin": "Manchester, United Kingdom"
    }
  ],
  "vessel_id": "vessel001"
}
`
	//var consignment *pb.Consignment
	consignment:=new(pb.Consignment)
	err := json.Unmarshal([]byte(jsonStr), consignment)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} else {
		fmt.Printf("ok")
	}

}
