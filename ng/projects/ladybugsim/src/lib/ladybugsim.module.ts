import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { LadybugsTableComponent } from './ladybugs-table/ladybugs-table.component'
import { LadybugSortingComponent } from './ladybug-sorting/ladybug-sorting.component'
import { LadybugDetailComponent } from './ladybug-detail/ladybug-detail.component'
import { LadybugPresentationComponent } from './ladybug-presentation/ladybug-presentation.component'

import { LadybugSimulationsTableComponent } from './ladybugsimulations-table/ladybugsimulations-table.component'
import { LadybugSimulationSortingComponent } from './ladybugsimulation-sorting/ladybugsimulation-sorting.component'
import { LadybugSimulationDetailComponent } from './ladybugsimulation-detail/ladybugsimulation-detail.component'
import { LadybugSimulationPresentationComponent } from './ladybugsimulation-presentation/ladybugsimulation-presentation.component'

import { UpdatePositionEventsTableComponent } from './updatepositionevents-table/updatepositionevents-table.component'
import { UpdatePositionEventSortingComponent } from './updatepositionevent-sorting/updatepositionevent-sorting.component'
import { UpdatePositionEventDetailComponent } from './updatepositionevent-detail/updatepositionevent-detail.component'
import { UpdatePositionEventPresentationComponent } from './updatepositionevent-presentation/updatepositionevent-presentation.component'

import { UpdateSpeedEventsTableComponent } from './updatespeedevents-table/updatespeedevents-table.component'
import { UpdateSpeedEventSortingComponent } from './updatespeedevent-sorting/updatespeedevent-sorting.component'
import { UpdateSpeedEventDetailComponent } from './updatespeedevent-detail/updatespeedevent-detail.component'
import { UpdateSpeedEventPresentationComponent } from './updatespeedevent-presentation/updatespeedevent-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		LadybugsTableComponent,
		LadybugSortingComponent,
		LadybugDetailComponent,
		LadybugPresentationComponent,

		LadybugSimulationsTableComponent,
		LadybugSimulationSortingComponent,
		LadybugSimulationDetailComponent,
		LadybugSimulationPresentationComponent,

		UpdatePositionEventsTableComponent,
		UpdatePositionEventSortingComponent,
		UpdatePositionEventDetailComponent,
		UpdatePositionEventPresentationComponent,

		UpdateSpeedEventsTableComponent,
		UpdateSpeedEventSortingComponent,
		UpdateSpeedEventDetailComponent,
		UpdateSpeedEventPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		LadybugsTableComponent,
		LadybugSortingComponent,
		LadybugDetailComponent,
		LadybugPresentationComponent,

		LadybugSimulationsTableComponent,
		LadybugSimulationSortingComponent,
		LadybugSimulationDetailComponent,
		LadybugSimulationPresentationComponent,

		UpdatePositionEventsTableComponent,
		UpdatePositionEventSortingComponent,
		UpdatePositionEventDetailComponent,
		UpdatePositionEventPresentationComponent,

		UpdateSpeedEventsTableComponent,
		UpdateSpeedEventSortingComponent,
		UpdateSpeedEventDetailComponent,
		UpdateSpeedEventPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class LadybugsimModule { }
