import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PolylineDB } from '../polyline-db'
import { PolylineService } from '../polyline.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface polylineDummyElement {
}

const ELEMENT_DATA: polylineDummyElement[] = [
];

@Component({
	selector: 'app-polyline-presentation',
	templateUrl: './polyline-presentation.component.html',
	styleUrls: ['./polyline-presentation.component.css'],
})
export class PolylinePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	polyline: PolylineDB = new (PolylineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private polylineService: PolylineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPolyline();

		// observable for changes in 
		this.polylineService.PolylineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPolyline()
				}
			}
		)
	}

	getPolyline(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.polyline = this.frontRepo.Polylines.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "polyline-detail", ID]
			}
		}]);
	}
}
