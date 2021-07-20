// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, NullInt64, SelectionMode } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { UpdatePositionEventDB } from '../updatepositionevent-db'
import { UpdatePositionEventService } from '../updatepositionevent.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-updatepositioneventstable',
  templateUrl: './updatepositionevents-table.component.html',
  styleUrls: ['./updatepositionevents-table.component.css'],
})
export class UpdatePositionEventsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode

  // used if the component is called as a selection component of UpdatePositionEvent instances
  selection: SelectionModel<UpdatePositionEventDB>;
  initialSelection = new Array<UpdatePositionEventDB>();

  // the data source for the table
  updatepositionevents: UpdatePositionEventDB[];
  matTableDataSource: MatTableDataSource<UpdatePositionEventDB>

  // front repo, that will be referenced by this.updatepositionevents
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort) sort: MatSort;
  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (updatepositioneventDB: UpdatePositionEventDB, property: string) => {
      switch (property) {
        // insertion point for specific sorting accessor
        case 'Name':
          return updatepositioneventDB.Name;

        case 'Duration':
          return updatepositioneventDB.Duration;

        default:
          return UpdatePositionEventDB[property];
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (updatepositioneventDB: UpdatePositionEventDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the updatepositioneventDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += updatepositioneventDB.Name.toLowerCase()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort;
    this.matTableDataSource.paginator = this.paginator;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private updatepositioneventService: UpdatePositionEventService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of updatepositionevent instances
    public dialogRef: MatDialogRef<UpdatePositionEventsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.updatepositioneventService.UpdatePositionEventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getUpdatePositionEvents()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Duration",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Duration",
      ]
      this.selection = new SelectionModel<UpdatePositionEventDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getUpdatePositionEvents()
    this.matTableDataSource = new MatTableDataSource(this.updatepositionevents)
  }

  getUpdatePositionEvents(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.updatepositionevents = this.frontRepo.UpdatePositionEvents_array;

        // insertion point for variables Recoveries
        // compute strings for durations
        for (let updatepositionevent of this.updatepositionevents) {
          updatepositionevent.Duration_string =
            Math.floor(updatepositionevent.Duration / (3600 * 1000 * 1000 * 1000)) + "H " +
            Math.floor(updatepositionevent.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000)) + "M " +
            updatepositionevent.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000) + "S"
        }

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          this.updatepositionevents.forEach(
            updatepositionevent => {
              let ID = this.dialogData.ID
              let revPointer = updatepositionevent[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(updatepositionevent)
              }
            }
          )
          this.selection = new SelectionModel<UpdatePositionEventDB>(allowMultiSelect, this.initialSelection);
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s"]
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)

          if (sourceInstance[this.dialogData.SourceField]) {
            for (let associationInstance of sourceInstance[this.dialogData.SourceField]) {
              let updatepositionevent = associationInstance[this.dialogData.IntermediateStructField]
              this.initialSelection.push(updatepositionevent)
            }
          }
          this.selection = new SelectionModel<UpdatePositionEventDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.updatepositionevents
      }
    )
  }

  // newUpdatePositionEvent initiate a new updatepositionevent
  // create a new UpdatePositionEvent objet
  newUpdatePositionEvent() {
  }

  deleteUpdatePositionEvent(updatepositioneventID: number, updatepositionevent: UpdatePositionEventDB) {
    // list of updatepositionevents is truncated of updatepositionevent before the delete
    this.updatepositionevents = this.updatepositionevents.filter(h => h !== updatepositionevent);

    this.updatepositioneventService.deleteUpdatePositionEvent(updatepositioneventID).subscribe(
      updatepositionevent => {
        this.updatepositioneventService.UpdatePositionEventServiceChanged.next("delete")
      }
    );
  }

  editUpdatePositionEvent(updatepositioneventID: number, updatepositionevent: UpdatePositionEventDB) {

  }

  // display updatepositionevent in router
  displayUpdatePositionEventInRouter(updatepositioneventID: number) {
    this.router.navigate(["github_com_fullstack_lang_ladybugsim_go-" + "updatepositionevent-display", updatepositioneventID])
  }

  // set editor outlet
  setEditorRouterOutlet(updatepositioneventID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "updatepositionevent-detail", updatepositioneventID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(updatepositioneventID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_presentation: ["github_com_fullstack_lang_ladybugsim_go-" + "updatepositionevent-presentation", updatepositioneventID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.updatepositionevents.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.updatepositionevents.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<UpdatePositionEventDB>()

      // reset all initial selection of updatepositionevent that belong to updatepositionevent
      this.initialSelection.forEach(
        updatepositionevent => {
          updatepositionevent[this.dialogData.ReversePointer].Int64 = 0
          updatepositionevent[this.dialogData.ReversePointer].Valid = true
          toUpdate.add(updatepositionevent)
        }
      )

      // from selection, set updatepositionevent that belong to updatepositionevent
      this.selection.selected.forEach(
        updatepositionevent => {
          let ID = +this.dialogData.ID
          updatepositionevent[this.dialogData.ReversePointer].Int64 = ID
          updatepositionevent[this.dialogData.ReversePointer].Valid = true
          toUpdate.add(updatepositionevent)
        }
      )

      // update all updatepositionevent (only update selection & initial selection)
      toUpdate.forEach(
        updatepositionevent => {
          this.updatepositioneventService.updateUpdatePositionEvent(updatepositionevent)
            .subscribe(updatepositionevent => {
              this.updatepositioneventService.UpdatePositionEventServiceChanged.next("update")
            });
        }
      )
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s"]
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedUpdatePositionEvent = new Set<number>()
      for (let updatepositionevent of this.initialSelection) {
        if (this.selection.selected.includes(updatepositionevent)) {
          // console.log("updatepositionevent " + updatepositionevent.Name + " is still selected")
        } else {
          console.log("updatepositionevent " + updatepositionevent.Name + " has been unselected")
          unselectedUpdatePositionEvent.add(updatepositionevent.ID)
          console.log("is unselected " + unselectedUpdatePositionEvent.has(updatepositionevent.ID))
        }
      }

      // delete the association instance
      if (sourceInstance[this.dialogData.SourceField]) {
        for (let associationInstance of sourceInstance[this.dialogData.SourceField]) {
          let updatepositionevent = associationInstance[this.dialogData.IntermediateStructField]
          if (unselectedUpdatePositionEvent.has(updatepositionevent.ID)) {

            this.frontRepoService.deleteService( this.dialogData.IntermediateStruct, associationInstance )
          }
        }
      }

      // is the source array is emptyn create it
      if (sourceInstance[this.dialogData.SourceField] == undefined) {
        sourceInstance[this.dialogData.SourceField] = new Array<any>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField]) {
        this.selection.selected.forEach(
          updatepositionevent => {
            if (!this.initialSelection.includes(updatepositionevent)) {
              // console.log("updatepositionevent " + updatepositionevent.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + updatepositionevent.Name,
              }

              associationInstance[this.dialogData.IntermediateStructField+"ID"] = new NullInt64
              associationInstance[this.dialogData.IntermediateStructField+"ID"].Int64 = updatepositionevent.ID
              associationInstance[this.dialogData.IntermediateStructField+"ID"].Valid = true

              associationInstance[this.dialogData.SourceStruct + "_" + this.dialogData.SourceField + "DBID"] = new NullInt64
              associationInstance[this.dialogData.SourceStruct + "_" + this.dialogData.SourceField + "DBID"].Int64 = sourceInstance["ID"]
              associationInstance[this.dialogData.SourceStruct + "_" + this.dialogData.SourceField + "DBID"].Valid = true

              this.frontRepoService.postService( this.dialogData.IntermediateStruct, associationInstance )

            } else {
              // console.log("updatepositionevent " + updatepositionevent.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<UpdatePositionEventDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
