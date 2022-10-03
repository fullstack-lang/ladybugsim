import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { CircleDB } from '../circle-db'
import { CircleService } from '../circle.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface circleDummyElement {
}

const ELEMENT_DATA: circleDummyElement[] = [
];

@Component({
	selector: 'app-circle-presentation',
	templateUrl: './circle-presentation.component.html',
	styleUrls: ['./circle-presentation.component.css'],
})
export class CirclePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	circle: CircleDB = new (CircleDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private circleService: CircleService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getCircle();

		// observable for changes in 
		this.circleService.CircleServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getCircle()
				}
			}
		)
	}

	getCircle(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.circle = this.frontRepo.Circles.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongsvg_go_presentation: ["github_com_fullstack_lang_gongsvg_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "circle-detail", ID]
			}
		}]);
	}
}
