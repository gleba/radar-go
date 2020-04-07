import { La, qubit, stored } from 'lasens'
import { LaStore } from '@store/frontStore'
import Fuse from 'fuse.js'
import { api } from '@store/api/http'
import { wake } from '~/front/entry'

let options = {
  threshold: 0.3,
  keys: ['symbol', 'name'],
}

export class Discover {
  @qubit ids: KV<TokenIds>
  @stored search
  results = []
  @qubit targetSize

  actions({ a, atoms }: La<Discover, LaStore>) {
    let fuse: Fuse<any, any>
    wake('discover')
    api.get('ids').then((v) => {
      const ar = Object.values(v)
      fuse = new Fuse(ar, options)
      a.ids(v)
      a.targetSize(ar.length)
      a.search.up((v) => {
        a.results(fuse.search(v).map((i) => i.item))
      })
    })
    return {}
  }
}
