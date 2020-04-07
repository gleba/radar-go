import { La, qubit } from 'lasens'
import { BackStore, sens } from '~/back/core/backStore'
import TelegramBot, { Message, SendMessageOptions } from 'node-telegram-bot-api'
import { Do } from 'lasens/core/dynamique'

const markdown: SendMessageOptions = {
  parse_mode: 'Markdown'
}

export default class Telegram {
  // @qubit instance: TelegramBot

  actions({ a, q, actions, dynamique }: Do<Telegram, BackStore>) {
    // const bot = new TelegramBot('708785520:AAFa4ns6-OpdXZzL7QXrh0TMpIpPiBOQl8w', { polling: true })
    // a.instance(bot)
    // bot.getMe().then(bot => {
    //   console.log('telegram as:', bot.username)
    // })
    // bot.on('message', (m: Message) => {
    //   const [cmd, sid] = m.text.split(' ')
    //   const chatId = m.chat.id
    //   console.log(m.chat.username, cmd, sid)
    //   console.log(m.from.id, m.chat.id)
    //
    //   switch (cmd.toLowerCase()) {
    //     case '/start':
    //       if (sid) {
    //         const session = dynamique.session(sid)
    //         session.actions.auth(m.from)
    //         const agent = session.atoms.agent.value
    //         bot.sendMessage(chatId, `Вернитесь к сессии : ${agent.browser}`)
    //       } else {
    //         const url = sens.actions.preAuth.preAuth(m.from)
    //         const body = `Для входа в личный кабинет откройте ссылку в рабочем браузере: [${url}](${url})`
    //         bot.sendMessage(chatId, body, markdown)
    //           .then(l => sens.actions.preAuth.clearLink(chatId, l.message_id))
    //       }
    //   }
    // })

    return {
      deleteMessage(chatId, messageId){
        bot.deleteMessage(chatId, messageId)
      },
      send(id: number, message: string) {
        bot.sendMessage(id, message)
      }
    }
  }
}

