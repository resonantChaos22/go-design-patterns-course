package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type internalDatabase struct {
	capitals map[string]int
}

func (db *internalDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var internalDB *internalDatabase

func GetSingletonDB() Database {
	once.Do(func() {
		fmt.Println("Initializing Database")
		db := internalDatabase{
			capitals: map[string]int{
				"Delhi":    12435323,
				"Seoul":    35432343,
				"New York": 53564454,
			},
		}
		internalDB = &db
	})
	return internalDB
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}

	return d.dummyData[name]
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDB().GetPopulation(city)
	}

	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}

	return result
}

func TestSingleton() {
	fmt.Println(GetSingletonDB().GetPopulation("Delhi"))
	fmt.Println(GetSingletonDB().GetPopulation("Seoul"))

	//	directly depends
	fmt.Println(GetTotalPopulation([]string{"Delhi", "Seoul", "New York"}))
	UnitTest()
}

func UnitTest() {
	names := []string{"alpha", "gamma"}

	//	dependency inversion
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
