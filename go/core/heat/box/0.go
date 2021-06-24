package box

//go:generate go run github.com/a8m/syncmap -pkg box -name Pulses map[uint32]sol.CoinPulse
//go:generate go run github.com/a8m/syncmap -pkg box -name Quotes map[uint32]sol.CoinQuote
//go:generate go run github.com/a8m/syncmap -pkg box -name Markets map[uint32]sol.CoinMarkets
//go:generate go run github.com/a8m/syncmap -pkg box -name DailyConsts map[uint32]sol.DailyConst
