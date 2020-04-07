import { HttpRequest, HttpResponse, WebSocket } from 'uWebSockets.js'
import { Do } from 'lasens/core/dynamique'
import { BackStore } from '../backStore'
import { holistic, qubit } from 'lasens'

const dec = new TextDecoder()

const ua = require('ua-parser-js')

export class SocketCtrl {
  @qubit goodbye
  @qubit route: RouteReq
  @qubit proxy
  @qubit agent

  public actions({ a, q, actions, dynamique, id }: Do<SocketCtrl, BackStore>) {
    //console.log(id, 'socket')
    let socket: WebSocket
    let sid: string
    a.route.up(rq => actions.routes.request(rq, id))
    a.proxy.stateless().holistic()

    return {
      new(ws) {
        socket = ws
      },
      send(string) {
        console.log('send', string)
        socket.send(string)
      },
      open(req: HttpRequest) {
        sid = req.getQuery()
        const agent = ua(req.getHeader('user-agent'))
        a.agent({
          browser: agent.browser.name,
          os: agent.os.name,
          device: agent.device.model
        })
        dynamique.session(sid).actions.add(id)
      },
      close() {
        dynamique.session(sid).actions.rm(id)
        a.goodbye(id)
        dynamique.socket.removeById(id)
      },
      handleMessage(arrayBuffer: ArrayBuffer) {
        const action: FrontAction = JSON.parse(dec.decode(arrayBuffer))
        const [cmd, data] = action
        console.log(cmd, data)
        switch (cmd) {
          case 'route':
            actions.routes.request(data, id)
            break
          default:
            a.proxy(cmd, data)
        }
      }
    }
  }
}
