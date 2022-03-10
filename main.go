package main

import (
	"context"
	"flag"
	"flow-event-fetcher/config"
	"fmt"
	"time"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

// examples
// go run main.go --type A.e4cf4bdc1751c65d.PackNFT.Revealed
// go run main.go --type A.e4cf4bdc1751c65d.PackNFT.Revealed --range 250 --env mainnet

const (
	DEFAULT_RANGE uint64 = 250 // 250 max
	DEFAULT_ENV   string = "mainnet"
)

func main() {
	eventType := flag.String("type", "test", "")
	rng := flag.Uint64("range", DEFAULT_RANGE, "")
	env := flag.String("env", DEFAULT_ENV, "")
	flag.Parse()

	if *eventType == "" {
		fmt.Println("fully qualified event type required")
		return
	}

	config.InitConf(*env)
	fetchEvents(*eventType, *rng)
}

func fetchEvents(eventType string, blockRange uint64) {
	ctx := context.Background()
	c, err := client.New(config.Conf.FlowAccessNode, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	for true {
		fmt.Println("\n")

		// get latest block height
		latestBlock, err := c.GetLatestBlock(ctx, true)
		if err != nil {
			panic(err)
		}
		latestBlockHeight := latestBlock.Height

		// specify height range
		startBlockHeight := latestBlockHeight - DEFAULT_RANGE

		fmt.Println(fmt.Printf("start height: %v", startBlockHeight))
		fmt.Println(fmt.Printf("end height: %v", latestBlockHeight))

		// todo: in order to get all relevant events, we must query all blocks
		// in height range - once per event type. this query returns all events grouped
		// by block in height range and *only* includes events that match the query type
		// this means that blocks will be appended multiple times
		// there is probably a better way to do this
		blockEvents := []client.BlockEvents{}
		bes, err := c.GetEventsForHeightRange(ctx, client.EventRangeQuery{
			Type:        eventType,
			StartHeight: startBlockHeight + 1,
			EndHeight:   latestBlockHeight,
		})
		if err != nil {
			panic(err)
		}
		blockEvents = append(blockEvents, bes...)

		fmt.Println(fmt.Printf("events processed: %v", len(blockEvents)))

		var eventsFound []flow.Event
		for _, b := range blockEvents {
			for _, e := range b.Events {
				if e.Type == eventType {
					eventsFound = append(eventsFound, e)
				}
			}
		}

		if len(eventsFound) > 0 {
			fmt.Println(fmt.Printf("%v", eventsFound))
		} else {
			fmt.Println("no matching events found")
		}

		time.Sleep(time.Second * 2)
	}
}
