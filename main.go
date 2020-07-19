package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/adasari/task/sl"
)

type input struct {
	Rect []sl.Rect `json:"rects"`
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "input file path")

	if filePath == "" {
		log.Printf("filepath is empty")
		return
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("failed to read to file %s: %v", filePath, err)
		return
	}

	in := input{}
	if err := json.Unmarshal(data, &in); err != nil {
		log.Printf("failed to unarmshal to file content %s: %v", filePath, err)
		return
	}

	/* in := input{
		Rect: []sl.Rect{
			{X: 100, Y: 300, W: 70, H: 180},
			{X: 150, Y: 220, W: 50, H: 70},
			{X: 160, Y: 250, W: 40, H: 50},
			{X: 250, Y: 250, W: 50, H: 100},
		},
	} */

	var input []sl.NamedRect
	for i, r := range in.Rect {
		r := r // go range problem.
		input = append(input, sl.NamedRect{
			Rect: &r,
			Name: fmt.Sprintf("%d", i+1),
		})
	}

	f := sl.NewFinder()
	r := f.FindIntersections(input)

	fmt.Printf("lenth of results: %d\n", len(r))
	for _, i := range r {
		log.Printf("%s - %+v\n", strings.Join(i.Names, ","), *i.Rect)
	}

}
