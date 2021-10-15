import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LadybugDB } from '../ladybug-db'
import { LadybugService } from '../ladybug.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface ladybugDummyElement {
}

const ELEMENT_DATA: ladybugDummyElement[] = [
];

@Component({
	selector: 'app-ladybug-presentation',
	templateUrl: './ladybug-presentation.component.html',
	styleUrls: ['./ladybug-presentation.component.css'],
})
export class LadybugPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	ladybug: LadybugDB = new (LadybugDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private ladybugService: LadybugService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLadybug();

		// observable for changes in 
		this.ladybugService.LadybugServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLadybug()
				}
			}
		)
	}

	getLadybug(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.ladybug = this.frontRepo.Ladybugs.get(id)!

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + "ladybug-detail", ID]
			}
		}]);
	}
}
