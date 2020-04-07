import { La, qubit } from 'lasens'
import { BackStore, ctrl } from '~/back/core/backStore'
import { newOutPool } from '~/shared/io.channels'
import { newLevelStore } from '~/back/core/db/levelStore'
import { Do } from 'lasens/core/dynamique'
import A from 'alak'
import { User } from 'node-telegram-bot-api'

const db = newLevelStore('profiles')


export default class UserCtrl {

  @qubit profile: User

  actions({ a, q, actions, id, dynamique }: Do<UserCtrl, BackStore>) {
    console.log(id, 'user')
    const sessions = newOutPool()
    db.restore(id, a.profile)

    function auth(send) {
      if (q.profile) {
        send(JSON.stringify(['auth', q.profile]))
      }
    }

    function userOut() {
      dynamique.user.removeById(id)
    }

    return {
      handleWeb(cmd:FrontCommand, data) {
        switch (cmd) {
        }
      },
      auth,
      add(sendFn) {
        console.log(id, 'user add session')
        sessions.add(sendFn)
      },
      rm(sendFn) {
        sessions.rm(sendFn, userOut)
      },
      init(user: User) {
        delete user.is_bot
        delete user.id
        db.put(id, user)
        a.profile(user)
        auth(sessions.send)
      }
    }
  }
}

