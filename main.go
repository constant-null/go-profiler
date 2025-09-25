package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/go-profiler/api"
	"github.com/go-profiler/profile"
)

func main() {
	f, _ := os.Open("alloc.pprof")
	r, _ := gzip.NewReader(f)
	d, _ := ioutil.ReadAll(r)
	profile.Parse(d)

	api.Serve(":3030")
}
