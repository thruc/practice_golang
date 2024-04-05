package main

import (
	"errors"
	"fmt"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ", ")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("Age must be greater than 0"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("Name cannot be empty"))
	}

	return m
}

func main() {
	c := Customer{
		Age:  35,
		Name: "John",
	}
	err := c.Validate()
	if err != nil {
		fmt.Printf("customer is invalid: %v\n", err) // customer is invalid: <nil>
	}

	var foo *Foo
	fmt.Println(foo.Bar()) // bar
}

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}
