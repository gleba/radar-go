import { ctrl } from '~/back/core/backStore'


export function newOutPool(timeout=120000) {
  const pool = []
  let clearStamp
  return {
    add(f) {
      clearInterval(clearStamp)
      pool.push(f)
    },
    // send(raw:string) {
    //   const json = JSON.stringify(raw)
    //   pool.forEach(f => f(json))
    // },
    send(raw:string) {
      pool.forEach(f => f(raw))
    },
    action(a:BackCommand, ...v:any) {
      const json = JSON.stringify([a, ...v])
      pool.forEach(f => f(json))
    },
    rm(f, onRemoveLast) {
      pool.splice(pool.indexOf(f), 1)
      if (!pool.length) {
        clearStamp = setTimeout(onRemoveLast, timeout)
      }
    }
  }
}
