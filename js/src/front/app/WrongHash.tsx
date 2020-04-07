import { h } from 'preact'
import { store } from 'src/front/store/frontStore'

export function WrongHash() {
  const goHome = e => {
    e.preventDefault()
    store.atoms.routes.current('/')
  }
  return (
    <div className='container home'>
      <h1 className='tracking-in-08'>Radar</h1>
      <div className="row tracking-in-03">
        Ссылка недействительна.
      </div>
      <div style='padding-top:20px'>
        <a className="button button-outline" href="/" onClick={goHome}>Reload</a>
      </div>
    </div>
  )
}
