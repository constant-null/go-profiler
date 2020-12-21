package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/google/profile_storage/api"
	"github.com/google/profile_storage/profile"
)

func main() {
	f, _ := os.Open("cpu.prof")
	r, _ := gzip.NewReader(f)
	d, _ := ioutil.ReadAll(r)
	profile.Parse(d)

	api.Serve(":3030")
}
