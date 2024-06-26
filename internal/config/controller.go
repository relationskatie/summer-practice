package config

import "fmt"

type Controller struct {
	BindAddres string `config:"bind-address, short=a"`
	BindPort   int    `config:"bind-port", short=b`
}

func (c Controller) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", c.BindAddres, c.BindPort)
}
