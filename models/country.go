package models

import "fmt"

type Country struct {
	Number int
	Name   string
	Code   string
}

func (c *Country) String() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Code)
}
