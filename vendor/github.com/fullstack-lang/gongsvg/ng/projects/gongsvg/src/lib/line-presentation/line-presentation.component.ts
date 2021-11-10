import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LineDB } from '../line-db'
import { LineService } from '../line.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface lineDummyElement {
}

const ELEMENT_DATA: lineDummyElement[] = [
];

@Component({
	selector: 'app-line-presentation',
	templateUrl: './line-presentation.component.html',
	styleUrls: ['./line-presentation.component.css'],
})
export class LinePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	line: LineDB = new (LineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private lineService: LineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLine();

		// observable for changes in 
		this.lineService.LineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLine()
				}
			}
		)
	}

	getLine(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.line = this.frontRepo.Lines.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "line-detail", ID]
			}
		}]);
	}
}
