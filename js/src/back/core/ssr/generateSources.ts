import * as path from 'path'
import { readdirSync, readFileSync, rmdirSync, statSync, unlinkSync } from 'fs'
import { artDir } from 'scripts/config'
import NumberFormat = Intl.NumberFormat

import { gzip } from 'zlib'
import { makeHeaderGzExtSize } from './headers'
import { SsrRoute } from '~/back/core/ssr/upRoutes'

export async function generateSources() {
  return await readDir(`${artDir}/front`)
}

const supportedRes = {
  js: true
}

async function readDir(directory) {
  directory = path.resolve(directory)
  const files = readdirSync(directory)
  let size = 0
  const promises = []
  files.forEach(file => {
    const ext = file.split('.').pop()
    if (!supportedRes[ext]) {
      return
    }
    const filePath = path.join(directory, file)
    const stat = statSync(filePath)
    const isDir = stat.isDirectory()
    if (isDir) {
      return
    } else {
      promises.push(
        new Promise(done => {
          gzip(readFileSync(filePath), (err, buffer) => {

            done({
              headers: makeHeaderGzExtSize(ext, buffer.length),
              path: '/' + file,
              buffer
            } as SsrRoute)
          })
        })
      )
    }
  })
  return await Promise.all(promises)
}
