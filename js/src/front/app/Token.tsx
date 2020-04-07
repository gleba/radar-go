import { h, Fragment } from 'preact'
import { xAtom, store } from '@store/frontStore'
import { useAtom } from '@store/preact'
import { AtomicButton } from '~/front/app/components/AtomicButton'
import { start } from '~/front/entry'
import A from 'alak'
import { BTC, days, eu, ru, USD } from '~/front/common/consts'
import { tgl } from '~/front/common/tails'
import { useCallback } from 'preact/hooks'

export function Token() {
  const goHome = (e) => {
    e.preventDefault()
    store.atoms.routes.current('/')
  }
  const now = useAtom(aNowMode)
  return (
    <div className="container">
      <div>
        <AtomicButton atom={aNowMode} name={'NOW % DIF'} />
        <AtomicButton atom={aUsdMode} name={'USD'} />
        <AtomicButton atom={aBtcMode} name={'BTC'} />
        {now && (
          <Fragment>
            <AtomicButton atom={aLocAtom} name={'FMT'} />
            <AtomicButton atom={aLabelAtom} name={'LABEL'} />
          </Fragment>
        )}
      </div>
      <Selected />
    </div>
  )
}

const aUsdMode = A(true)
const aBtcMode = A()
const aNowMode = A(false).setWrapper(() => !aNowMode.value)
const aLocAtom = A(true).setWrapper(() => !aLocAtom.value)
const aLabelAtom = A(true).setWrapper(() => !aLabelAtom.value)
const aFmt = A.from(aLocAtom).some((v) => (v ? ru : eu))
const aLV = A()

const aVolFormat = A.from(aLV, aFmt, aLabelAtom).some((lv, loc, label) => {
  return [
    loc,
    {
      style: label ? 'currency' : 'decimal',
      currencyDisplay: 'symbol',
      currency: lv,
    },
  ]
})
const aPriceFormat = A.from(aLV, aFmt, aLabelAtom).some((lv, loc, label) => {
  return [
    loc,
    {
      style: label ? 'currency' : 'decimal',
      currencyDisplay: 'symbol',
      currency: lv,
      maximumFractionDigits: lv == BTC ? 8 : 2,
    },
  ]
})

const pct = (target, from) => Math.round((from / target) * 100) - 100 + ' %'
aBtcMode.upSome(() => {
  aUsdMode(null)
  aLV(BTC)
})
aUsdMode.upSome(() => {
  aBtcMode(null)
  aLV(USD)
})

function Selected() {
  // const co = fromAtom.token.coefficients
  const LV = useAtom(aLV)
  console.log(LV)
  const volForm = useAtom(aVolFormat)
  const priceForm = useAtom(aPriceFormat)
  const now = useAtom(aNowMode)
  const selected = xAtom.token.selected
  const fmt = (v, isVol) =>
    v > 100000000
      ? v.toFixed(0)
      : isVol
      ? v.toLocaleString(...volForm)
      : v.toLocaleString(...priceForm)
  const maybePct = useCallback(
    (v, isVol?) =>
      now
        ? isVol
          ? pct(v, selected.quote[LV].volume_24h)
          : pct(v, selected.quote[LV].price)
        : fmt(v, isVol),
    [now, LV, volForm, priceForm]
  )

  if (!selected) {
    return null
  }
  const { tails } = selected
  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>Value</th>
            {days.map((v, i) => (
              <th className={`cell-price maybe-col-${i}`} key={i}>
                {v}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {Object.keys(tgl.volumes).map((k) => {
            const kid = tgl.volumes[k]
            return (
              <tr>
                <td>{k}</td>
                {tails &&
                  tails.map((d, i) => (
                    <td className={`cell-price maybe-col-${i}`}>
                      {maybePct(d[LV][kid], true)}
                    </td>
                  ))}
              </tr>
            )
          })}

          {Object.keys(tgl.prices).map((k) => {
            const kid = tgl.prices[k]
            return (
              <tr>
                <td>{k}</td>
                {tails &&
                  tails.map((d, i) => (
                    <td className={`cell-price maybe-col-${i}`}>
                      {maybePct(d[LV][kid])}
                    </td>
                  ))}
              </tr>
            )
          })}

          {Object.keys(tgl.percentile).map((k) => {
            const kid = tgl.percentile[k]
            return (
              <tr>
                <td>{k}</td>
                {tails &&
                  tails.map((d, i) => (
                    <td className={`cell-price maybe-col-${i}`}>
                      {fmt(d[LV][kid], false)}
                    </td>
                  ))}
              </tr>
            )
          })}
          {Object.keys(tgl.volatile).map((k) => {
            const kid = tgl.volatile[k]
            return (
              <tr>
                <td>{k}</td>
                {tails &&
                  tails.map((d, i) => (
                    <td className={`cell-price maybe-col-${i}`}>
                      {d[LV][kid].toFixed(0)}
                    </td>
                  ))}
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>
  )
}
