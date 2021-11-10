import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PolygoneDB } from '../polygone-db'
import { PolygoneService } from '../polygone.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface polygoneDummyElement {
}

const ELEMENT_DATA: polygoneDummyElement[] = [
];

@Component({
	selector: 'app-polygone-presentation',
	templateUrl: './polygone-presentation.component.html',
	styleUrls: ['./polygone-presentation.component.css'],
})
export class PolygonePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	polygone: PolygoneDB = new (PolygoneDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private polygoneService: PolygoneService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPolygone();

		// observable for changes in 
		this.polygoneService.PolygoneServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPolygone()
				}
			}
		)
	}

	getPolygone(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.polygone = this.frontRepo.Polygones.get(id)!

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "polygone-detail", ID]
			}
		}]);
	}
}
