import { h, Fragment } from 'preact'
import { useCallback } from 'preact/hooks'
import { res } from '../common/static'
import { HeadUpDisplay } from '../hud/HeadUpDisplay'

// import { Button } from 'preact-fluid';
export function Auth(props) {
  const nav = e => {
    e.preventDefault()
  }
  return (
    <div className='route'>
      <div class='container home'>
        <h1 class='tracking-in-08'>Radar auth processing</h1>
        {/*<div class="row tracking-in-03">*/}
        {/*  /!*<div class="column column-50 column-offset-25">.column column-50 column-offset-25</div>*!/*/}
        {/*</div>*/}
      </div>
    </div>
  )
}
