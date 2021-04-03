package gee

import "log"

func Cors() HandlerFunc {
	return func(c *Context) {
		log.Println("call cors middleware")
	}
}
