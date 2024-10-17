package main

import "flag"

func main() {
	listen := flag.String("l", "[::]:10080", "listen host")
	database := flag.String("d", "judge.db", "database path")
	flag.Parse()
	s := server{
		Listen:   *listen,
		Database: *database,
	}
	s.init()
	s.serve()
}
