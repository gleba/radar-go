import { generateSources } from '~/back/core/ssr/generateSources'
import A from 'alak'
import { generateJsFooter, generatePages, renderPage } from '~/back/core/ssr/generatePages'
import { IAtom } from 'alak/atom'
import { sens } from '../backStore'
import { routesMap } from '~/router/routes'

export type SsrRoute = {
  path: string
  headers: string[][]
  buffer: Buffer
}

export function upRoutes(): IAtom<SsrRoute> {
  const upRoute = A()
  generateSources().then(async js => {
    js.forEach(upRoute)
    generateJsFooter(js)
    const pages = await generatePages()
    pages.forEach(upRoute)
    Object.keys(routesMap).forEach(path => {
      const ro = routesMap[path]
      if (ro.state) {
        sens.atoms.routes[ro.state].next(() =>
          renderPage(path).then(upRoute)
        )
      }
    })
  })
  return upRoute
}
