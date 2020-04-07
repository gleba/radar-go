import { routesMap } from '~/router/routes'

import render from 'preact-render-to-string'
import { indexHtml } from '~/back/core/ssr/indexHtml'
import { SsrRoute } from '~/back/core/ssr/upRoutes'
import { gzip } from 'zlib'
import { makeHeaderGzExtSize } from './headers'
import { sens } from '../backStore'

const head = title => `<title>${title}</title>`
const source = f => `<script type="text/javascript" src="${f.path}"></script>`
let jsFooter = ''

export function generateJsFooter(js) {
  jsFooter = js.map(source)
}

export async function generatePages() {
  const promises = Object.keys(routesMap).map(renderPage)
  return await Promise.all(promises)
}

export async function renderPage(path) {
  return new Promise(done => {
    const ro = routesMap[path]
    const data = ro.state ? sens.atoms.routes[ro.state]?.value : {}
    const page = indexHtml(head(ro.title), render(ro.component(data)), jsFooter, data?.hash)
    gzip(page,
      (err, buffer) => {
        if (err) {
          console.log('gzip error:', err)
        }
        done({
          headers: makeHeaderGzExtSize('html', buffer.length),
          path,
          buffer
        } as SsrRoute)
      })
  })
}
