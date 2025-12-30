package main

import (
	"fmt"
	"log"

	"github.com/oscar-jung/kiwoom-quant/pkg/kiwoom"
)

func main() {
	c, err := kiwoom.NewClient("./.env")
	if err != nil {
		panic(err)
	}
	err = c.FetchToken()
	if err != nil {
		panic(err)
	}
	log.Printf("FETCHED TOKEN: %v", c.AccessToken)
	err = c.RevokeToken()
	if err != nil {
		panic(err)
	}
	leaders, err := c.GetVolumeLeaders()
	if err != nil {
		panic(err)
	}
	fmt.Println(leaders)
}
