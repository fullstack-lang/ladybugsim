import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AnimateDB } from '../animate-db'
import { AnimateService } from '../animate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface animateDummyElement {
}

const ELEMENT_DATA: animateDummyElement[] = [
];

@Component({
	selector: 'app-animate-presentation',
	templateUrl: './animate-presentation.component.html',
	styleUrls: ['./animate-presentation.component.css'],
})
export class AnimatePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	animate: AnimateDB = new (AnimateDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private animateService: AnimateService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAnimate();

		// observable for changes in 
		this.animateService.AnimateServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAnimate()
				}
			}
		)
	}

	getAnimate(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.animate = this.frontRepo.Animates.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "animate-detail", ID]
			}
		}]);
	}
}
