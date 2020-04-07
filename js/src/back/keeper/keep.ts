import { makeCollection, MCollection } from '~/back/keeper/db/makeCollection'
import { openMongo } from '~/back/keeper/db/openMongo'
import A from 'alak'

export function makeKeeper() {
  const namespaces = ['daily', 'pulse', 'coin', 'market']
  const collections = {} as {
    daily: MCollection<CollDailyItem>
    pulse: MCollection<CoinPulse>
    coin: MCollection<CollCoinItem>
    market: MCollection<CollMarketItem>
  }
  const ready = A()
  openMongo().then(() => {
    namespaces.forEach((n) => (collections[n] = makeCollection(n)))
    ready.resend()
  })
  return {
    ns: collections,
    ready,
  }
}
