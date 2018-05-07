package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"math/big"
	"encoding/json"
	"strings"
)

type Tour struct {
	Name,  string
}
func main() {
	//一定是json格式的么？ 如何把html转成json？
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := cpntentFromServer(url)

	tours := toursFromJson(content)

	for _, tour := range tours {
		//这句什么意思
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
		//[问题1， 怎么就给转成价格了？？]
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func cpntentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	//readall这个是干什么的
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) [] Tour {


	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	fmt.Println(decoder)
	_, err := decoder.Token()
	checkError(err)
	//decode 在干嘛
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}