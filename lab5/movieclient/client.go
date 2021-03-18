// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"lab6/lab5/movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())
	cancel()

	// Adding new movie details to database
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	r2, err := c.SetMovieInfo(ctx, &movieapi.MessageData{
		Title:    "Avengers: Endgame",
		Year:     2019,
		Director: "Russo Brothers",
		Cast:     []string{"Chris Evans, Chris Hemsworth, RDJ, etc."}})
	if err != nil {
		log.Fatalf("Couldn't add movie info %s", err)
	}
	log.Printf("Addition of the Movie to the Datase was %s \n", r2.Code)
	cancel()

	// Trying to retrieve the details of the new movie added
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	r, err = c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: "Avengers: Endgame"})
	if err != nil {
		log.Fatalf("Couldn't get movie info: %v", err)
	}
	log.Printf("Movie Info for Avengers: Endgame %d %s %v", r.GetYear(), r.GetDirector(), r.GetCast())
	cancel()
}