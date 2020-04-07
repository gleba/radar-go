const types = {
  ico: 'image/x-icon',
  html: 'text/html; charset=utf-8',
  js: 'text/javascript; charset=utf-8',
  json: 'application/json',
  css: 'text/css',
  png: 'image/png',
  jpg: 'image/jpeg',
  wav: 'audio/wav',
  mp3: 'audio/mpeg',
  svg: 'image/svg+xml',
  pdf: 'application/pdf',
  doc: 'application/msword',
}

export function makeHeaderGzExtSize(ext, size) {
  return [
    ['Content-Type', types[ext]],
    ['Content-Encoding', 'gzip'],
    ['Content-Length', size.toString()],
  ]
}
//
export function makeContentHeader(ext: keyof typeof types) {
  return ['Content-Type', types[ext]]
}
// export function makeHeaderSize(ext, size) {
//   return {
//     'Content-Type': types[ext],
//     'Content-Length': size
//   }
// }
//
