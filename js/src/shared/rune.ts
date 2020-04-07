export function newRune(length) {
  let charset =
      'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789',
    i,
    result = ''
  if ((process as any).browser) {
    if (window.crypto && window.crypto.getRandomValues) {
      let values = new Uint32Array(length)
      window.crypto.getRandomValues(values)
      for (i = 0; i < length; i++) {
        result += charset[values[i] % charset.length]
      }
      return result
    }
    for (i = 0; i < length; i++) {
      result += charset[Math.floor(Math.random() * charset.length)]
    }
    return result
  } else {
    return Buffer.alloc(length, require('crypto').randomBytes(length)).toString('hex')
  }
}
