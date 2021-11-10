import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { TextDB } from '../text-db'
import { TextService } from '../text.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface textDummyElement {
}

const ELEMENT_DATA: textDummyElement[] = [
];

@Component({
	selector: 'app-text-presentation',
	templateUrl: './text-presentation.component.html',
	styleUrls: ['./text-presentation.component.css'],
})
export class TextPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	text: TextDB = new (TextDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private textService: TextService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getText();

		// observable for changes in 
		this.textService.TextServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getText()
				}
			}
		)
	}

	getText(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.text = this.frontRepo.Texts.get(id)!

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
				github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "text-detail", ID]
			}
		}]);
	}
}
