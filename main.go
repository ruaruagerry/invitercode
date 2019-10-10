package main

import (
	"log"

	"github.com/ruaruagerry/invitercode/code"
)

func main() {
	userID := "12345678"
	newcode := code.CreateInviterCode(userID)
	log.Println("newcode:", newcode)
	newuserID := code.DecodeInviterCode(newcode)
	log.Println("newuserID:", newuserID)
}
