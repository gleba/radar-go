import A, {IAtom} from "alak";
import { lo } from '~/front/common/log'

export function atomicWebSocket(url: string) {
  let delay = 100
  let trial = 0

  const open = A.stateless().setId('ws.open') as IAtom<Event>
  const message = A.stateless().setId('ws.message') as IAtom<MessageEvent>
  const close = A.stateless().setId('ws.close') as IAtom<CloseEvent>
  const send = A.stateless().setId('ws.send') as IAtom<string | ArrayBufferLike | Blob | ArrayBufferView>

  function reconnect(e) {
    console.log(`ws: retry in ${delay}ms`);
    delay = delay+(delay/2)
    if (delay>60000)
      delay = 60000
    setTimeout(function () {
      send.clearValue()
      subscribe(new WebSocket(url))
    }, delay);
  }


  function subscribe(ws:WebSocket) {
    send.up(v=>{
      lo.ws('↑', v)
      ws.send(v)
    })
    ws.onopen = open
    ws.onmessage = message
    ws.onclose =  (e) => {
      console.log("ws error", e.code)
      close(e)
      trial++
      switch (e.code) {
        case 1000:	// CLOSE_NORMAL
          console.log("WebSocket: closed");
          break;
        default:	// Abnormal closure
          reconnect(e);
          break;
      }
    }
    ws.onerror = (e:any) => {
      console.log("ws error", e.code)
      close(e)
      // switch (e.code) {
      //   case 'ECONNREFUSED':
      //     reconnect(e);
      //     break;
      //   default:
      //     ws.onerror(e);
      //     break;
      // }
    }
  }
  subscribe(new WebSocket(url))
  return {
    open, message, close, send
  }
}
