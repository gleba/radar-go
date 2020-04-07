import _ from 'lodash'
import { keep } from '~/back/keeper/entry'

const tokenBaseProps = [
  'id',
  'slug',
  'name',
  'symbol',
  'cmc_rank',
  'date_added',
]

export function getTune() {
  const o = keep.ns.daily.ram
  const r = []
  for (const id in o) {
    const tails = o[id].tails
    const coin = keep.ns.coin.ram[id]
    const pulse = keep.ns.pulse.ram[id]
    delete pulse._id
    const ex = keep.ns.market.ram[id].markets.map((m) =>
      _.pick(m, ['name', 'id'])
    )
    if (tails && ex) {
      r.push(
        Object.assign(
          {
            tails,
            pulse,
            ex,
          },
          _.pick(coin, ...tokenBaseProps)
        )
      )
    }
  }
}

export function getSlug(id) {
  return Object.assign(
    {
      tails: keep.ns.daily.ram[id]?.tails,
    },
    _.pick(keep.ns.coin.ram[id], 'quote', ...tokenBaseProps)
  )
}
export function getIds() {
  const o = keep.ns.coin.ram
  const r = {} as any
  for (const id in o) {
    const coin = o[id]
    const target = !!keep.ns.daily.ram[id]
    r[id] = Object.assign({ target }, _.pick(coin, ...tokenBaseProps))
  }
  return r
}
