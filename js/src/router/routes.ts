import { Landing } from '~/front/pages/Landing'
import { error404 } from '~/front/pages/error404'
import A from 'alak'
import { Auth } from '~/front/pages/Auth'
///<reference path="../typings/index.d.ts"/>
export const routesMap = {
  '/': {
    title: 'Radar',
    component: Landing,
    state: 'land'
  },
  '/auth': {
    title: 'Radar:auth',
    component: Auth
  },
  404: {
    title: 'Radar:none',
    component: error404
  }
} as TKV<RouteName, RouteBox>


export function makeSsrRouteStore() {
  const dataByRoute = {}
  const dataByState = {}
  const routeByState = {}
  const hashByRoute = {}
  const map = {}
  Object.keys(routesMap).forEach(k => {
    map[k] = true
    routeByState[routesMap[k].state] = k
  })
  return {
    has(route) {
      return map[route]
    },
    dataByRoute(route) {
      return dataByRoute[route]
    },
    setSsrHash(route: string, hash) {
      const data = { time: hash }
      dataByState[routesMap[route].state] = data
      dataByRoute[route] = data
    },
    set(state: string, data) {
      const route = routeByState[state]
      dataByState[state] = data
      dataByRoute[route] = data
      return route
    }
  }
}

