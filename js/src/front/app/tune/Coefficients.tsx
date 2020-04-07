import { h } from 'preact'
import { useAtom } from '@store/preact'
import { store } from '@store/frontStore'
import { useEffect, useState } from 'preact/hooks'
import A from 'alak'
import { api } from '@store/api/http'


const selectedAtom = A()
const coefficientsAtom = A.setWrapper(raw => {
  return raw
})
selectedAtom.up(id =>
  api.get('alive?' + id).then(coefficientsAtom)
)

function Selected({ selected }) {
  const co = useAtom(coefficientsAtom)
  return <table>
    <thead>
    <tr>
      <th>Days</th>
      <th>BTC</th>
      <th>USD</th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td>2</td>
      <td>4324324</td>
      <td>4234234</td>
    </tr>
    </tbody>
  </table>
  // return <pre>{JSON.stringify(co, null, 2)}</pre>
}

function Item({ i }) {
  const select = () => selectedAtom(i.id)
  return <div onClick={select} className='k-item'>{i.symbol} {i.name} </div>
}

export function Coefficients() {
  const results = useAtom(store.atoms.token.results)
  const search = useAtom(store.atoms.token.search)
  const size = useAtom(store.atoms.token.targetSize)
  const selected = useAtom(selectedAtom)
  const onSearch = e =>
    store.atoms.token.search(e.target.value)
  console.log({ selected })
  if (selected) {
    return <Selected selected={selected}/>
  } else {
    return (
      <div>
        <div>Отслеживается целей: {size}</div>
        <input placeholder='Поиск криптоактива' className='up-style' value={search} onInput={onSearch}/>
        {results.map(i => <Item key={i.id} i={i}/>)}
      </div>
    )
  }
}
