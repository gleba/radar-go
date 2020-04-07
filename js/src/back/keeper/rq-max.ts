import frequencyDistribution from 'frequency-distribution'
import { keep } from '~/back/keeper/entry'
import { tgl } from '~/front/common/tails'
import { BTC, dayN, days, USD } from '~/front/common/consts'
import { createPrivateKey } from 'crypto'

const pro = {
  update(v) {
    this.min = Math.round(Math.min(this.min, v))
    this.max = Math.round(Math.max(this.max, v))
  },
  fin() {
    this.max = Math.round(this.max / 8)
  },
}
const minmaxLimits = () => {
  return {
    min: Number.MAX_SAFE_INTEGER,
    max: 0,
    __proto__: pro,
  }
}

const capVolLimits = () => {
  return {
    cap: minmaxLimits(),
    vol: minmaxLimits(),
  } as any
}

export function getLimits() {
  let lva = [BTC, USD]
  let res = {
    BTC: capVolLimits(),
    USD: capVolLimits(),
  }
  Object.keys(keep.ns.daily.ram).forEach((id: any) => {
    if (id == '1027') {
      return
    }
    lva.forEach((lv) => {
      let q = keep.ns.coin.ram[id].quote
      res[lv].cap.update(q[lv].market_cap)
      res[lv].vol.update(q[lv].volume_24h)
    })
  })

  lva.forEach((lv) => {
    res[lv].cap.fin()
    res[lv].vol.fin()
  })
  res.BTC.vol.min = 5
  res.BTC.cap.min = 100
  return res
}

export function getVolatile() {
  console.log('getVolatile')
  let va = []
  let max = 0
  let min = Number.MAX_SAFE_INTEGER
  let z = 0
  const countZero = (v) => {
    let lt = v.toFixed(8).split('.')[1]
    let l = lt.length
    let i = 0
    let cz = 0
    while (i < l) {
      if (lt[i] == '0') cz++
      else i = l
      i++
    }
    return cz
  }
  Object.keys(keep.ns.daily.ram).forEach((id) => {
    let tails = keep.ns.daily.ram[id].tails
    if (tails) {
      let d = tails[3]
      if (!d) return
      if (!d.BTC) return
      let v = d.BTC[tgl.volatile.mean]
      let p = d.BTC[tgl.prices.mean]
      if (!v || !p) return
      let vv = Math.round((v / p) * 10000) / 10
      max = Math.max(vv, max)
      min = Math.min(vv, min)
      let k = Math.pow(10, Math.round(countZero(p)))
      let vk = (k / v) * 100000000
      va.push([
        vk,
        vk.toFixed(8),
        v,
        k,
        v.toFixed(8),
        p.toFixed(8),
        keep.ns.pulse.ram[id].PriceBTC.toFixed(8),
        Math.round(p * 100000000),
        keep.ns.coin.ram[id].slug,
      ])
    }
  })
  va.sort((a, b) => a[0] - b[0])
  // min = min.toFixed(3)
  // max = max.toFixed(3)
  console.log(va)
  let x = frequencyDistribution(va)
  // console.log()
  return va
}
