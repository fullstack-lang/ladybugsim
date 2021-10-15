// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { LadybugDB } from '../ladybug-db'
import { LadybugService } from '../ladybug.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-ladybug-sorting',
  templateUrl: './ladybug-sorting.component.html',
  styleUrls: ['./ladybug-sorting.component.css']
})
export class LadybugSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Ladybug instances that are in the association
  associatedLadybugs = new Array<LadybugDB>();

  constructor(
    private ladybugService: LadybugService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of ladybug instances
    public dialogRef: MatDialogRef<LadybugSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getLadybugs()
  }

  getLadybugs(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let ladybug of this.frontRepo.Ladybugs_array) {
          let ID = this.dialogData.ID
          let revPointerID = ladybug[this.dialogData.ReversePointer as keyof LadybugDB] as unknown as NullInt64
          let revPointerID_Index = ladybug[this.dialogData.ReversePointer + "_Index" as keyof LadybugDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedLadybugs.push(ladybug)
          }
        }

        // sort associated ladybug according to order
        this.associatedLadybugs.sort((t1, t2) => {
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
    moveItemInArray(this.associatedLadybugs, event.previousIndex, event.currentIndex);

    // set the order of Ladybug instances
    let index = 0

    for (let ladybug of this.associatedLadybugs) {
      let revPointerID_Index = ladybug[this.dialogData.ReversePointer + "_Index" as keyof LadybugDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedLadybugs.forEach(
      ladybug => {
        this.ladybugService.updateLadybug(ladybug)
          .subscribe(ladybug => {
            this.ladybugService.LadybugServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}
