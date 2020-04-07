import { La, qubit } from 'lasens'
import { BackStore } from '~/back/core/backStore'
import { newOutPool } from '~/shared/io.channels'
import { Do } from 'lasens/core/dynamique'
import { User } from 'node-telegram-bot-api'
import { newLevelStore } from '~/back/core/db/levelStore'

const db = newLevelStore('session.uid')

export default class SessionCtrl {
  @qubit send
  @qubit uid

  @qubit webProxy
  @qubit agent
  actions({ a, q, actions, dynamique, id }: Do<SessionCtrl, BackStore>) {
    console.log(id, 'session')
    const sockets = newOutPool()
    db.restore(id, a.uid)

    function sessionClose() {
      q.uid && dynamique.user(q.uid).actions.rm(sockets.send)
      dynamique.session.removeById(id)
    }

    a.webProxy.stateless().holistic()
    a.webProxy.up((cmd: FrontCommand, data) => {
      switch (cmd) {
        case 'hash':
          const u:User = actions.preAuth.hasUser(data)
          if (u) {
            actions.bot.send(u.id, `Успешная авторизация: ${q.agent.browser} ${q.agent.os} ${q.agent.device && q.agent.device}`)
            auth(u)
          } else {
            sockets.action("wrong-hash")
          }
      }
    })
    function auth(user: User) {
      const uid = user.id
      dynamique.user(uid).actions.init(user)
      a.uid(uid)
      db.put(id, uid)
    }
    return {
      add(xid) {
        const sock = dynamique.socket(xid)
        sock.atoms.proxy.up(a.webProxy)
        sock.atoms.agent.up(a.agent)
        sockets.add(sock.actions.send)
        q.uid && dynamique.user(q.uid).actions.auth(sock.actions.send)
      },
      rm(xid) {
        const sock = dynamique.socket(xid)
        sockets.rm(sock.actions.send, sessionClose)
      },
      auth
    }
  }
}
