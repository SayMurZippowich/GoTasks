package main

import "fmt"

type AppendErr string

func (err AppendErr) Error() string {
	return fmt.Sprintf("%s", err)
}
