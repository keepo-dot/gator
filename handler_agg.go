package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("agg command must have a duration, ex: 1m")
	}
	timeInterval, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}
	ticker := time.NewTicker(timeInterval)
	fmt.Printf("Fetching feeds every %s:\n", cmd.arguments[0])
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
