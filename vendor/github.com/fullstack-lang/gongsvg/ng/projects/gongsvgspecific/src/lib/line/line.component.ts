import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'lib-line',
  templateUrl: './line.component.svg',
  styleUrls: ['./line.component.css']
})
export class LineComponent implements OnInit {

  @Input() Line?: gongsvg.LineDB

  constructor() { }

  ngOnInit(): void {

  }

}
