import { useAtom } from '@store/preact'
import { h } from 'preact'

export function AtomicButton({ atom, name }) {
  const selected = useAtom(atom)
  let cssClass = 'button button-deep button-small'
  if (!selected) cssClass = cssClass + ' button-clear'
  return <a style='margin-right:5px'
            onClick={atom}
            className={cssClass}>
    {name}
  </a>
}
