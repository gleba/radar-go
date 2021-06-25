package _

import (
	"fmt"
	"radar.cash/core/sol"
)

func upFloor(coin sol.CoinQuote, alives []AliveTune) {
	for _, alive := range alives {
		fmt.Println(alive)
	}
}
