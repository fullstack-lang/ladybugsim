import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { SVGDB } from '../svg-db'
import { SVGService } from '../svg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface svgDummyElement {
}

const ELEMENT_DATA: svgDummyElement[] = [
];

@Component({
	selector: 'app-svg-presentation',
	templateUrl: './svg-presentation.component.html',
	styleUrls: ['./svg-presentation.component.css'],
})
export class SVGPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	svg: SVGDB = new (SVGDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private svgService: SVGService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSVG();

		// observable for changes in 
		this.svgService.SVGServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSVG()
				}
			}
		)
	}

	getSVG(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.svg = this.frontRepo.SVGs.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "svg-detail", ID]
			}
		}]);
	}
}
