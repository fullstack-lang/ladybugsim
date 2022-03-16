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
import { LadybugSimulationDB } from '../ladybugsimulation-db'
import { LadybugSimulationService } from '../ladybugsimulation.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-ladybugsimulationstable',
  templateUrl: './ladybugsimulations-table.component.html',
  styleUrls: ['./ladybugsimulations-table.component.css'],
})
export class LadybugSimulationsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of LadybugSimulation instances
  selection: SelectionModel<LadybugSimulationDB> = new (SelectionModel)
  initialSelection = new Array<LadybugSimulationDB>()

  // the data source for the table
  ladybugsimulations: LadybugSimulationDB[] = []
  matTableDataSource: MatTableDataSource<LadybugSimulationDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.ladybugsimulations
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
    this.matTableDataSource.sortingDataAccessor = (ladybugsimulationDB: LadybugSimulationDB, property: string) => {
      switch (property) {
        case 'ID':
          return ladybugsimulationDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return ladybugsimulationDB.Name;

        case 'EventNb':
          return ladybugsimulationDB.EventNb;

        case 'NbOfCollision':
          return ladybugsimulationDB.NbOfCollision;

        case 'LadybugRadius':
          return ladybugsimulationDB.LadybugRadius;

        case 'AbsoluteSpeed':
          return ladybugsimulationDB.AbsoluteSpeed;

        case 'SimulationStep':
          return ladybugsimulationDB.SimulationStep;

        case 'MaxDistanceInOneStep':
          return ladybugsimulationDB.MaxDistanceInOneStep;

        case 'NbLadybugs':
          return ladybugsimulationDB.NbLadybugs;

        case 'NbLadybugsOnTheGround':
          return ladybugsimulationDB.NbLadybugsOnTheGround;

        case 'LeftRelayInitialPosX':
          return ladybugsimulationDB.LeftRelayInitialPosX;

        case 'RightRelayInitialPosX':
          return ladybugsimulationDB.RightRelayInitialPosX;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (ladybugsimulationDB: LadybugSimulationDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the ladybugsimulationDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += ladybugsimulationDB.Name.toLowerCase()
      mergedContent += ladybugsimulationDB.EventNb.toString()
      mergedContent += ladybugsimulationDB.NbOfCollision.toString()
      mergedContent += ladybugsimulationDB.LadybugRadius.toString()
      mergedContent += ladybugsimulationDB.AbsoluteSpeed.toString()
      mergedContent += ladybugsimulationDB.MaxDistanceInOneStep.toString()
      mergedContent += ladybugsimulationDB.NbLadybugs.toString()
      mergedContent += ladybugsimulationDB.NbLadybugsOnTheGround.toString()
      mergedContent += ladybugsimulationDB.LeftRelayInitialPosX.toString()
      mergedContent += ladybugsimulationDB.RightRelayInitialPosX.toString()

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
    private ladybugsimulationService: LadybugSimulationService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of ladybugsimulation instances
    public dialogRef: MatDialogRef<LadybugSimulationsTableComponent>,
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
    this.ladybugsimulationService.LadybugSimulationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getLadybugSimulations()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "EventNb",
        "NbOfCollision",
        "LadybugRadius",
        "AbsoluteSpeed",
        "SimulationStep",
        "MaxDistanceInOneStep",
        "NbLadybugs",
        "NbLadybugsOnTheGround",
        "LeftRelayInitialPosX",
        "RightRelayInitialPosX",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "EventNb",
        "NbOfCollision",
        "LadybugRadius",
        "AbsoluteSpeed",
        "SimulationStep",
        "MaxDistanceInOneStep",
        "NbLadybugs",
        "NbLadybugsOnTheGround",
        "LeftRelayInitialPosX",
        "RightRelayInitialPosX",
      ]
      this.selection = new SelectionModel<LadybugSimulationDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getLadybugSimulations()
    this.matTableDataSource = new MatTableDataSource(this.ladybugsimulations)
  }

  getLadybugSimulations(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.ladybugsimulations = this.frontRepo.LadybugSimulations_array;

        // insertion point for time duration Recoveries
        // compute strings for durations
        for (let ladybugsimulation of this.ladybugsimulations) {
          ladybugsimulation.SimulationStep_string =
            Math.floor(ladybugsimulation.SimulationStep / (3600 * 1000 * 1000 * 1000)) + "H " +
            Math.floor(ladybugsimulation.SimulationStep % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000)) + "M " +
            ladybugsimulation.SimulationStep % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000) + "S"
        }
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let ladybugsimulation of this.ladybugsimulations) {
            let ID = this.dialogData.ID
            let revPointer = ladybugsimulation[this.dialogData.ReversePointer as keyof LadybugSimulationDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(ladybugsimulation)
            }
            this.selection = new SelectionModel<LadybugSimulationDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LadybugSimulationDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as LadybugSimulationDB[]
          for (let associationInstance of sourceField) {
            let ladybugsimulation = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LadybugSimulationDB
            this.initialSelection.push(ladybugsimulation)
          }

          this.selection = new SelectionModel<LadybugSimulationDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.ladybugsimulations
      }
    )
  }

  // newLadybugSimulation initiate a new ladybugsimulation
  // create a new LadybugSimulation objet
  newLadybugSimulation() {
  }

  deleteLadybugSimulation(ladybugsimulationID: number, ladybugsimulation: LadybugSimulationDB) {
    // list of ladybugsimulations is truncated of ladybugsimulation before the delete
    this.ladybugsimulations = this.ladybugsimulations.filter(h => h !== ladybugsimulation);

    this.ladybugsimulationService.deleteLadybugSimulation(ladybugsimulationID).subscribe(
      ladybugsimulation => {
        this.ladybugsimulationService.LadybugSimulationServiceChanged.next("delete")
      }
    );
  }

  editLadybugSimulation(ladybugsimulationID: number, ladybugsimulation: LadybugSimulationDB) {

  }

  // display ladybugsimulation in router
  displayLadybugSimulationInRouter(ladybugsimulationID: number) {
    this.router.navigate(["github_com_fullstack_lang_ladybugsim_go-" + "ladybugsimulation-display", ladybugsimulationID])
  }

  // set editor outlet
  setEditorRouterOutlet(ladybugsimulationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "ladybugsimulation-detail", ladybugsimulationID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(ladybugsimulationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_presentation: ["github_com_fullstack_lang_ladybugsim_go-" + "ladybugsimulation-presentation", ladybugsimulationID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.ladybugsimulations.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.ladybugsimulations.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<LadybugSimulationDB>()

      // reset all initial selection of ladybugsimulation that belong to ladybugsimulation
      for (let ladybugsimulation of this.initialSelection) {
        let index = ladybugsimulation[this.dialogData.ReversePointer as keyof LadybugSimulationDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(ladybugsimulation)

      }

      // from selection, set ladybugsimulation that belong to ladybugsimulation
      for (let ladybugsimulation of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = ladybugsimulation[this.dialogData.ReversePointer as keyof LadybugSimulationDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(ladybugsimulation)
      }


      // update all ladybugsimulation (only update selection & initial selection)
      for (let ladybugsimulation of toUpdate) {
        this.ladybugsimulationService.updateLadybugSimulation(ladybugsimulation)
          .subscribe(ladybugsimulation => {
            this.ladybugsimulationService.LadybugSimulationServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LadybugSimulationDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedLadybugSimulation = new Set<number>()
      for (let ladybugsimulation of this.initialSelection) {
        if (this.selection.selected.includes(ladybugsimulation)) {
          // console.log("ladybugsimulation " + ladybugsimulation.Name + " is still selected")
        } else {
          console.log("ladybugsimulation " + ladybugsimulation.Name + " has been unselected")
          unselectedLadybugSimulation.add(ladybugsimulation.ID)
          console.log("is unselected " + unselectedLadybugSimulation.has(ladybugsimulation.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let ladybugsimulation = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LadybugSimulationDB
      if (unselectedLadybugSimulation.has(ladybugsimulation.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<LadybugSimulationDB>) = new Array<LadybugSimulationDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          ladybugsimulation => {
            if (!this.initialSelection.includes(ladybugsimulation)) {
              // console.log("ladybugsimulation " + ladybugsimulation.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + ladybugsimulation.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = ladybugsimulation.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = ladybugsimulation.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("ladybugsimulation " + ladybugsimulation.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<LadybugSimulationDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
