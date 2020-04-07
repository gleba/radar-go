import { La, qubit, stored } from 'lasens'
import { LaStore } from '@store/frontStore'
import { api } from '@store/api/http'

export class Token {
  @qubit selected: TokenTail
  @qubit coefficients = [1]
  @qubit LV
  @qubit volFormat
  @qubit priceFormat
  actions({ a, atoms, actions }: La<Token, LaStore>) {
    return {
      select(t: CoinInfo) {
        atoms.routes.current('-/' + t.slug)
      },
      load(slug: string) {
        api.get('slug/' + slug).then(a.selected)
      },
    }
  }
}
