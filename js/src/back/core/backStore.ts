
import {Dynamique, IDynamique, LaSens} from "lasens";


import RouteStore from "src/back/core/controllers/RouteStore";

import Telegram from '~/back/core/controllers/Telegram'

import UserCtrl from '~/back/core/controllers/UserCtrl'
import { SocketCtrl } from '~/back/core/controllers/SocketCtrl'
import SessionCtrl from '~/back/core/controllers/SessionCtrl'
import PreAuth from '~/back/core/controllers/PreAuth'


const modules = {
  routes: RouteStore,
  bot:Telegram,
  preAuth:PreAuth
}
const controllers = {
  socket:SocketCtrl,
  user:UserCtrl,
  session:SessionCtrl
}

export type BackStore = IDynamique<typeof modules, typeof controllers>
const backStore = Dynamique(LaSens(modules), controllers)
backStore.renew()
export const ctrl = backStore.dynamique
export const sens = {
  atoms: backStore.atoms,
  state: backStore.state,
  actions: backStore.actions
}

