import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { SvgComponent } from './svg/svg.component';
import { RectComponent } from './rect/rect.component';
import { BrowserModule } from '@angular/platform-browser';
import { TextComponent } from './text/text.component';
import { CircleComponent } from './circle/circle.component';
import { LineComponent } from './line/line.component';
import { EllipseComponent } from './ellipse/ellipse.component';
import { PolylineComponent } from './polyline/polyline.component';
import { PathComponent } from './path/path.component';
import { PolygoneComponent } from './polygone/polygone.component'



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    SvgComponent,
    RectComponent,
    TextComponent,
    CircleComponent,
    LineComponent,
    EllipseComponent,
    PolylineComponent,
    PathComponent,
    PolygoneComponent
  ],
  imports: [
    BrowserModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent
  ]
})
export class GongsvgspecificModule { }
