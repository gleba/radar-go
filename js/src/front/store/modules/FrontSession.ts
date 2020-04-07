import { La, stored } from 'lasens'
import { atomicWebSocket } from '../api/atomicWebSocket'
import { newRune } from '~/shared/rune'
import { LaStore } from '@store/frontStore'
import { domain } from '~/front/common/static'

const host = location.hostname == domain ? `wss://${domain}/ws?` : 'ws://localhost:4001/ws?'

export class FrontSession {
  @stored auth = false
  @stored rune: string
  online = false

  actions({ a, actions, q }: La<FrontSession, LaStore>) {

    const queue = []
    if (a.rune.isEmpty) {
      a.rune(newRune(24))
    }
    const aws = atomicWebSocket(host + q.rune)
    aws.close.up(() => a.online(false))
    aws.open.up(() => {
      a.online(true)
      while (queue.length) {
        aws.send(queue.pop())
      }
    })
    aws.message.up(v => {
        actions.income.receive(JSON.parse(v.data))
      }
    )

    function send(json: string) {
      if (q.online) {
        aws.send(json)
      } else {
        queue.push(json)
      }
    }

    return {
      action(cmd: FrontCommand, data?) {
        send(JSON.stringify([cmd, data]))
      }
    }
  }
}
