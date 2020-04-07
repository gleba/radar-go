import { createHash } from 'crypto'
import { mean } from '~/shared/median'
import { idBySlug, keep, nc } from '~/back/keeper/entry'

let times = []
let lastHash = ''
let lastTime = Date.now()

export function updateLand() {
  let land = {} as any
  const earth = {
    hour: 0,
    day: 0,
    week: 0,
  } as any
  Object.keys(keep.ns.coin.ram).forEach((id) => {
    const coin = keep.ns.coin.ram[id]
    idBySlug[coin.slug] = coin.id
    if (coin.quote.USD.percent_change_1h < 0) earth.hour++
    if (coin.quote.USD.percent_change_24h < 0) earth.day++
    if (coin.quote.USD.percent_change_7d < 0) earth.week++
  })
  earth.hourUp = land.earth
    ? earth.hour > land.earth.hour
      ? true
      : earth.hour < land.earth.hour
      ? false
      : null
    : null
  earth.dayUp = land.earth
    ? earth.day > land.earth.day
      ? true
      : earth.day < land.earth.day
      ? false
      : null
    : null
  earth.weekUp = land.earth
    ? earth.week > land.earth.week
      ? true
      : earth.week < land.earth.week
      ? false
      : null
    : null

  land = {
    earth,
    all: keep.ns.pulse.size(),
    alive: keep.ns.daily.size(),
  }
  const hash = createHash('md5').update(JSON.stringify(land)).digest('hex')
  if (lastHash != hash) {
    let time = Date.now()
    lastHash = hash
    const changes = Math.round((time - lastTime) / 1000)
    if (changes > 1) times.push(changes)
    if (times.length > 100) {
      times.pop()
    }
    console.log(times, changes)
    lastTime = time
    nc.publish(
      'land',
      JSON.stringify(
        Object.assign(land, {
          hash,
          time,
          step: {
            max: Math.max(...times),
            min: Math.min(...times),
            mean: mean(times),
          },
        })
      )
    )
  }
}
