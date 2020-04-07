package df

import (
	"radar.cash/core/intel"
)

var DailyConst = flow{
	"daily",
}
var Pulse = flow{
	"pulse",
}

var Market = flow{
	"market",
}

type flow struct {
	namespace string
}
type callBack interface{}

func (c *flow) Up(cb callBack) {
	intel.Request(c.ns("init"), cb)
	intel.Subscribe(c.ns("updates.>"), cb)
}

func (c *flow) Init(cb callBack) {
	intel.Request(c.ns("init"), cb)
}

func (c *flow) UpdateItem(id uint32, item interface{}) {
	s:= make([]interface{}, 1)
	s[0] = item
	intel.Publish(c.ns("updates"), s)
}

func (c *flow) ns(path string) string {
	ns := "flow." + c.namespace
	if path != "" {
		return ns + "." + path
	}
	return ns
}
