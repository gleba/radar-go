import { La, qubit } from 'lasens'
import { LaStore } from '@store/frontStore'
import { setCssVar } from '~/front/common/static'

type PopUps = 'markets' | 'auth' | null

export class Hud {
  @qubit popup: PopUps
  @qubit authUrl: string
  @qubit shake = false
  hideText = false

  actions({ a, q, atoms, actions }: La<Hud, LaStore>) {
    a.popup.upSome(() => {
      setCssVar('--route-filter-time', '987ms')
      setCssVar('--route-filter', 'blur(5px)')
    })
    a.popup.upNone(() => {
      setCssVar('--route-filter-time', '233ms')
      setCssVar('--route-filter', 'none')
    })
    const shakeOff = () => a.shake(false)
    a.shake.upTrue(() => setTimeout(shakeOff, 1000))
    atoms.account.user.up((v) => a.popup(null))
    atoms.session.rune.up((v) =>
      a.authUrl(
        'tg://resolve?domain=radar_cashbot&start=' + atoms.session.rune.value
      )
    )

    function auth() {
      a.shake(true)
      a.popup('auth')
      window.open(q.authUrl, '_self')
    }

    return {
      auth,
      touch() {
        document.body.requestFullscreen({
          navigationUI: 'hide',
        })
        if (atoms.account.user.isEmpty) {
          auth()
        } else {
          a.shake(true)
          atoms.routes.current('tune')
        }
      },
      clarity() {
        a.popup(null)
      },
    }
  }
}
