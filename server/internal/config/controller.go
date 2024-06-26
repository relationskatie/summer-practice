package config

type Controller struct {
	BindAddres string `config:"bind-address, short=a"`
	BindPort   int    `config:"bind-port", short=b"`
}
