// import { commands } from '~/back/core/tg/commands'
//
// (process as any).env["NTBA_FIX_319"] = 1 as any
//
//
// const TelegramBot = require('node-telegram-bot-api')
// const token = '708785520:AAFa4ns6-OpdXZzL7QXrh0TMpIpPiBOQl8w'
// export const bot = new TelegramBot(token, {polling: true})
//
//
// bot.getMe().then(bot => {
//   console.log("telegram bot:", bot.username)
// })
//
//
// bot.onText(/\/echo (.+)/, (msg, match) => {
//   const chatId = msg.chat.id
//   const resp = match[1] // the captured "whatever"
//   bot.sendMessage(chatId, resp)
// })
//
// bot.on('message', (msg) => {
//   console.log('telegram bot message:', msg.text)
//   const chatId = msg.chat.id
//   const fn = commands[msg.text.toLowerCase().split(" ")[0]]
//   if (fn) {
//     fn(msg)
//   } else {
//     // bot.sendMessage(msg.chat.id, 'Доступна команда "войти"');
//   }
// })
//
