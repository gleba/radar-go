import { App, HttpResponse, SHARED_COMPRESSOR, WebSocket } from 'uWebSockets.js'

import { ramRoutes, writeRoute } from './ramRoutes'
import chalk = require('chalk')
import { newRune } from '~/shared/rune'
import { ctrl } from '~/back/core/backStore'

const app = App()
// const pool = new Map<WebSocket, any>()
app
  .ws('/ws', {
    // compression: SHARED_COMPRESSOR,
    maxPayloadLength: 16 * 1024 * 1024,
    idleTimeout: 360,
    open(ws, req) {
      ws.id = newRune(6)
      ctrl.socket(ws).actions.open(req)
    },
    message(ws, message, isBinary) {
      ctrl.socket(ws).actions.handleMessage(message)
    },
    close(ws) {
      ctrl.socket(ws).actions.close()
    },
  })
  .any('/*', (res, req) => {
    const path = req.getUrl()
    const route = ramRoutes[path]
    if (route) {
      writeRoute(route, res, '200')
    } else {
      res.writeStatus('404')
      res.end('404 bro')
      // writeRoute(ramRoutes["/"], res, '200')
    }
  })
  .listen(4001, (fine) => console.log(chalk.yellow('• uws started on 4001')))
