package main

type Mesurament struct{
	City string
	Average float64 `json:"average"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Count int64 `json:"count"`
	Sum float64 `json:"sum"`
}