import { BackStore } from '~/back/core/backStore'
import { Do } from 'lasens/core/dynamique'
import { User } from 'node-telegram-bot-api'
import { newRune } from '~/shared/rune'

export default class PreAuth {
  actions({ a, q, actions }: Do<PreAuth, BackStore>) {
    const runes = {}
    const clearMessage = {}
    return {
      clearLink(chatId, messageId) {
        clearMessage[chatId] = messageId
      },
      preAuth(user: User) {
        const rune = `#` + newRune(42)
        runes[rune] = user
        return `http://localhost:4444/auth` + rune
      },
      hasUser(rune: string) {
        const u = runes[rune]
        if (u){
          const mid = clearMessage[u.id]
          actions.bot.deleteMessage(u.id, mid)
          delete clearMessage[u.id]
          delete runes[rune]
          return u
        }

      }
    }
  }
}

