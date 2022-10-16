package main

import (
	"log"
	"os"

	"github.com/pramineni01/salad_cloud_ex/processor"
	"github.com/pramineni01/salad_cloud_ex/utils"
)

func main() {
	// fetch tcp source and port
	// if invalid input, print error and return
	// hardcoded for now
	source := "data.salad.com"
	port := uint(5000)

	// create processor and connect
	 p, err := processor.NewProcessor(source, port)
	 if err != nil {
		log.Fatal("Failed to create processor: ", err)
		os.Exit(1)
	 }

	 //get context and initiate
	 ctx := utils.GetContext(nil)
	 
	 p.Process()
}