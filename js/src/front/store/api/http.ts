import { lo } from '~/front/common/log'
import { domain } from '~/front/common/static'

const apiUrl = location.hostname == domain ? `/api/` : 'http://localhost:4002/'
const get = path => fetch(apiUrl+path)
  .then((response) =>
    response.json()
      .then(v=>{
        lo.api("↓",path,v)
        return v
      })
  )
  // .then((data) => {
  //   console.log(data);
  // })

export const api = {
  get
}
