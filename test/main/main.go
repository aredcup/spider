package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var rule = `<dd class="job_bt">[^<]*<h3 class="description">职位描述：</h3>[^<]*<div class="job-detail">([^</d]+)</div>[^<]*</dd>`

func main() {

}

func r() {
	fmt.Println("start to regexp")
	f, err := os.Open("./regexp")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)

	matches := regexp.MustCompile(rule).FindAllSubmatch(b, -1)

	for _, match := range matches {
		fmt.Printf("%s \n", match)
	}
}
