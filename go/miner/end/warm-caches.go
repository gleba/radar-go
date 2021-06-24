package end

import "radar.cash/core/heat"

func WarmCaches() {
	restoreMarkets()
	//heat.RestorePulse()
	heat.RestoreQuota()
	heat.RestoreDaily()
}
