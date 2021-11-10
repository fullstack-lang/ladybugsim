import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PathDB } from '../path-db'
import { PathService } from '../path.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface pathDummyElement {
}

const ELEMENT_DATA: pathDummyElement[] = [
];

@Component({
	selector: 'app-path-presentation',
	templateUrl: './path-presentation.component.html',
	styleUrls: ['./path-presentation.component.css'],
})
export class PathPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	path: PathDB = new (PathDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private pathService: PathService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPath();

		// observable for changes in 
		this.pathService.PathServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPath()
				}
			}
		)
	}

	getPath(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.path = this.frontRepo.Paths.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "path-detail", ID]
			}
		}]);
	}
}
