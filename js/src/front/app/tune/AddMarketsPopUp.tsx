import { h } from 'preact'
import { mix, store, xAtom } from '@store/frontStore'
import { newRune } from '~/shared/rune'
import { start } from 'src/front/entry'
import A from 'alak'

const alphaSort = (a, b) => {
  if (a.name < b.name) return -1
  if (a.name > b.name) return 1
  return 0
}
const searchAtom = A('')
export function AddMarketsPopUp() {
  const block = (e: MouseEvent) => {
    e.preventDefault()
    e.stopImmediatePropagation()
    e.stopPropagation()
  }
  const markets = mix(
    (a) => [a.scan.markets, a.scan.selectedMarkets, searchAtom],
    (mark, sel, searchValue) => {
      if (mark) {
        let m = Object.keys(mark).map((id) => {
          return { id, name: mark[id], selected: sel[id], key: newRune(3) }
        })
        m.sort(alphaSort)

        return m.filter((t) => t.name.indexOf(searchValue) != -1)
      } else {
        return []
      }
    }
  )
  const mark = (v) =>
    store.atoms.scan.selectedMarkets.fmap((s) => {
      markets.forEach((m) => {
        s[m.id] = v
      })
      return s
    })
  const markAll = () => mark(true)
  const unmarkAll = () => mark(false)
  const onSearch = (e) => searchAtom(e.target.value)
  return (
    <div class="popup-zone">
      <div class="popup-close" onClick={store.actions.hud.clarity} />
      <div class="popup market slide-in-elliptic-top-fwd" onClick={block}>
        {/*<div className="title">Markets</div>Ю*/}
        <div className="row-stretch">
          {/*<button class="btn btn-border-rev-o btn-green">Cancel</button>*/}
          <button className="btn btn-border-rev-o btn-green" onClick={markAll}>
            Check All
          </button>
          <button
            className="btn btn-border-rev-o btn-green"
            onClick={unmarkAll}
          >
            Uncheck All
          </button>
        </div>
        <input
          onInput={onSearch}
          placeholder={`search `}
          className="up-style"
        />
        <div class="content">
          <div class="markets">
            {markets &&
              markets.map((m) => {
                return <Market market={m} key={m.key} />
              })}
          </div>
        </div>
        <div className="actions">
          <div class="row-stretch">
            {/*<button class="btn btn-border-rev-o btn-green">Cancel</button>*/}
            <button
              class="btn btn-border-rev-o btn-green"
              onClick={store.actions.hud.clarity}
            >
              Ok
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

// start.up((v) => {

// store.atoms.scan.selectedMarkets.up((v) => {})
// A.from(store.atoms.token.LV).some((v) => {
//   console.log('xx', v)
// })

// })
function Market({ market }) {
  const uid = 'xx' + market.id
  const change = (e) => {
    // console.log('change', market.id)
    // store.atoms.token.LV.fmap((v) => Math.random())
    store.atoms.scan.selectedMarkets.fmap((s) => {
      s[market.id] = !s[market.id]
      return s
    })
    store.atoms.scan.markets.resend()
  }
  return (
    <div className="checkbox-root" onClick={change}>
      <div className="checkbox path">
        <input type="checkbox" id={uid} checked={market.selected} />
        <svg viewBox="0 0 21 21">
          <path d="M5,10.75 L8.5,14.25 L19.4,2.3 C18.8333333,1.43333333 18.0333333,1 17,1 L4,1 C2.35,1 1,2.35 1,4 L1,17 C1,18.65 2.35,20 4,20 L17,20 C18.65,20 20,18.65 20,17 L20,7.99769186"></path>
        </svg>
      </div>
      <label htmlFor={uid} className="label">
        {market.name}
      </label>
    </div>
  )
}
