// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { PolylineDB } from '../polyline-db'
import { PolylineService } from '../polyline.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-polyline-sorting',
  templateUrl: './polyline-sorting.component.html',
  styleUrls: ['./polyline-sorting.component.css']
})
export class PolylineSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Polyline instances that are in the association
  associatedPolylines = new Array<PolylineDB>();

  constructor(
    private polylineService: PolylineService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of polyline instances
    public dialogRef: MatDialogRef<PolylineSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getPolylines()
  }

  getPolylines(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let polyline of this.frontRepo.Polylines_array) {
          let ID = this.dialogData.ID
          let revPointerID = polyline[this.dialogData.ReversePointer as keyof PolylineDB] as unknown as NullInt64
          let revPointerID_Index = polyline[this.dialogData.ReversePointer + "_Index" as keyof PolylineDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedPolylines.push(polyline)
          }
        }

        // sort associated polyline according to order
        this.associatedPolylines.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedPolylines, event.previousIndex, event.currentIndex);

    // set the order of Polyline instances
    let index = 0

    for (let polyline of this.associatedPolylines) {
      let revPointerID_Index = polyline[this.dialogData.ReversePointer + "_Index" as keyof PolylineDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedPolylines.forEach(
      polyline => {
        this.polylineService.updatePolyline(polyline)
          .subscribe(polyline => {
            this.polylineService.PolylineServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}
