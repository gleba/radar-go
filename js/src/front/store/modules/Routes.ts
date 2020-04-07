///<reference path="../../../typings/index.d.ts"/>
import { La, qubit, stored } from 'lasens'
import { render } from 'preact'
import { makeSsrRouteStore, routesMap } from '~/router/routes'
import { LaStore } from '@store/frontStore'
import { App } from '~/front/app/App'

export default class Routes {
  // @qubit path
  @stored current: AppRoutes | string
  // @qubit isApp = false
  @qubit isToken
  @qubit land

  actions({ a, q, atoms, actions }: La<Routes, LaStore>) {
    const hashes = {}
    if (window['routeHash']) {
      hashes[location.pathname] = window['routeHash']
    }
    window['ssrRoute'] = a.current
    if (!q.current) {
      a.current(location.pathname)
    }
    let isDrawn = false
    a.current.up(route => {
      console.log("→ ", route)
      const ssr: RouteBox = routesMap[route]
      if (!ssr) {
        const isToken = route[0] == '-'
        if (isToken) {
          window.history.pushState(null, 'radar', "/"+route)
          actions.token.load(route.split('/')[1])
        } else {
          window.document.title = route
          window.history.pushState(null, 'radar', "/")
        }
        a.isToken(isToken)
      } else {
        a.isToken.is(true) && a.isToken(false)
        window.history.replaceState(null, ssr.title, route)
      }
      actions.session.action('route', {
        route,
        state: ssr?.state,
        time: hashes[route]?.time
      })
      if (!isDrawn) {
        render(App(), document.getElementById('root'))
        isDrawn = true
      }
    })
    return {
      updateState({ state, data }: RouteResp) {
        a[state](data)
      }
    }
  }
}
