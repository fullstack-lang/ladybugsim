import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AnimatesTableComponent } from './animates-table/animates-table.component'
import { AnimateDetailComponent } from './animate-detail/animate-detail.component'
import { AnimatePresentationComponent } from './animate-presentation/animate-presentation.component'

import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'
import { CirclePresentationComponent } from './circle-presentation/circle-presentation.component'

import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
import { EllipseDetailComponent } from './ellipse-detail/ellipse-detail.component'
import { EllipsePresentationComponent } from './ellipse-presentation/ellipse-presentation.component'

import { LinesTableComponent } from './lines-table/lines-table.component'
import { LineDetailComponent } from './line-detail/line-detail.component'
import { LinePresentationComponent } from './line-presentation/line-presentation.component'

import { PathsTableComponent } from './paths-table/paths-table.component'
import { PathDetailComponent } from './path-detail/path-detail.component'
import { PathPresentationComponent } from './path-presentation/path-presentation.component'

import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
import { PolygoneDetailComponent } from './polygone-detail/polygone-detail.component'
import { PolygonePresentationComponent } from './polygone-presentation/polygone-presentation.component'

import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
import { PolylineDetailComponent } from './polyline-detail/polyline-detail.component'
import { PolylinePresentationComponent } from './polyline-presentation/polyline-presentation.component'

import { RectsTableComponent } from './rects-table/rects-table.component'
import { RectDetailComponent } from './rect-detail/rect-detail.component'
import { RectPresentationComponent } from './rect-presentation/rect-presentation.component'

import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'
import { SVGPresentationComponent } from './svg-presentation/svg-presentation.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextDetailComponent } from './text-detail/text-detail.component'
import { TextPresentationComponent } from './text-presentation/text-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongsvg_go-animates', component: AnimatesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-adder', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-adder/:id/:originStruct/:originStructFieldName', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-detail/:id', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-presentation/:id', component: AnimatePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-presentation-special/:id', component: AnimatePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_goanimatepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-circles', component: CirclesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-adder', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-adder/:id/:originStruct/:originStructFieldName', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-detail/:id', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-presentation/:id', component: CirclePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-presentation-special/:id', component: CirclePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gocirclepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipses', component: EllipsesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-adder', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-adder/:id/:originStruct/:originStructFieldName', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-detail/:id', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-presentation/:id', component: EllipsePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-presentation-special/:id', component: EllipsePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_goellipsepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-lines', component: LinesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-adder', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-adder/:id/:originStruct/:originStructFieldName', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-detail/:id', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-presentation/:id', component: LinePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-presentation-special/:id', component: LinePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_golinepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-paths', component: PathsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-adder', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-adder/:id/:originStruct/:originStructFieldName', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-detail/:id', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-presentation/:id', component: PathPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-presentation-special/:id', component: PathPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gopathpres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-polygones', component: PolygonesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-adder', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-adder/:id/:originStruct/:originStructFieldName', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-detail/:id', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-presentation/:id', component: PolygonePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-presentation-special/:id', component: PolygonePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gopolygonepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-polylines', component: PolylinesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-adder', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-adder/:id/:originStruct/:originStructFieldName', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-detail/:id', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-presentation/:id', component: PolylinePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-presentation-special/:id', component: PolylinePresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gopolylinepres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-rects', component: RectsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-adder', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-adder/:id/:originStruct/:originStructFieldName', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-detail/:id', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-presentation/:id', component: RectPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-presentation-special/:id', component: RectPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gorectpres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-svgs', component: SVGsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder/:id/:originStruct/:originStructFieldName', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-detail/:id', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-presentation/:id', component: SVGPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-presentation-special/:id', component: SVGPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gosvgpres' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-texts', component: TextsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-adder', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-adder/:id/:originStruct/:originStructFieldName', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-detail/:id', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-presentation/:id', component: TextPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-presentation-special/:id', component: TextPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gotextpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
