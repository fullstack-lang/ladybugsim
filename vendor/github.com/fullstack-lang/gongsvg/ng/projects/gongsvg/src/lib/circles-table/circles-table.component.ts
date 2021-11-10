// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { CircleDB } from '../circle-db'
import { CircleService } from '../circle.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-circlestable',
  templateUrl: './circles-table.component.html',
  styleUrls: ['./circles-table.component.css'],
})
export class CirclesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Circle instances
  selection: SelectionModel<CircleDB> = new (SelectionModel)
  initialSelection = new Array<CircleDB>()

  // the data source for the table
  circles: CircleDB[] = []
  matTableDataSource: MatTableDataSource<CircleDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.circles
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (circleDB: CircleDB, property: string) => {
      switch (property) {
        case 'ID':
          return circleDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return circleDB.Name;

        case 'CX':
          return circleDB.CX;

        case 'CY':
          return circleDB.CY;

        case 'Radius':
          return circleDB.Radius;

        case 'Color':
          return circleDB.Color;

        case 'FillOpacity':
          return circleDB.FillOpacity;

        case 'Stroke':
          return circleDB.Stroke;

        case 'StrokeWidth':
          return circleDB.StrokeWidth;

        case 'StrokeDashArray':
          return circleDB.StrokeDashArray;

        case 'Transform':
          return circleDB.Transform;

        case 'SVG_Circles':
          return this.frontRepo.SVGs.get(circleDB.SVG_CirclesDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (circleDB: CircleDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the circleDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += circleDB.Name.toLowerCase()
      mergedContent += circleDB.CX.toString()
      mergedContent += circleDB.CY.toString()
      mergedContent += circleDB.Radius.toString()
      mergedContent += circleDB.Color.toLowerCase()
      mergedContent += circleDB.FillOpacity.toString()
      mergedContent += circleDB.Stroke.toLowerCase()
      mergedContent += circleDB.StrokeWidth.toString()
      mergedContent += circleDB.StrokeDashArray.toLowerCase()
      mergedContent += circleDB.Transform.toLowerCase()
      if (circleDB.SVG_CirclesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.SVGs.get(circleDB.SVG_CirclesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private circleService: CircleService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of circle instances
    public dialogRef: MatDialogRef<CirclesTableComponent>,
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
    this.circleService.CircleServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getCircles()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "CX",
        "CY",
        "Radius",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "Transform",
        "SVG_Circles",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "CX",
        "CY",
        "Radius",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "Transform",
        "SVG_Circles",
      ]
      this.selection = new SelectionModel<CircleDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getCircles()
    this.matTableDataSource = new MatTableDataSource(this.circles)
  }

  getCircles(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.circles = this.frontRepo.Circles_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let circle of this.circles) {
            let ID = this.dialogData.ID
            let revPointer = circle[this.dialogData.ReversePointer as keyof CircleDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(circle)
            }
            this.selection = new SelectionModel<CircleDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, CircleDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as CircleDB[]
          for (let associationInstance of sourceField) {
            let circle = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as CircleDB
            this.initialSelection.push(circle)
          }

          this.selection = new SelectionModel<CircleDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.circles
      }
    )
  }

  // newCircle initiate a new circle
  // create a new Circle objet
  newCircle() {
  }

  deleteCircle(circleID: number, circle: CircleDB) {
    // list of circles is truncated of circle before the delete
    this.circles = this.circles.filter(h => h !== circle);

    this.circleService.deleteCircle(circleID).subscribe(
      circle => {
        this.circleService.CircleServiceChanged.next("delete")
      }
    );
  }

  editCircle(circleID: number, circle: CircleDB) {

  }

  // display circle in router
  displayCircleInRouter(circleID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongsvg_go-" + "circle-display", circleID])
  }

  // set editor outlet
  setEditorRouterOutlet(circleID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "circle-detail", circleID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(circleID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsvg_go_presentation: ["github_com_fullstack_lang_gongsvg_go-" + "circle-presentation", circleID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.circles.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.circles.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<CircleDB>()

      // reset all initial selection of circle that belong to circle
      for (let circle of this.initialSelection) {
        let index = circle[this.dialogData.ReversePointer as keyof CircleDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(circle)

      }

      // from selection, set circle that belong to circle
      for (let circle of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = circle[this.dialogData.ReversePointer as keyof CircleDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(circle)
      }


      // update all circle (only update selection & initial selection)
      for (let circle of toUpdate) {
        this.circleService.updateCircle(circle)
          .subscribe(circle => {
            this.circleService.CircleServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, CircleDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedCircle = new Set<number>()
      for (let circle of this.initialSelection) {
        if (this.selection.selected.includes(circle)) {
          // console.log("circle " + circle.Name + " is still selected")
        } else {
          console.log("circle " + circle.Name + " has been unselected")
          unselectedCircle.add(circle.ID)
          console.log("is unselected " + unselectedCircle.has(circle.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let circle = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as CircleDB
      if (unselectedCircle.has(circle.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<CircleDB>) = new Array<CircleDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          circle => {
            if (!this.initialSelection.includes(circle)) {
              // console.log("circle " + circle.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + circle.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = circle.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = circle.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("circle " + circle.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<CircleDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}