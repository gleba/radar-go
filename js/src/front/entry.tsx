import A from 'alak'

export const start = A()
export const wake = A.stateless()
import { store } from '@store/frontStore'
import { h, render, Component } from 'preact'
import { HeadUpDisplay } from './hud/HeadUpDisplay'
import './styles/root.scss'
import { ComputeStrategy } from 'alak/ext-computed'

render(<HeadUpDisplay />, document.getElementById('hud'))

switch (location.pathname) {
  case '/auth':
    store.actions.session.action('hash', location.hash)
}

store.atoms.routes.current.is(null)
// if (process.env.NODE_ENV == 'development')

// declare type Route = ''
declare module 'alak' {
  interface IAtom<T> {
    from<A extends IAtom<any>[]>(...a: A): ComputeStrategy<T, A>
    match<A extends IAtom<any>[]>(...a: any): any
  }
}
