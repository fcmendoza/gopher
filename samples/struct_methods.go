package main

import (
	"fmt"
	"time"
)

type Employee struct {
	FirstName, LastName string
}

func fullName(firtname string, lastName string) (fullname string) {
	fullname = firtname + " " + lastName
	return
}

func (e Employee) FullName() string { // attaching a method to the Employee struct
	return e.FirstName + " " + e.LastName
}

// DasError is an error implementation that includes a time and message.
type DasError struct {
	When time.Time
	What string
}

func (de DasError) Error() string {
	return fmt.Sprintf("%v: %v", de.When, de.What)
}

func (de DasError) Errore() string {
	return fmt.Sprintf("Errore: %v: %v", de.When, de.What)
}

/*
Because we're returning the bult-in 'error' type the retuning type must implement the Error() method. 
In other words the returning type must implment 'error' for this to work. That is DasError type must have an Error method.
Otherwise the program will fail with the following error:
cannot use DasError literal (type DasError) as type error in return argument:
	DasError does not implement error (missing Error method)
*/
func oops() error { // we're retuning `error` here, therefore DasError must implement `error`, therefore DasError must have an Error method
	return DasError {
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func ups() DasError {
	de := DasError { time.Date(2019, 3, 15, 22, 30, 0, 0, time.UTC), "The file failed.", }
	return de
}

func main() {
	e := Employee {
		FirstName: "Jon",
		LastName: "Connor",
	}
	fmt.Println(fullName(e.FirstName, e.LastName))
	fmt.Println(e.FullName())

	err := oops();
	if err != nil {
		fmt.Println(err)
	}

	dae := ups()
	msg := dae.Errore()
	fmt.Println(msg)
	fmt.Println(dae)
}
