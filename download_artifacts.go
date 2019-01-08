package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
)

const WPTURL = "http://wpt1.speedcurve.com/xmlResult/"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func download(r result) {
	fmt.Println(r)
	// Create the directory with test_id
	newpath := filepath.Join(".", r.testid)
	os.MkdirAll(newpath, os.ModePerm)

	// Download the rendered screen as test_id/screen
	urldownload(newpath, r.testid+".jpg", r.screen)

	// Download the HAR file
	urldownload(newpath, r.testid+".har", r.har)

	// Download the WPT XML Result
	urldownload(newpath, r.testid+".xml", r.wptxml)

}

func urldownload(savedir string, savefile string, downurl string) {
	// Download the rendered screen as test_id/screen
	newFile, err := os.Create(filepath.Join(savedir, savefile))
	check(err)
	defer newFile.Close()

	resp, err := http.Get(downurl)
	check(err)
	defer resp.Body.Close()

	_, err = io.Copy(newFile, resp.Body)
	check(err)
}

func main() {
	fmt.Println("Parsing the Speed Curve data")
	content, err := ioutil.ReadFile("sample_export.json")
	check(err)

	jsonparser.ArrayEach(content, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		// fmt.Println(jsonparser.GetString(value, "test_id"))
		tid, e := jsonparser.GetString(value, "test_id")
		ts, e := jsonparser.GetString(value, "screen")
		th, e := jsonparser.GetString(value, "har")
		tw := WPTURL + tid + "/" // The trailing back slash is essential or 404
		check(e)

		r := result{tid, ts, th, tw}
		download(r)
	}, "tests")
}

type result struct {
	testid string
	screen string
	har    string
	wptxml string
}
