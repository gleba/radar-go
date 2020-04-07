export function ABox() {
  const values = new Map()
  return {
    each(key, iterator) {
      let ar = values[key]
      if (ar) {
        ar.forEach(iterator)
      }
    },
    push(key, value) {
      let ar = values[key]
      if (ar) {
        ar.push(value)
      } else {
        values[key] = [value]
      }
    },
    removeAll(key) {
      delete values[key]
    },
    has(key) {
      return !!values[key]
    },
    get(key) {
      return values[key]
    },
    remove(key, value) {
      let ar = values[key]
      if (ar && ar.length) {
        ar.splice(ar.indexOf(value), 1)
      }
      if (!ar.length) {
        delete values[key]
      }
    }
  }
}
