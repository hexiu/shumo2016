package main

import (
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	resp, err := http.Get("http://data.earthquake.cn/datashare/globeEarthquake_csn.html")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	input, err := ioutil.ReadAll(resp.Body)
	out := make([]byte, len(input))
	out = out[:]
	iconv.Convert(input, out, "gb2312", "utf-8")
	fmt.Println(out)
	ioutil.WriteFile("out.html", out, 0644)
}
