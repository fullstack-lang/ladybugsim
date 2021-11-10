import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { EllipseDB } from '../ellipse-db'
import { EllipseService } from '../ellipse.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface ellipseDummyElement {
}

const ELEMENT_DATA: ellipseDummyElement[] = [
];

@Component({
	selector: 'app-ellipse-presentation',
	templateUrl: './ellipse-presentation.component.html',
	styleUrls: ['./ellipse-presentation.component.css'],
})
export class EllipsePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	ellipse: EllipseDB = new (EllipseDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private ellipseService: EllipseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEllipse();

		// observable for changes in 
		this.ellipseService.EllipseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEllipse()
				}
			}
		)
	}

	getEllipse(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.ellipse = this.frontRepo.Ellipses.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "ellipse-detail", ID]
			}
		}]);
	}
}
