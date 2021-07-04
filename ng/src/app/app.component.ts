import { Component } from '@angular/core';

import * as ladybugsim from 'ladybugsim'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ng';

  constructor(
    private ladybugService: ladybugsim.LadybugService) {
  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (updateDisplay: boolean): void => {

    // refresh the ladybug splitter
    this.ladybugService.LadybugServiceChanged.next("update")
  }
}
