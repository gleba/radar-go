import * as path from 'path'
import { existsSync, mkdirSync, readdir, readdirSync, rmdirSync, statSync, unlink, unlinkSync } from 'fs'
import { start } from 'repl'
import * as os from 'os'


export function mkDir(directory) {
  let fullPath = path.resolve(directory)
  let isExists = existsSync(fullPath)
  if (!isExists) {
    mkdirSync(fullPath)
  }
  return isExists
}

export function clearDir(directory) {
  // console.log("clear::", directory)
  let fullPath = path.resolve(directory)
  if (existsSync(fullPath)) {
    const files = readdirSync(fullPath)
    files.forEach(file => {
      const filePath = path.join(fullPath, file)
      if (statSync(filePath).isDirectory()) {
        clearDir(filePath)
        rmdirSync(filePath)
      } else {
        unlinkSync(filePath)
      }
    })
  }
}

export function prepareDir(directory, keepPrev?) {
  let fullPath = path.resolve(directory)
  // console.log("prepareDir", directory)
  if (!keepPrev &&existsSync(fullPath)) {
    // console.log("clear →", directory)
    clearDir(directory)
  } else {
    const parentDir = fullPath.split(path.sep)
    parentDir.pop()
    const parent = parentDir.join(path.sep)
    if (existsSync(parent)) {
      // console.log("have a parent", parent)
      mkdirSync(fullPath)
      // console.log(fullPath, existsSync(fullPath))
      // setTimeout(()=>console.log(fullPath, existsSync(fullPath)), 1000)
    } else {
      // console.log("↓ keep", parent)
      prepareDir(parent, true)
      mkdirSync(fullPath)
    }
  }
}

