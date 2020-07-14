package main

import "github.com/denisacostaq/glanguage/src"

func main() {
	s := src.NewServer(8080)
	s.Start()
}
