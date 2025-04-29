package cliente

import "time"

type Cliente struct {
	Id        int
	Name      string
	BirthDate time.Time
	Address   *Address
	UpdatedAt time.Time
}

func (c *Cliente) AddAddress(a *Address) {
	c.Address = a
	c.changeUpdatedTime()
}

func (c *Cliente) changeUpdatedTime() {
	c.UpdatedAt = time.Now()
}

type Address struct {
	Street     string
	PostalCode string
	City       string
}
