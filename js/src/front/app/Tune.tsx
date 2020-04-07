import { h } from 'preact'
import { store } from 'src/front/store/frontStore'
import { api } from '@store/api/http'
import A from 'alak'
import { tgl } from '~/front/common/tails'
import { days } from '~/front/common/consts'
import { useAtom } from '@store/preact'

const tune = A<TokenTune[]>()

// const pct = (target, from) => (from / target) * 100
//
// api.get('tune').then(tune)
// const targets = A.id('targets')
// const preinit = localStorage.getItem(targets.id)
// // console.log(preinit)
// preinit && targets(JSON.parse(preinit))
// tune.up(a => {
//   const tmap = {}
//   a.forEach(token => {
//     token.tails?.forEach((lv, day) => {
//       let rate = pct(token.pulse.VolumeBTC, lv.BTC[tailGroups.volumes.VolumeMean]) / 100
//       if (rate > 3) {
//         let ratedToken = tmap[token.slug]
//         if (!ratedToken) ratedToken = tmap[token.slug] = Object.assign(token, {
//           rates: []
//         })
//         delete ratedToken.tails
//         ratedToken.rates.push([day, rate.toFixed(1)])
//       }
//     })
//   })
//   targets(tmap)
// })
// targets.up((v, a) => {
//   localStorage.setItem(a.id,JSON.stringify(v))
// })

export function Tune() {
  // const rep = useAtom(targets)
  return (
    <div className="container">
      {/*{rep &&*/}
      {/*<div> targets {Object.values(rep).length}*/}
      {/*  {Object.values(rep).map((t: TokenTune) =>*/}
      {/*    <div>*/}
      {/*      <h2>{t.name}</h2>*/}
      {/*    </div>)*/}
      {/*  }*/}
      {/*</div>}*/}
    </div>
  )
}
