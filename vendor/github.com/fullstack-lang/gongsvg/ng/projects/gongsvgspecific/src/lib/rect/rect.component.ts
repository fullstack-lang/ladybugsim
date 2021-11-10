import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit {

  @Input() Rect?: gongsvg.RectDB

  constructor() { }

  ngOnInit(): void {
  }
}
