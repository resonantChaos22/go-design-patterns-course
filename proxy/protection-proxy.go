package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age < 18 {
		fmt.Println("Driver is too young to drive")
		return
	}

	c.car.Drive()
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		car:    Car{},
		driver: driver,
	}
}

func TestProtectionProxy() {
	car := NewCarProxy(&Driver{
		Age: 18,
	})
	car.Drive()
}
