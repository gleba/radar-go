import { La, qubit, stored } from 'lasens'
import { LaStore } from '@store/frontStore'
import { api } from '@store/api/http'
import { BTC } from '~/front/common/consts'

const zeroLimits = () => Object.create({ min: 0, max: 0 })

export class Scan {
  capMinMax: MinMax = zeroLimits()
  @stored capLV = BTC
  @stored capValue: number
  volMinMax: MinMax = zeroLimits()
  @stored volLV = BTC
  @stored volValue: number

  @qubit markets: KV<string>
  @stored selectedMarkets: KV<boolean> = {}
  actions({ a, atoms, actions }: La<Scan, LaStore>) {
    api.get('limits').then((mm: CVLimits) => {
      a.capLV.up((lv) => {
        const mc = mm[lv].cap
        a.capMinMax(mc)
        a.capValue(mc.min * 2)
      })
      a.volLV.up((lv) => {
        const mv = mm[lv].vol
        a.volMinMax(mv)
        a.volValue(mv.min * 3)
      })
    })
    api.get('markets').then(a.markets)
    return {}
  }
}
