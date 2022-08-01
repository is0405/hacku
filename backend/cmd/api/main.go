package main

import (
	"flag"
	"log"
	"os"

	server "github.com/is0405/hacku"
)

func main() {
	var databaseDatasource string
	var port int

	flag.StringVar(&databaseDatasource, "databaseDatasource", "root:password@tcp(localhost:3306)/hacku", "Should looks like root:password@tcp(hostname:port)/dbname")
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	if err := s.Init(databaseDatasource); err != nil {
		log.Fatal(err)
	}
	s.Run(port)
}
