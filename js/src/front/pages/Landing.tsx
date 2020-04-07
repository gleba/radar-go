import { h, Fragment } from 'preact'
import { useState } from 'preact/hooks'
import { ssrNav } from '~/front/common/ssrNav'

const variations = {
  true: <span className="upcolor">uptrend</span>,
  false: <span className="downcolor">downtrend</span>,
  null: <span>stable</span>,
}

function Pcent({ pct, label, up }) {
  return (
    <div>
      <span class="upcolor">{pct.toFixed(2)}% </span>
      grown up per {label}, now is {variations[up]}
    </div>
  )
}

const sec = (time) => (Math.round((Date.now() - time) / 100) / 10).toFixed(1)
let t

function Watch({ time }) {
  const [dist, setDist] = useState(sec(time))
  if (process['browser']) {
    clearInterval(t)
    t = setInterval(() => setDist(sec(time)), 144)
    return (
      <div className="column tracking-in-08">
        last update at {dist} seconds ago{' '}
      </div>
    )
  } else {
    return <div style="visibility:hidden"> seconds ago </div>
  }
}

export function Landing(props) {
  if (!props?.alive) return <div>☉ loading... </div>
  const alive = (props.alive / props.all) * 100
  const g1 = (props.earth.hour / props.all) * 100
  const g2 = (props.earth.day / props.all) * 100
  const g3 = (props.earth.week / props.all) * 100
  const { hourUp, hour, dayUp, day, weekUp, week } = props.earth

  return (
    <div className="route">
      <div className="full-size">
        <Pie key={g3} value={g3} s={3} />
        <Pie key={g2} value={g2} s={2} />
        <Pie key={g1} value={g1} s={1} />
      </div>
      <h1
        className="tracking-in-08"
        style="padding-left:24px; padding-top:12px"
      >
        Radar{' '}
      </h1>
      <div className="container home">
        <h2 className="tracking-in-03">Market Stats</h2>
        <div class="column tracking-in-03">
          <div>Assets : {props.all} </div>
          <Pcent up={hourUp} pct={g1} key={hour} label={'hour'} />
          <Pcent up={dayUp} pct={g2} key={day} label={'day'} />
          <Pcent up={weekUp} pct={g3} key={week} label={'week'} />
          <div>
            Targeted: <b>{props.alive}</b>
          </div>
          <div class="updatezone">
            <Watch time={props.time} />
            {props.step.mean && (
              <div>
                hour dynamic ~{props.step.mean.toFixed(0)}({props.step.min}-
                {props.step.max}) sec
              </div>
            )}
          </div>
        </div>
        <div className="discover">
          <a
            className="button button"
            href="/discover"
            onClick={ssrNav('discover')}
          >
            discover
          </a>
        </div>
      </div>
      <div class="footer">★ всё клёво ★</div>
    </div>
  )
}

const P = 2 * Math.PI * 180

function Pie({ value, s }) {
  const v = 100 - value
  const pv = (P / 100) * v
  const r = (360 / 100) * (v / 2) + 270
  return (
    <div class="pie">
      <figure style={`width:${s * 100}px; height:${s * 100}px; opacity: 0.9;`}>
        <svg
          width="100%"
          height="100%"
          viewBox="0,0,720, 720"
          style={`transform: rotate(-${r}deg);`}
        >
          <circle
            r="180"
            cx="360"
            cy="360"
            className="pie"
            style={`
                stroke-width: 360px;
                stroke-dasharray: ${pv} ${P - pv}; stroke: var(--color);
                transition: stroke-dasharray .3s ease;
                `}
          />
        </svg>
      </figure>
    </div>
  )
}
