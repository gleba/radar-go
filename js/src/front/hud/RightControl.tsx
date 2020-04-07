import { Component, h } from 'preact'
import { res } from '../common/static'

import { useAtom } from '@store/preact'
import { store } from '@store/frontStore'

export function RadarButton() {
  const shake = useAtom(store.atoms.hud.shake)
  return (
    <div class='radar-button rotate-in-center '>
      <img src={res.logo} className={shake ? 'shake-lr' : ''}/>
    </div>
  )
}


const Offline = () => <div class='tracking-in-08'>
  <div class='login offline'>offline</div>
</div>

function Online() {
  const user = useAtom(store.atoms.account.user)
  return <div class='login tracking-in-03'>{user ? user.username : 'auth'}</div>
}


export function RightControl() {
  const online = useAtom(store.atoms.session.online)
  const popup = useAtom(store.atoms.hud.popup)
  return (
    <div className={`right-control${online ? ' ' : ' gray'}`}
         onClick={store.actions.hud.touch}>
      <div className={popup ? 'hide' : 'show'}>
        {online ? <Online/> : <Offline/>}
      </div>
      <RadarButton/>
    </div>
  )
}
