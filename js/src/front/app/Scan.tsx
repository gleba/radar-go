import { h } from 'preact'
import { mix, store, xAtom } from 'src/front/store/frontStore'
import { api } from '@store/api/http'
import A from 'alak'
import { tgl } from '~/front/common/tails'
import { BTC, days, dayCod, USD } from '~/front/common/consts'
import { useAtom } from '@store/preact'
import { createPrivateKey } from 'crypto'
import { Checkbox } from '~/front/app/components/Checkbox'
import { AtomicButton } from '~/front/app/components/AtomicButton'
import { AtomicRadioBoxes } from '~/front/app/components/AtomicRadioBoxes'
import { money } from '~/front/common/cur'

// const tune = A<TokenTune[]>()
//
// const pct = (target, from) => (from / target) * 100
//
// const targets = A.id('targets')
// const preinit = localStorage.getItem(targets.id)
// preinit && targets(JSON.parse(preinit))
//
// targets.up((v, a) => {
//   localStorage.setItem(a.id, JSON.stringify(v))
// })
//
// const strategies = [] as {
//   kid: number
//   did: number
//   lv: string
//   name: string
// }[]
//
// function addX(prefix, did, kid, lv) {
//   strategies.push({ name: `${prefix}${lv}-${dayCod[did]}`, kid, did, lv })
// }
//
// function addS(prefix, did, kid) {
//   addX(prefix, did, kid, BTC)
//   addX(prefix, did, kid, USD)
// }
//
// ;[2, 3].forEach((n, did) => {
//   // addS('vMean', did, tailGroups.volumes.VolumeMean)
//   // addS('vMedian', did, tailGroups.volumes.VolumeMedian)
//   // addS('vHarm', did, tailGroups.volumes.VolumeHarmonic)
//   // addS('pMean', did, tailGroups.prices.PriceMedian)
//   // addS('pMedian', did, tailGroups.prices.PriceMedian)
//   // addS('pHarm', did, tailGroups.prices.PriceHarmonic)
//   // addS('pVar', did, tailGroups.prices.PriceVariance)
//   addS('vvMean', n, tgl.volatile.mean)
//   // addS('vvHarm', did, tailGroups.volatile.VolatilityHarmonic)
//   // addS('vvMedian', did, tailGroups.volatile.VolatilityMedian)
// })
// // console.log(strategies.length)
//
// tune.up((a) => {
//   const tmap = {}
//   strategies.forEach((s) => {
//     a.forEach((token) => {
//       if (!token.tails) {
//         return
//       }
//       let vv = token.tails[s.did][s.lv][s.kid]
//       if (vv > 1 && vv < 4000) {
//         console.log(token.slug, vv.toFixed(0), dayCod[s.did])
//       }
//     })
//   })
//   // a.forEach(token => {
//   //   token.tails?.forEach((lv, day) => {
//   //     let rate = pct(token.pulse.VolumeBTC, lv.BTC[tailGroups.volumes.VolumeMean]) / 100
//   //     if (rate > 3) {
//   //       let ratedToken = tmap[token.slug]
//   //       if (!ratedToken) ratedToken = tmap[token.slug] = Object.assign(token, {
//   //         rates: []
//   //       })
//   //       delete ratedToken.tails
//   //       ratedToken.rates.push([days[day], rate.toFixed(1)])
//   //     }
//   //   })
//   // })
//   targets(tmap)
// })
// console.log(strategies)
//
// function scan() {
//   api.get('tune').then(tune)
// }
const mv = Math.exp(12)

function AtomicFilter({ label, valueAtom, limitsAtom, lvAtom }) {
  const value = useAtom(valueAtom) as number
  const limit = useAtom(limitsAtom) as any
  const lv = useAtom(lvAtom)
  // console.log({ capMinMax })
  const capMove = (e) => {
    capChange(e)
  }
  const rv = limit.max - limit
  const capChange = (event) => {
    const v = event.target.value
    const e = Math.exp(v)
    const p = e / mv
    valueAtom(limit.min + limit.max * p)
  }

  const fp = (value - limit.min) / limit.max
  const fe = mv * fp
  const logValue = Math.log(fe)

  return (
    <div className="rule">
      <div className="mini-range">
        <label key={value}>
          {label} greater than {money(value, lv)}
        </label>
        <AtomicRadioBoxes atom={lvAtom} values={[USD, BTC]} />
        <div className="minmax" key={value}>
          <div className="min">{money(limit.min, lv)}</div>
          <div className="max">{money(limit.max, lv)}</div>
        </div>
        <input
          min="0"
          max="12"
          step="0.001"
          type="range"
          value={logValue}
          onChange={capChange}
          onInput={capMove}
        />
      </div>
    </div>
  )
}

// atomic atomic-badge
// <div
//   className="atomic-badge"
//   onClick={() => store.atoms.hud.popup('markets')}
// >
//   <div className="icon-badge">
//     <div className="text">Add</div>
//     <div className="icon-button">
//       <div className="icon">+</div>
//     </div>
//   </div>
// </div>
const mInfo = A()

function MarketsFilter() {
  const sizes = mix(
    (a) => [a.scan.markets, a.scan.selectedMarkets],
    (markets, selected) => {
      if (!markets) return '-/-'
      let sel = 0
      let names = []

      Object.keys(markets).forEach((id) => {
        if (selected[id]) {
          sel++
          console.log(markets[id], id)
          names.push(markets[id])
        }
      })
      const mtotal = Object.keys(markets).length
      if (names.length > 0) {
        if (names.length > 8) mInfo(names.slice(0, 8).join(', ') + '...')
        else if (names.length == mtotal) mInfo(names.join('everyone'))
        else mInfo(names.join(', '))
      } else {
        mInfo('no one, select someone')
      }
      return `${sel}/${mtotal}`
    }
  )
  const info = useAtom(mInfo)
  const change = () => store.atoms.hud.popup('markets')
  return (
    <div class="rule">
      <label>Markets </label>
      {info}
      <div className="add-markets" onClick={change}>
        {sizes}
      </div>
    </div>
  )
}

export function Scan() {
  // const rep = useAtom(targets)
  return (
    <div className="container route">
      <AtomicFilter
        label="Capitalisation"
        valueAtom={store.atoms.scan.capValue}
        lvAtom={store.atoms.scan.capLV}
        limitsAtom={store.atoms.scan.capMinMax}
      />
      <AtomicFilter
        label="Volume 24h"
        valueAtom={store.atoms.scan.volValue}
        lvAtom={store.atoms.scan.volLV}
        limitsAtom={store.atoms.scan.volMinMax}
      />

      <MarketsFilter />
      {/*<div className="rule">*/}
      {/*  <div>*/}
      {/*    <AtomicButton atom={A()} name={'Mean'} />*/}
      {/*    <AtomicButton atom={A()} name={'Median'} />*/}
      {/*    <AtomicButton atom={A()} name={'Harmonic'} />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className="rule">*/}
      {/*  <div>*/}
      {/*    <AtomicButton atom={A()} name={'30 days'} />*/}
      {/*    <AtomicButton atom={A()} name={'45 days'} />*/}
      {/*    <AtomicButton atom={A()} name={'60 days'} />*/}
      {/*    <AtomicButton atom={A()} name={'90 days'} />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className="mini-range">*/}
      {/*  <label>Rate</label>*/}
      {/*  <input type="range" />*/}
      {/*</div>*/}
      {/*<div className="scan-con">*/}
      {/*  <button className="logo-but fxb">*/}
      {/*    <div className="label">SCAN</div>*/}
      {/*  </button>*/}
      {/*</div>*/}
      {/*{rep && (*/}
      {/*  <div>*/}
      {/*    {' '}*/}
      {/*    targets {Object.values(rep).length}*/}
      {/*    {Object.values(rep).map((t: RatedToken) => (*/}
      {/*      <div>*/}
      {/*        <h2>{t.name}</h2>*/}
      {/*        <div>*/}
      {/*          {t.rates.map((r) => (*/}
      {/*            <div>{r.join(' - ')}</div>*/}
      {/*          ))}*/}
      {/*        </div>*/}
      {/*      </div>*/}
      {/*    ))}*/}
      {/*  </div>*/}
      {/*)}*/}
    </div>
  )
}
