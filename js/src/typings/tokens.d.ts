///collections

declare interface CollCoinItem extends CoinInfo {
  quote: LV<Quote>
}

declare interface CollDailyItem {
  id: number
  day: string
  tails: Days
}

declare interface CollMarketItem {
  id: number
  markets: Market[]
}

declare interface Market {
  id: number
  name: string
  slug: string
  pairs: MarketPair[]
}

declare interface MarketPair {
  pair: string
  url: string
}

declare interface CoinInfo {
  id: number
  slug: string
  name: string
  symbol: string
  cmc_rank: string
  date_added: string
}
declare interface TokenIds extends CoinInfo {
  target: boolean
}

declare type LiberalValue = string
declare type LV<T> = {
  [k: string]: T
  BTC: T
  USD: T
}

declare type Days = LV<number[]>[]
declare type Quote = {
  last_updated: string
  market_cap: number
  percent_change_1h: number
  percent_change_24h: number
  percent_change_7d: number
  price: number
  volume_24h: number
}

declare interface TokenTail extends CoinInfo {
  tails: Days
  quote: LV<Quote>
}

declare interface RatedToken extends CoinInfo {
  rates: any
}

declare interface TokenTune extends CoinInfo {
  tails: Days
  pulse: CoinPulse
}

declare interface CoinPulse {
  MarketCapBTC: number
  MarketCapUSD: number
  PriceBTC: number
  PriceUSD: number
  Time: string
  VolumeBTC: number
  VolumeUSD: number
}
