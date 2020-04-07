import { h } from 'preact'
import { useAtom } from '@store/preact'
import { store } from '@store/frontStore'
import { WrongHash } from '~/front/app/WrongHash'
import { Tune } from '~/front/app/Tune'
import { Discover } from '~/front/app/Discover'
import { useCallback, useState } from 'preact/hooks'
import { routesMap } from '~/router/routes'
import { Token } from '~/front/app/Token'
import { start } from '~/front/entry'
import A from 'alak'
import { Scan } from '~/front/app/Scan'

const routes: TKV<AppRoutes, any> = {
  discover: <Discover />,
  tune: <Tune />,
  scan: <Scan />,
  'wrong-hash': <WrongHash />,
  '-': <Token />,
}

const menuItems: AppRoutes[] = ['discover', 'scan']

function NavBar({ route }) {
  const links = menuItems.filter((r) => r != route)
  return (
    <div className="top-bar">
      <NavBtn way="/" label="Home" />
      {links.map((r) => (
        <NavBtn key={r} way={r} label={r} />
      ))}
    </div>
  )
}

const NavBtn = ({ way, label }) => {
  const go = useCallback(() => store.atoms.routes.current(way), [])
  return (
    <a className="button button-clear" onClick={go}>
      {label}
    </a>
  )
}

export function App() {
  return <FunApp />
}

const capitalize = (s) => {
  if (typeof s !== 'string') return ''
  return s.charAt(0).toUpperCase() + s.slice(1)
}
let lastTitle = ''
const hideClass = 'title hide'
const taClass = 'title tracking-in-08'
const titleAtom = A()
start.up(() => {
  store.atoms.routes.isToken.up((v) => {
    if (v) store.atoms.token.selected.once((v) => titleAtom(v.name))
    else store.atoms.routes.current.once(titleAtom)
  })
})

function Title() {
  const name = useAtom(titleAtom)
  const [titleClass, setTC] = useState(taClass)
  const title = capitalize(name)
  if (lastTitle != title) {
    lastTitle = title
    setTC(hideClass)
    setTimeout(() => setTC(taClass), 12)
  }
  return <h1 className={titleClass}>{title}</h1>
}

function FunApp() {
  const route = useAtom(store.atoms.routes.current)
  const isToken = useAtom(store.atoms.routes.isToken)
  if (isToken || routes[route]) {
    return (
      <div className="route">
        <Title />
        <NavBar route={route} />
        {isToken ? <Token /> : routes[route]}
      </div>
    )
  }
  const ro = routesMap[route]
  if (ro) {
    return <SsrPage route={ro} />
  }
}

function SsrPage({ route }) {
  const data = useAtom(store.atoms.routes[route.state])
  return route.component(data)
}
