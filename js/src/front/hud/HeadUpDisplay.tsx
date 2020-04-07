import {Component, h} from 'preact'


import {RightControl} from "./RightControl";
import { PopUps } from '~/front/hud/PopUps'


export class HeadUpDisplay extends Component<any, any> {
  render() {
    return (
      <div class="hud">
        <RightControl/>
        <PopUps/>
      </div>
    )
  }
}
