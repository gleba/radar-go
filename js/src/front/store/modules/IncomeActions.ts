import {La} from 'lasens'
import { LaStore } from '@store/frontStore'
import { lo } from '~/front/common/log'

export class IncomeActions {
  actions({a, actions, atoms}: La<IncomeActions, LaStore>) {
    return {
      receive(action:BackAction){
        const [command, data] = action
        lo.ws("↓", command, data)
        switch (command) {
          case "route":
            actions.routes.updateState(data)
            break
          case 'wrong-hash':
            atoms.routes.current(command)
            break
          case 'auth':
            atoms.account.user(data)
            break
        }
      }
    }
  }
}
