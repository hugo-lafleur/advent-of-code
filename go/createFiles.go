package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	date := strings.Split(args[0], "/")
	year, day := date[0], date[1]

	err := os.MkdirAll(year+"/"+day, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll("../inputs/"+year+"/"+day, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	template, err := os.ReadFile("template.txt")
	if err != nil {
		log.Fatalln(err)
	}

	str := string(template)
	str = strings.ReplaceAll(str, "2015", year)
	str = strings.ReplaceAll(str, "day01", day)

	f, err := os.Create(year + "/" + day + "/" + day + ".go")
	if err != nil {
		log.Fatalln(err)
	}

	f.Write([]byte(str))
	f.Close()

	_, err = os.Create("../inputs/" + year + "/" + day + "/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Create("../inputs/" + year + "/" + day + "/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

}
