package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	//
	var filePath string
	flag.StringVar(&filePath, "file", "", "input file path")

	if filePath == "" {
		// empty file error
		return
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// error
		return
	}

	in := input{}
	if err := json.Unmarshal(data, &in); err != nil {
		// unmarshal error
		return
	}

	/*
		input := []Rect{
			Rect{x: 100, y: 100, w: 250, h: 80},
			Rect{x: 120, y: 200, w: 250, h: 150},
			Rect{x: 140, y: 160, w: 250, h: 100},
			Rect{x: 160, y: 140, w: 350, h: 190},
		}

	*/
	/* 	f := NewFinder()
	   	r := f.find(in.Rect)
	   	log.Printf("%+v", r) */
}
