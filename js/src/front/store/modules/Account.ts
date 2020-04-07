import { La, qubit } from 'lasens'
import { LaStore } from '@store/frontStore'


export class Account {
  @qubit user

  actions({ a, atoms }: La<Account, LaStore>) {
    return {
      logout() {
      }
    }
  }
}
