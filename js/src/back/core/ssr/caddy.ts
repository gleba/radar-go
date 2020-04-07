const gzip = {
  handle: [{
    encodings: { gzip: {} },
    handler: 'encode'
  }]
}
const jsHead = {
  handle: [{
    handler: 'headers',
    response: {
      set:
        {
          'Content-Type': ['text/javascript']
        }
    }
  }]
}

function makeRoute(paths, body) {
  return {
    handle: [{
      body,
      handler: 'static_response',
      status_code: 200
    }],
    match: [{ path: paths }]
  }
}

import { request } from 'http'

export function upCaddy(files) {

  const rq = request({
    host: 'localhost',
    port: 2019,
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    path: '/config/apps/http/servers/'
  })
  rq.write(JSON.stringify({
    listen: [':2020/js'],
    routes: [gzip, jsHead, ...files.map(f => {
      console.log(f[0])
      return makeRoute(['/' + f[0]], f[2].toString())
    })]
  }))
  rq.end()
  console.log('caddy ok')
}
