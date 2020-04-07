import { Worker } from 'worker_threads'
import { devArtEntry, thoughts, artDir } from 'scripts/config'

import { ChildProcess, fork } from 'child_process'


const artifacts = {}
const artifactEntry = process.env.isDev ? devArtEntry : "app"
thoughts.forEach(name => (artifacts[name] = `./${artDir}/${name}/${artifactEntry}.js`))


const activeWorkers = new Map<string, ChildProcess>()
// //
async function reRunThread(name) {
  if (activeWorkers.has(name)) {
    await activeWorkers.get(name).kill()
  }
  let process = fork(artifacts[name])
  activeWorkers.set(name, process)
  console.log("reRunThread:", name)
}

thoughts.forEach(reRunThread)

if (process.env.isDev) {
  process.on('message', reRunThread)
}
