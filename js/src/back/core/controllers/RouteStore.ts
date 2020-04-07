import { BackStore, ctrl } from '~/back/core/backStore'
import { Do } from 'lasens/core/dynamique'
import { User } from 'node-telegram-bot-api'

export default class RouteStore {
  land = {
    daily: 0,
    pulse: 0,
    time: Date.now()
  }

  actions({ a, q, dynamique }: Do<RouteStore, BackStore>) {
    const listeners = {}
    const land = 'land'

    function packAction(state, data) {
      const action: BackAction = ['route', { state, data }]
      return JSON.stringify(action)
    }

    a.land.up(data => {
      const json = packAction(land, data)
      Object.keys(listeners).forEach(id => {
        if (listeners[id] == 'land')
          dynamique.socket(id).actions.send(json)
      })
    })
    const onSessionEnd = sid => {
      delete listeners[sid]
    }
    return {
      request({ state, time }: RouteReq, id) {
        const sock = dynamique.socket.getById(id)
        if (!sock) return
        if (!state) {
          if (listeners[id]) {
            console.log("removed", listeners[id], id)
            delete listeners[id]
          }
          return
        }
        console.log("add", state, id)
        listeners[id] = state
        let current = q[state]
        if (current?.time != time) {
          sock.actions.send(packAction(state, current))
        }
        sock.atoms.goodbye.up(onSessionEnd)
      }
    }
  }
}

