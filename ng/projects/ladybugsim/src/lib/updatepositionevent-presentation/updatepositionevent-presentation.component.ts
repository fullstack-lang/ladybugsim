import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { UpdatePositionEventDB } from '../updatepositionevent-db'
import { UpdatePositionEventService } from '../updatepositionevent.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface updatepositioneventDummyElement {
}

const ELEMENT_DATA: updatepositioneventDummyElement[] = [
];

@Component({
	selector: 'app-updatepositionevent-presentation',
	templateUrl: './updatepositionevent-presentation.component.html',
	styleUrls: ['./updatepositionevent-presentation.component.css'],
})
export class UpdatePositionEventPresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from Duration
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	updatepositionevent: UpdatePositionEventDB = new (UpdatePositionEventDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private updatepositioneventService: UpdatePositionEventService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUpdatePositionEvent();

		// observable for changes in 
		this.updatepositioneventService.UpdatePositionEventServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUpdatePositionEvent()
				}
			}
		)
	}

	getUpdatePositionEvent(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.updatepositionevent = this.frontRepo.UpdatePositionEvents.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.updatepositionevent.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.updatepositionevent.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.updatepositionevent.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_ladybugsim_go_presentation: ["github_com_fullstack_lang_ladybugsim_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "updatepositionevent-detail", ID]
			}
		}]);
	}
}
