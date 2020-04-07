import { fusebox } from 'fuse-box/core/fusebox'
import { thoughts, artDir, arts } from './config'
import { clearDir, prepareDir } from './tool'
import * as path from 'path'
import chalk from "chalk";
import {createPrivateKey} from "crypto";
const bundleLog = (a) => console.log('🗸', chalk.green(a.bundles.map(b => b.absPath.toString())))
export async function build(isDev) {
  let distRoot
  clearDir(artDir)
  clearDir('./.cache')
  const uglify = { uglify: true }
  const backRoot = `${artDir}/back/`
  prepareDir(backRoot)
  let app = 'index.js'
  for (const name of thoughts) {
    const box = back(isDev, name)
    distRoot = path.join(backRoot, name)
    const bundles = {
      distRoot,
      app
    }
    isDev ?
      box.runDev({ bundles }).then(
        r => r.onComplete(h => {
          bundleLog(r)
          console.log(chalk.green("🗸  start "+ name))
          h.server.start()
        })
      )
      :
      box.runProd({
        bundles,
        uglify
      }).then(bundleLog)

  }
  distRoot = `${artDir}/front/`
  prepareDir(distRoot)
  for (const a of arts) {
    const [name] = a
    const box = front(isDev, a[0], a[1])
    const bundles = {
      distRoot,
      app: `${name}-.$hash.js`
    }
    isDev ?
      box.runDev({ bundles }).then(bundleLog)
      :
      box.runProd({ bundles, uglify }).then(bundleLog)
  }
}

//
function front(isDev, name, useIndex) {
  return fusebox({
    target: 'browser',
    // logging: { level: 'disabled' },
    entry: `src/front/entry.tsx`,
    watcher: {
      root: path.resolve("src/front")
      // }
    },
    env:{
      API: JSON.stringify("localhost:4001")
    },
    cache: true,
    sourceMap: isDev,
    devServer: {
      enabled: useIndex,
      hmrServer: {
        useCurrentURL:false,
        port:4003,
        connectionURL: "wss:/radar.cash/hmr",
        enabled:useIndex,
        // connectionURL: "ws://localhost:4444"
      },
    },
    webIndex: useIndex && {
      template: 'src/front/index.html'
    }
  })
}

//
function back(isDev, name) {
  return fusebox({
    target: 'server',
    // logging: { level: 'disabled' },
    entry: `src/back/${name}/entry.ts`,
    watcher: isDev,
    // watcher: {
    //   // include:[`src/back/${name}/*`],
    //   // @ts-ignore
    //   root: path.resolve("packages")
    //   // include:[`src/back//*`],
    //   // chokidarOptions:{
    //   //
    //   // }
    // },
    cache: false,
    sourceMap: isDev,
    dependencies: {
      // include: ['alak'],
      // ignoreAllExternal: isDev,
      ignore: ['level', 'uWebSockets.js']
    }
  })
}
