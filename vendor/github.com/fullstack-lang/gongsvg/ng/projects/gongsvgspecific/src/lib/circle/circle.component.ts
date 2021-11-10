import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'
import { CircleDB } from 'gongsvg';
@Component({
  selector: 'lib-circle',
  templateUrl: './circle.component.svg',
  styleUrls: ['./circle.component.css']
})
export class CircleComponent implements OnInit {

  @Input() Circle?: gongsvg.CircleDB

  constructor() { }

  ngOnInit(): void {
  }

}
