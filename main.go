package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	file, err := os.Open("temperature.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mesuraments := make(map[string]Mesurament)
	for scanner.Scan(){
		row := scanner.Text()
		semicolon := strings.Index(row, ";")
		city := row[0:semicolon]
		temperature, err := strconv.ParseFloat(strings.TrimSpace(row[(semicolon+1):]), 64)
		if err != nil{
			fmt.Println(err)
			panic(err)
		}
		var mesurament Mesurament
		if _, ok := mesuraments[city]; !ok {
			mesurament = Mesurament{
				City: city,
				Max: temperature,
				Min: temperature,
				Average: temperature,
				Count: 1,
			}
		} else {
			mesurament = mesuraments[city]
			mesurament.Count += 1
		}
		
		if temperature < mesurament.Min {
			mesurament.Min = temperature
		}
		if temperature > mesurament.Max {
			mesurament.Max = temperature
		}
		mesurament.Sum += temperature 
		mesurament.Average = mesurament.Sum / float64(mesurament.Count)
		mesuraments[city] = mesurament
	}

	cities := make([]string, 0, len(mesuraments))
	for key, _ := range mesuraments {
		cities = append(cities, key)
	}
	sort.Strings(cities)
	
	fmt.Printf("{")
	for _, value := range cities {
		mesurament := mesuraments[value]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",value, mesurament.Min, mesurament.Average, mesurament.Max)
	}
	fmt.Printf("}\n")

	
	fmt.Errorf("%v", err)
	errors.New()


}