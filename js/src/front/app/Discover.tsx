import { h } from 'preact'
import { store } from '@store/frontStore'
import { AtomicButton } from '~/front/app/components/AtomicButton'
import A from 'alak'
import { useAtom } from '@store/preact'
import { start, wake } from '~/front/entry'

const isAllAtom = A(false).setWrapper((v) => !isAllAtom.value)
const resultsAtom = A([])
const sizeAtom = A(0)

wake.match('discover', (_) => {
  sizeAtom.from(isAllAtom, store.atoms.discover.ids).some((all, ids) => {
    const a = Object.values(ids)
    const t = all ? a : a.filter((i) => i.target)
    return t.length
  })
  resultsAtom
    .from(isAllAtom, store.atoms.discover.results)
    .some((all, results) => (all ? results : results.filter((i) => i.target)))
})

export function Discover() {
  const results = useAtom(resultsAtom)
  const search = useAtom(store.atoms.discover.search)
  const size = useAtom(sizeAtom)
  const onSearch = (e) => store.atoms.discover.search(e.target.value)
  return (
    <div className="container discover">
      <AtomicButton atom={isAllAtom} name={'all tokens'} />
      <input
        value={search}
        onInput={onSearch}
        placeholder={`search in ${size} assets`}
        className="up-style"
      />
      {results.map((i) => (
        <Item key={i.id} i={i} />
      ))}
    </div>
  )
}

function Item({ i }) {
  const select = () => store.actions.token.select(i)
  const className = i.target ? 'item ' : 'item all'
  return (
    <div onClick={select} className={className}>
      {i.symbol} {i.name}{' '}
    </div>
  )
}
