import { makeContentHeader } from '~/back/core/ssr/headers'

export const head = (res) => {
  res.writeHeader('Access-Control-Allow-Origin', '*')
  res.writeHeader(
    'Access-Control-Allow-Headers',
    'Origin, X-Requested-With, Content-Type, Accept'
  )
  res.writeHeader(
    'Access-Control-Allow-Headers',
    'Origin, X-Requested-With, Content-Type, Accept'
  )
  res.writeHeader(...makeContentHeader('json'))
  res.writeStatus('200')
}
