package vars

import (
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars/asg"
)

var Alerts asg.SyncAlerts
var ActiveAlerts asg.ActiveAlerts

var FrontDict = &wings.FrontDict{
	Rules: map[int]*wings.FrontRule{},
	Alert: map[int]*wings.FrontAlert{},
}
