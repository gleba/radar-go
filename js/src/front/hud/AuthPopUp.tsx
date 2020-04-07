import { h } from 'preact'
import { store } from '@store/frontStore'
import { useAtom } from '@store/preact'

export function AuthPopUp() {
  const authUtl = useAtom(store.atoms.hud.authUrl)
  console.log(authUtl)

  const block = (e: MouseEvent) => {
    e.preventDefault()
    e.stopImmediatePropagation()
    e.stopPropagation()
  }
  return (
    <div class='popup-zone'>
      <div class='popup-close' onClick={store.actions.hud.clarity}/>
      <div class='popup slide-in-elliptic-top-fwd' onClick={block}>
        <div class='title'>Авторизация</div>
        <div class='info'>Для следующего шага необходим установленный телеграм на этом устройстве.
          <p>После открытия <a onClick={store.actions.hud.auth} target='_self'>ссылки</a>, должен открытся телеграм чат
            с ботом @radar_cashbot, где необходимо нажать на старт и
            вернутся сюда.
          </p>
          <p>Если чат не открылся или открылся без кнопки старт, необходимо самостоятельно найти
            бота <b>@radar_cashbot</b> и отправить
            команду <b>/start</b>
          </p>
        </div>
        <div class='icon' onClick={store.actions.hud.auth}>
          <svg width="24px" height="24px" version="1.1"
               style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;">
            <path id="telegram-1"
                  d="M18.384,22.779c0.322,0.228 0.737,0.285 1.107,0.145c0.37,-0.141 0.642,-0.457 0.724,-0.84c0.869,-4.084 2.977,-14.421 3.768,-18.136c0.06,-0.28 -0.04,-0.571 -0.26,-0.758c-0.22,-0.187 -0.525,-0.241 -0.797,-0.14c-4.193,1.552 -17.106,6.397 -22.384,8.35c-0.335,0.124 -0.553,0.446 -0.542,0.799c0.012,0.354 0.25,0.661 0.593,0.764c2.367,0.708 5.474,1.693 5.474,1.693c0,0 1.452,4.385 2.209,6.615c0.095,0.28 0.314,0.5 0.603,0.576c0.288,0.075 0.596,-0.004 0.811,-0.207c1.216,-1.148 3.096,-2.923 3.096,-2.923c0,0 3.572,2.619 5.598,4.062Zm-11.01,-8.677l1.679,5.538l0.373,-3.507c0,0 6.487,-5.851 10.185,-9.186c0.108,-0.098 0.123,-0.262 0.033,-0.377c-0.089,-0.115 -0.253,-0.142 -0.376,-0.064c-4.286,2.737 -11.894,7.596 -11.894,7.596Z"/>
          </svg>
        </div>
      </div>
    </div>
  )
}
