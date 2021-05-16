package main

type config struct {
	Mongo mongoConfig
}

type mongoConfig struct {
	URI string `required:"true"`
	DB  string `required:"true"`
}
