///<reference path="../../typings/index.d.ts"/>
import { ISens, LaSens } from 'lasens'
import { Account } from './modules/Account'
//
//
import Routes from './modules/Routes'
//
import { FrontSession } from '@store/modules/FrontSession'
import { IncomeActions } from '@store/modules/IncomeActions'
import { Hud } from '@store/modules/Hud'
import A, { IAtom } from 'alak'
import { Token } from '@store/modules/Token'
import { start } from '~/front/entry'
import { Discover } from '@store/modules/Discover'
import { useAtom } from '@store/preact'
import { Scan } from '@store/modules/Scan'
import { useEffect, useState } from 'preact/hooks'

//
const modules = {
  token: Token,
  scan: Scan,
  account: Account,
  session: FrontSession,
  income: IncomeActions,
  routes: Routes,
  hud: Hud,
  discover: Discover,
}

export type LaStore = ISens<typeof modules>
const frontStore = LaSens(modules).renew()
export const store = frontStore.newContext('root')

export const xAtom = preactAtomsProxyHook(
  store.atoms
) as typeof frontStore.state

interface Mixer {
  (f: (a: typeof store.atoms) => IAtom<any>[], x: any): any
}

export const mix: Mixer = (a, f) => {
  const [state, setState] = useState(null)
  useEffect(() => {
    const atom = A.from(...a(store.atoms)).weak(f)
    atom.up(setState)
  }, [])
  return state
}

start.resend()

function preactAtomsProxyHook(target) {
  const activeModules = {}
  const activeHandlers = {
    get(a, way) {
      return useAtom(a[way])
    },
  }
  return new Proxy(target, {
    get(m, key) {
      let am = activeModules[key]
      if (!am) am = activeModules[key] = new Proxy(m[key], activeHandlers)
      return am
    },
  })
}
