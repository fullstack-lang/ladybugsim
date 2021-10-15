import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { UpdateSpeedEventDB } from '../updatespeedevent-db'
import { UpdateSpeedEventService } from '../updatespeedevent.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface updatespeedeventDummyElement {
}

const ELEMENT_DATA: updatespeedeventDummyElement[] = [
];

@Component({
	selector: 'app-updatespeedevent-presentation',
	templateUrl: './updatespeedevent-presentation.component.html',
	styleUrls: ['./updatespeedevent-presentation.component.css'],
})
export class UpdateSpeedEventPresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from Duration
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	updatespeedevent: UpdateSpeedEventDB = new (UpdateSpeedEventDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private updatespeedeventService: UpdateSpeedEventService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getUpdateSpeedEvent();

		// observable for changes in 
		this.updatespeedeventService.UpdateSpeedEventServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getUpdateSpeedEvent()
				}
			}
		)
	}

	getUpdateSpeedEvent(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.updatespeedevent = this.frontRepo.UpdateSpeedEvents.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.updatespeedevent.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.updatespeedevent.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.updatespeedevent.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "updatespeedevent-detail", ID]
			}
		}]);
	}
}
