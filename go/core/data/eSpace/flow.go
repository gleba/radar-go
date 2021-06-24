package eSpace

import (
	"radar.cash/core/data/service"
)

var DailyConst = flow{"daily"}

var Quotes = flow{"quotes"}
var Pulses = flow{"pulses"}

var Markets = flow{"markets"}

type flow struct {
	namespace string
}
type callBack interface{}

func (c *flow) Sub(cb callBack) {
	service.NatsEncoded.Subscribe(c.ns(""), cb)
}

func (c *flow) Publish(item interface{}) {
	service.NatsEncoded.Publish(c.ns(""), item)
}

func (c *flow) ns(path string) string {
	ns := "flow." + c.namespace
	if path != "" {
		return ns + "." + path
	}
	return ns
}
