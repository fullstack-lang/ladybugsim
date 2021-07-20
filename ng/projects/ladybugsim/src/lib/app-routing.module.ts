import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { LadybugsTableComponent } from './ladybugs-table/ladybugs-table.component'
import { LadybugDetailComponent } from './ladybug-detail/ladybug-detail.component'
import { LadybugPresentationComponent } from './ladybug-presentation/ladybug-presentation.component'

import { LadybugSimulationsTableComponent } from './ladybugsimulations-table/ladybugsimulations-table.component'
import { LadybugSimulationDetailComponent } from './ladybugsimulation-detail/ladybugsimulation-detail.component'
import { LadybugSimulationPresentationComponent } from './ladybugsimulation-presentation/ladybugsimulation-presentation.component'

import { UpdatePositionEventsTableComponent } from './updatepositionevents-table/updatepositionevents-table.component'
import { UpdatePositionEventDetailComponent } from './updatepositionevent-detail/updatepositionevent-detail.component'
import { UpdatePositionEventPresentationComponent } from './updatepositionevent-presentation/updatepositionevent-presentation.component'

import { UpdateSpeedEventsTableComponent } from './updatespeedevents-table/updatespeedevents-table.component'
import { UpdateSpeedEventDetailComponent } from './updatespeedevent-detail/updatespeedevent-detail.component'
import { UpdateSpeedEventPresentationComponent } from './updatespeedevent-presentation/updatespeedevent-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugs', component: LadybugsTableComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_table' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybug-adder', component: LadybugDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybug-adder/:id/:originStruct/:originStructFieldName', component: LadybugDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybug-detail/:id', component: LadybugDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybug-presentation/:id', component: LadybugPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_presentation' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybug-presentation-special/:id', component: LadybugPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_goladybugpres' },

	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulations', component: LadybugSimulationsTableComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_table' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulation-adder', component: LadybugSimulationDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulation-adder/:id/:originStruct/:originStructFieldName', component: LadybugSimulationDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulation-detail/:id', component: LadybugSimulationDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulation-presentation/:id', component: LadybugSimulationPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_presentation' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-ladybugsimulation-presentation-special/:id', component: LadybugSimulationPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_goladybugsimulationpres' },

	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevents', component: UpdatePositionEventsTableComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_table' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevent-adder', component: UpdatePositionEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevent-adder/:id/:originStruct/:originStructFieldName', component: UpdatePositionEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevent-detail/:id', component: UpdatePositionEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevent-presentation/:id', component: UpdatePositionEventPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_presentation' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatepositionevent-presentation-special/:id', component: UpdatePositionEventPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_goupdatepositioneventpres' },

	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevents', component: UpdateSpeedEventsTableComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_table' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevent-adder', component: UpdateSpeedEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevent-adder/:id/:originStruct/:originStructFieldName', component: UpdateSpeedEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevent-detail/:id', component: UpdateSpeedEventDetailComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_editor' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevent-presentation/:id', component: UpdateSpeedEventPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_go_presentation' },
	{ path: 'github_com_fullstack_lang_ladybugsim_go-updatespeedevent-presentation-special/:id', component: UpdateSpeedEventPresentationComponent, outlet: 'github_com_fullstack_lang_ladybugsim_goupdatespeedeventpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
