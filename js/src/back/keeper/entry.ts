import NATS from 'nats'
import { makeCollection } from '~/back/keeper/db/makeCollection'
import { openMongo } from '~/back/keeper/db/openMongo'
import { App, HttpResponse, SHARED_COMPRESSOR, WebSocket } from 'uWebSockets.js'
import chalk from 'chalk'
import _ from 'lodash'
import { head } from '~/back/keeper/http'
import { createHash } from 'crypto'
import { mean, median } from '~/shared/median'
import { makeKeeper } from '~/back/keeper/keep'
import { updateLand } from '~/back/keeper/updateland'
import { getIds, getSlug, getTune } from '~/back/keeper/rq-tune'
import { getLimits, getVolatile } from '~/back/keeper/rq-max'

export const nc = NATS.connect({
  url: 'nats://localhost:4222',
  token: '2yKnjkfXCtA8ik2yKnjkfXCtA8ik',
})

export const keep = makeKeeper()

export const idBySlug = {}
keep.ready.up(() => {
  nc.subscribe('flow.>', (msg, reply, subject) => {
    const path = subject.split('.')
    const [_, namespace, action, id] = path
    const collection = keep.ns[namespace]
    collection && collection.action(action, msg)
    updateLand()
  })
  nc.subscribe('front', updateLand)
  setTimeout(updateLand, 1000)
})

const app = App()

app
  .get('/slug/:id', (res, req) => {
    head(res)
    const id = idBySlug[req.getParameter(0)]
    if (id) {
      res.end(JSON.stringify(getSlug(id)))
    } else {
      res.end(JSON.stringify('nope'))
    }
  })
  .get('/vv', (res, req) => {
    head(res)
    res.end(JSON.stringify(getVolatile()))
  })
  .get('/limits', (res, req) => {
    head(res)
    res.end(JSON.stringify(getLimits()))
  })
  .get('/markets', (res, req) => {
    head(res)
    const markets = {}
    keep.ns.market.each((i) =>
      i.markets.forEach((m) => (markets[m.id] = m.name))
    )
    res.end(JSON.stringify(markets))
  })
  .get('/tune', (res, req) => {
    head(res)
    res.end(JSON.stringify(getTune()))
  })
  .get('/ids', (res, req) => {
    head(res)
    res.end(JSON.stringify(getIds()))
  })
  .get('/alive', (res, req) => {
    const id = req.getQuery()
    head(res)
    res.end(
      JSON.stringify({
        k: keep.ns.daily.ram[id],
        now: keep.ns.coin.ram[id],
      })
    )
  })
  .any('/*', (res, req) => {
    const path = req.getUrl()
    console.log(path)
    res.writeStatus('404')
    res.end(JSON.stringify('•'))
  })
  .listen(4002, (fine) => console.log(chalk.yellow('• keeper started on 4002')))
