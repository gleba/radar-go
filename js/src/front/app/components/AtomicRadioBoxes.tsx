import { h } from 'preact'
import { useAtom } from '@store/preact'

export function AtomicRadioBoxes({ atom, values }) {
  const selected = useAtom(atom)
  return (
    <div>
      {values.map((v) => {
        let cssClass = 'button button-deep button-small'
        if (selected != v) cssClass = cssClass + ' button-clear'
        return (
          <a
            style="margin-right:5px"
            onClick={() => atom(v)}
            className={cssClass}
          >
            {v}
          </a>
        )
      })}
    </div>
  )
}
