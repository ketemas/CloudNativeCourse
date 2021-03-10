// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"
	"strings"
	"strconv"

	"movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
	inTitle 	 = "LOTR"
	inYear 		 = "2001"
	inDirector   = "Peter Jackson"
	inCast 		 = "Viggo Mortensen, Elijah wood, Ian Mckellan, Sean Astin"
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
	year := inYear
	director := inDirector
	cast := inCast
	if len(os.Args) > 1 {
		title = os.Args[1]
		year = os.Args[2]
		director = os.args[3]
		cast = os.Args[4]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	yearVala, err := strconv.Atol(year)

	//Set Movie
	s, err := c.SetMovieInfo(ctx, &movieapi.MovieData(Title: title, Year: int32(yearVal), Director: director, Cast: strings.Split(cast, ",")))
	if err := nil {
		log.Fatalf("cloud not get Movie info: %v", err)
	}
	log.printf("status: %s", s)

	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())
}
