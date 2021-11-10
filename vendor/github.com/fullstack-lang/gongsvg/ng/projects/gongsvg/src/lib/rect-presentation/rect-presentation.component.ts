import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { RectDB } from '../rect-db'
import { RectService } from '../rect.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface rectDummyElement {
}

const ELEMENT_DATA: rectDummyElement[] = [
];

@Component({
	selector: 'app-rect-presentation',
	templateUrl: './rect-presentation.component.html',
	styleUrls: ['./rect-presentation.component.css'],
})
export class RectPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	rect: RectDB = new (RectDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private rectService: RectService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getRect();

		// observable for changes in 
		this.rectService.RectServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getRect()
				}
			}
		)
	}

	getRect(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.rect = this.frontRepo.Rects.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "rect-detail", ID]
			}
		}]);
	}
}
