package main

import "fmt"

type AppendErr string

func (err AppendErr) Error() string {
	return fmt.Sprint(err)
}
