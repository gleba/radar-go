
var ws = "color: #7B86BA;";
var gray = "color: gray;";
var green = "color: green;";


const { log } = console
export const lo = {
  ws: (...v) => log(`%c${v[0]} %c ${v[1]}`, ws, gray, ...v.slice(2)),
  api: (...v) => {
    console.groupCollapsed(`%c${v[0]} %c ${v[1]}`, green, gray, )
    log(...v.slice(2))
    console.groupEnd()
  }
}
