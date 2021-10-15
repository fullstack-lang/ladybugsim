import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LadybugSimulationDB } from '../ladybugsimulation-db'
import { LadybugSimulationService } from '../ladybugsimulation.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface ladybugsimulationDummyElement {
}

const ELEMENT_DATA: ladybugsimulationDummyElement[] = [
];

@Component({
	selector: 'app-ladybugsimulation-presentation',
	templateUrl: './ladybugsimulation-presentation.component.html',
	styleUrls: ['./ladybugsimulation-presentation.component.css'],
})
export class LadybugSimulationPresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from SimulationStep
	SimulationStep_Hours: number = 0
	SimulationStep_Minutes: number = 0
	SimulationStep_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	ladybugsimulation: LadybugSimulationDB = new (LadybugSimulationDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private ladybugsimulationService: LadybugSimulationService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLadybugSimulation();

		// observable for changes in 
		this.ladybugsimulationService.LadybugSimulationServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLadybugSimulation()
				}
			}
		)
	}

	getLadybugSimulation(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.ladybugsimulation = this.frontRepo.LadybugSimulations.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for SimulationStep
				this.SimulationStep_Hours = Math.floor(this.ladybugsimulation.SimulationStep / (3600 * 1000 * 1000 * 1000))
				this.SimulationStep_Minutes = Math.floor(this.ladybugsimulation.SimulationStep % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.SimulationStep_Seconds = this.ladybugsimulation.SimulationStep % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "ladybugsimulation-detail", ID]
			}
		}]);
	}
}
