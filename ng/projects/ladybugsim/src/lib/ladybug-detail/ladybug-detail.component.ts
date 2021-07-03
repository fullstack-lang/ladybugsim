// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LadybugDB } from '../ladybug-db'
import { LadybugService } from '../ladybug.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// LadybugDetailComponent is initilizaed from different routes
// LadybugDetailComponentState detail different cases 
enum LadybugDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_LadybugSimulation_Ladybugs_SET,
}

@Component({
	selector: 'app-ladybug-detail',
	templateUrl: './ladybug-detail.component.html',
	styleUrls: ['./ladybug-detail.component.css'],
})
export class LadybugDetailComponent implements OnInit {

	// insertion point for declarations

	// the LadybugDB of interest
	ladybug: LadybugDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LadybugDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private ladybugService: LadybugService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id');
		this.originStruct = this.route.snapshot.paramMap.get('originStruct');
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName');

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = LadybugDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LadybugDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Ladybugs":
						console.log("Ladybug" + " is instanciated with back pointer to instance " + this.id + " LadybugSimulation association Ladybugs")
						this.state = LadybugDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_LadybugSimulation_Ladybugs_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLadybug()

		// observable for changes in structs
		this.ladybugService.LadybugServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLadybug()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLadybug(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LadybugDetailComponentState.CREATE_INSTANCE:
						this.ladybug = new (LadybugDB)
						break;
					case LadybugDetailComponentState.UPDATE_INSTANCE:
						this.ladybug = frontRepo.Ladybugs.get(this.id)
						break;
					// insertion point for init of association field
					case LadybugDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_LadybugSimulation_Ladybugs_SET:
						this.ladybug = new (LadybugDB)
						this.ladybug.LadybugSimulation_Ladybugs_reverse = frontRepo.LadybugSimulations.get(this.id)
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.ladybug.LadybugSimulation_Ladybugs_reverse != undefined) {
			if (this.ladybug.LadybugSimulation_LadybugsDBID == undefined) {
				this.ladybug.LadybugSimulation_LadybugsDBID = new NullInt64
			}
			this.ladybug.LadybugSimulation_LadybugsDBID.Int64 = this.ladybug.LadybugSimulation_Ladybugs_reverse.ID
			this.ladybug.LadybugSimulation_LadybugsDBID.Valid = true
			if (this.ladybug.LadybugSimulation_LadybugsDBID_Index == undefined) {
				this.ladybug.LadybugSimulation_LadybugsDBID_Index = new NullInt64
			}
			this.ladybug.LadybugSimulation_LadybugsDBID_Index.Valid = true
			this.ladybug.LadybugSimulation_Ladybugs_reverse = undefined // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LadybugDetailComponentState.UPDATE_INSTANCE:
				this.ladybugService.updateLadybug(this.ladybug)
					.subscribe(ladybug => {
						this.ladybugService.LadybugServiceChanged.next("update")
					});
				break;
			default:
				this.ladybugService.postLadybug(this.ladybug).subscribe(ladybug => {
					this.ladybugService.LadybugServiceChanged.next("post")
					this.ladybug = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: SelectionMode,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string ) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.ladybug.ID
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
			dialogData.ID = this.ladybug.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Ladybug"
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
			ID: this.ladybug.ID,
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

	fillUpNameIfEmpty(event) {
		if (this.ladybug.Name == undefined) {
			this.ladybug.Name = event.value.Name
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
			return this.mapFields_displayAsTextArea.get(fieldName)
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
