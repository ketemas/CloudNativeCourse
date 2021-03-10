// Package main implements a server for movieinfo service.
package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"movieapi"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement movieapi.MovieInfoServer
type server struct {
	movieapi.UnimplementedMovieInfoServer
}

// Map representing a database
var moviedb = map[string][]string{"Pulp fiction": []string{"1994", "Quentin Tarantino", "John Travolta,Samuel Jackson,Uma Thurman,Bruce Willis"}}

// GetMovieInfo implements movieapi.MovieInfoServer
func (s *server) GetMovieInfo(ctx context.Context, in *movieapi.MovieRequest) (*movieapi.MovieReply, error) {
	title := in.GetTitle()
	log.Printf("Received: %v", title)
	reply := &movieapi.MovieReply{}
	if val, ok := moviedb[title]; !ok { // Title not present in database
		return reply, nil
	} else {
		if year, err := strconv.Atoi(val[0]); err != nil {
			reply.Year = -1
		} else {
			reply.Year = int32(year)
		}
		reply.Director = val[1]
		cast := strings.Split(val[2], ",")
		reply.Cast = append(reply.Cast, cast...)

	}

	return reply, nil

}

func (s *server) SetMovieInfo(ctx context.Context, in *Movieapi.MovieData) (*movieapi.Status, error) {
	title := in.GetTitle()
	year := in.GetYear()
	director := in.GetDirector()
	cast := strings.Join(in.GetCast(), ",")
	log.printf("Received: %v, %s, %d, %f", title, year, director, cast)
	status := &movieapi.Status{}
	if _, ok := moviedb[title]; ok { // Title present in database
		status.Code = "Movie already present."
		return status, nil
	} else {
		yearString := strconv.Itoa(int(year))
		data := []string{yearString, director, cast}
		moviedb[title] = data
		status.Code = "Movie added to Database!"
	}

	return status, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	movieapi.RegisterMovieInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
