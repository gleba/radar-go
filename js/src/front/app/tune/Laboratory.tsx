import { h } from 'preact'
import { useAtom } from '@store/preact'
import { store } from '@store/frontStore'
import { useEffect, useState } from 'preact/hooks'
import A from 'alak'
import { api } from '@store/api/http'
import { tailGroups, tailNames } from '~/shared/tailmap'
import { clearLine } from 'readline'


const selectedAtom = A()
const coefficientsAtom = A.setWrapper(raw => {
  return raw
})
selectedAtom.up(id =>
  api.get('alive?' + id).then(coefficientsAtom)
)


const aUsdMode = A(true)
const aBtcMode = A()
const aNowMode = A()
const aLocAtom = A(true).setWrapper(() => !aLocAtom.value)
const aLabelAtom = A(true).setWrapper(() => !aLabelAtom.value)
const aFmt = A.from(aLocAtom).some(v => v ? ru : eu)
const aLV = A()
const aVolFormat = A.from(aLV, aFmt, aLabelAtom).some((lv, loc, label) => {
  return [loc, {
    style: label ? 'currency' : 'decimal',
    currencyDisplay: 'symbol',
    currency: lv
  }]
})
const aPriceFormat = A.from(aLV, aFmt, aLabelAtom).some((lv, loc, label) => {
  return [loc, {
    style: label ? 'currency' : 'decimal',
    currencyDisplay: 'symbol',
    currency: lv,
    maximumFractionDigits: lv == BTC ? 8 : 2
  }]
})


aBtcMode.upSome(() => {
  aUsdMode(null)
  aLV(BTC)
})
aUsdMode.upSome(() => {
  aBtcMode(null)
  aLV(USD)
})

function AtomicButton({ atom, name }) {
  const selected = useAtom(atom)
  let cssClass = 'button button-deep button-small'
  if (!selected) cssClass = cssClass + ' button-clear'
  return <a style='margin-right:5px'
            onClick={atom}
            className={cssClass}>
    {name}
  </a>
}

function FmtButton() {
  let cssClass = 'button button-deep button-small'
  // if (!selected) cssClass = cssClass + ' button-clear'
  return <a style='margin-right:5px'
            // onClick={atom}
            className={cssClass}>
    {name}
  </a>
}

function Selected({ selected }) {
  const co = useAtom(coefficientsAtom)
  const LV = useAtom(aLV)
  const volForm = useAtom(aVolFormat)
  const priceForm = useAtom(aPriceFormat)
  console.log(volForm)
  console.log(priceForm)

  const tails = co?.k.tails
  const fmt = (v, isVol) =>
    v > 100000000 ? v.toFixed(0) : isVol ? v.toLocaleString(...volForm) : v.toLocaleString(...priceForm)

  return (
    <div>
      <div>
        <AtomicButton atom={aNowMode} name={'NOW DIF'}/>
        <AtomicButton atom={aUsdMode} name={'USD'}/>
        <AtomicButton atom={aBtcMode} name={'BTC'}/>
        <AtomicButton atom={aLocAtom} name={'FMT'}/>
        <AtomicButton atom={aLabelAtom} name={'LABEL'}/>
      </div>
      <table>
        <thead>
        <tr>
          <th>Value</th>
          {days.map((i, v) => <th key={v}>{i}</th>)}
        </tr>
        </thead>
        <tbody>
        {Object.keys(tailGroups.volumes).map(k => {
          const kid = tailGroups.volumes[k]
          return (
            <tr>
              <td>{k}</td>
              {tails && tails.map(d => <td class='cell-price'>{fmt(d[LV][kid], true)}</td>)}
            </tr>
          )
        })}

        {Object.keys(tailGroups.prices).map(k => {
          const kid = tailGroups.prices[k]
          return (
            <tr>
              <td>{k}</td>
              {tails && tails.map(d => <td class='cell-price'>{fmt(d[LV][kid])}</td>)}
            </tr>
          )
        })}

        {Object.keys(tailGroups.volatile).map(k => {
          const kid = tailGroups.volatile[k]
          return (
            <tr>
              <td>{k}</td>
              {tails && tails.map(d => <td class='cell-price'>{d[LV][kid].toFixed(0)}</td>)}
            </tr>
          )
        })}

        </tbody>
      </table>
    </div>
  )
}
