import { h } from 'preact'
import { useAtom } from '@store/preact'
import { store } from '@store/frontStore'
import { AuthPopUp } from '~/front/hud/AuthPopUp'
import { AddMarketsPopUp } from '~/front/app/tune/AddMarketsPopUp'

export function PopUps() {
  const popup = useAtom(store.atoms.hud.popup)
  switch (popup) {
    case 'markets':
      return <AddMarketsPopUp />
    case 'auth':
      return <AuthPopUp />
  }
  return null
}
