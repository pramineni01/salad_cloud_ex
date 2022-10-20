package main

import (
	"log"
	"os"

	"github.com/pramineni01/salad_cloud_ex/processor"
	"github.com/pramineni01/salad_cloud_ex/utils"
)

const(
	source = "data.salad.com"
	port = uint(5000)
)

func main() {
	// create processor
	 p, err := processor.NewProcessor(source, port)
	 if err != nil {
		log.Fatal("Failed to create processor: ", err)
		os.Exit(1)
	 }

	 //get context and initiate
	 ctx, cancel := utils.GetContext(nil)
	 defer cancel()
	 
	 p.Process(ctx)
}