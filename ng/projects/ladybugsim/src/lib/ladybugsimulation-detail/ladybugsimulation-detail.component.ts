// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LadybugSimulationDB } from '../ladybugsimulation-db'
import { LadybugSimulationService } from '../ladybugsimulation.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LadybugSimulationDetailComponent is initilizaed from different routes
// LadybugSimulationDetailComponentState detail different cases 
enum LadybugSimulationDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-ladybugsimulation-detail',
	templateUrl: './ladybugsimulation-detail.component.html',
	styleUrls: ['./ladybugsimulation-detail.component.css'],
})
export class LadybugSimulationDetailComponent implements OnInit {

	// insertion point for declarations
	SimulationStep_Hours: number = 0
	SimulationStep_Minutes: number = 0
	SimulationStep_Seconds: number = 0

	// the LadybugSimulationDB of interest
	ladybugsimulation: LadybugSimulationDB = new LadybugSimulationDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LadybugSimulationDetailComponentState = LadybugSimulationDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private ladybugsimulationService: LadybugSimulationService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = LadybugSimulationDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LadybugSimulationDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLadybugSimulation()

		// observable for changes in structs
		this.ladybugsimulationService.LadybugSimulationServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLadybugSimulation()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLadybugSimulation(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LadybugSimulationDetailComponentState.CREATE_INSTANCE:
						this.ladybugsimulation = new (LadybugSimulationDB)
						break;
					case LadybugSimulationDetailComponentState.UPDATE_INSTANCE:
						let ladybugsimulation = frontRepo.LadybugSimulations.get(this.id)
						console.assert(ladybugsimulation != undefined, "missing ladybugsimulation with id:" + this.id)
						this.ladybugsimulation = ladybugsimulation!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				// computation of Hours, Minutes, Seconds for SimulationStep
				this.SimulationStep_Hours = Math.floor(this.ladybugsimulation.SimulationStep / (3600 * 1000 * 1000 * 1000))
				this.SimulationStep_Minutes = Math.floor(this.ladybugsimulation.SimulationStep % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.SimulationStep_Seconds = this.ladybugsimulation.SimulationStep % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.ladybugsimulation.SimulationStep =
			this.SimulationStep_Hours * (3600 * 1000 * 1000 * 1000) +
			this.SimulationStep_Minutes * (60 * 1000 * 1000 * 1000) +
			this.SimulationStep_Seconds * (1000 * 1000 * 1000)

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case LadybugSimulationDetailComponentState.UPDATE_INSTANCE:
				this.ladybugsimulationService.updateLadybugSimulation(this.ladybugsimulation)
					.subscribe(ladybugsimulation => {
						this.ladybugsimulationService.LadybugSimulationServiceChanged.next("update")
					});
				break;
			default:
				this.ladybugsimulationService.postLadybugSimulation(this.ladybugsimulation).subscribe(ladybugsimulation => {
					this.ladybugsimulationService.LadybugSimulationServiceChanged.next("post")
					this.ladybugsimulation = new (LadybugSimulationDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.ladybugsimulation.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.ladybugsimulation.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "LadybugSimulation"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.ladybugsimulation.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.ladybugsimulation.Name == "") {
			this.ladybugsimulation.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
