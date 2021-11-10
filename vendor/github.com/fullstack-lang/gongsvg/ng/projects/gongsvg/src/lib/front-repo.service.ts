import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { AnimateDB } from './animate-db'
import { AnimateService } from './animate.service'

import { CircleDB } from './circle-db'
import { CircleService } from './circle.service'

import { EllipseDB } from './ellipse-db'
import { EllipseService } from './ellipse.service'

import { LineDB } from './line-db'
import { LineService } from './line.service'

import { PathDB } from './path-db'
import { PathService } from './path.service'

import { PolygoneDB } from './polygone-db'
import { PolygoneService } from './polygone.service'

import { PolylineDB } from './polyline-db'
import { PolylineService } from './polyline.service'

import { RectDB } from './rect-db'
import { RectService } from './rect.service'

import { SVGDB } from './svg-db'
import { SVGService } from './svg.service'

import { TextDB } from './text-db'
import { TextService } from './text.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Animates_array = new Array<AnimateDB>(); // array of repo instances
  Animates = new Map<number, AnimateDB>(); // map of repo instances
  Animates_batch = new Map<number, AnimateDB>(); // same but only in last GET (for finding repo instances to delete)
  Circles_array = new Array<CircleDB>(); // array of repo instances
  Circles = new Map<number, CircleDB>(); // map of repo instances
  Circles_batch = new Map<number, CircleDB>(); // same but only in last GET (for finding repo instances to delete)
  Ellipses_array = new Array<EllipseDB>(); // array of repo instances
  Ellipses = new Map<number, EllipseDB>(); // map of repo instances
  Ellipses_batch = new Map<number, EllipseDB>(); // same but only in last GET (for finding repo instances to delete)
  Lines_array = new Array<LineDB>(); // array of repo instances
  Lines = new Map<number, LineDB>(); // map of repo instances
  Lines_batch = new Map<number, LineDB>(); // same but only in last GET (for finding repo instances to delete)
  Paths_array = new Array<PathDB>(); // array of repo instances
  Paths = new Map<number, PathDB>(); // map of repo instances
  Paths_batch = new Map<number, PathDB>(); // same but only in last GET (for finding repo instances to delete)
  Polygones_array = new Array<PolygoneDB>(); // array of repo instances
  Polygones = new Map<number, PolygoneDB>(); // map of repo instances
  Polygones_batch = new Map<number, PolygoneDB>(); // same but only in last GET (for finding repo instances to delete)
  Polylines_array = new Array<PolylineDB>(); // array of repo instances
  Polylines = new Map<number, PolylineDB>(); // map of repo instances
  Polylines_batch = new Map<number, PolylineDB>(); // same but only in last GET (for finding repo instances to delete)
  Rects_array = new Array<RectDB>(); // array of repo instances
  Rects = new Map<number, RectDB>(); // map of repo instances
  Rects_batch = new Map<number, RectDB>(); // same but only in last GET (for finding repo instances to delete)
  SVGs_array = new Array<SVGDB>(); // array of repo instances
  SVGs = new Map<number, SVGDB>(); // map of repo instances
  SVGs_batch = new Map<number, SVGDB>(); // same but only in last GET (for finding repo instances to delete)
  Texts_array = new Array<TextDB>(); // array of repo instances
  Texts = new Map<number, TextDB>(); // map of repo instances
  Texts_batch = new Map<number, TextDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private animateService: AnimateService,
    private circleService: CircleService,
    private ellipseService: EllipseService,
    private lineService: LineService,
    private pathService: PathService,
    private polygoneService: PolygoneService,
    private polylineService: PolylineService,
    private rectService: RectService,
    private svgService: SVGService,
    private textService: TextService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<AnimateDB[]>,
    Observable<CircleDB[]>,
    Observable<EllipseDB[]>,
    Observable<LineDB[]>,
    Observable<PathDB[]>,
    Observable<PolygoneDB[]>,
    Observable<PolylineDB[]>,
    Observable<RectDB[]>,
    Observable<SVGDB[]>,
    Observable<TextDB[]>,
  ] = [ // insertion point sub template 
      this.animateService.getAnimates(),
      this.circleService.getCircles(),
      this.ellipseService.getEllipses(),
      this.lineService.getLines(),
      this.pathService.getPaths(),
      this.polygoneService.getPolygones(),
      this.polylineService.getPolylines(),
      this.rectService.getRects(),
      this.svgService.getSVGs(),
      this.textService.getTexts(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            animates_,
            circles_,
            ellipses_,
            lines_,
            paths_,
            polygones_,
            polylines_,
            rects_,
            svgs_,
            texts_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var animates: AnimateDB[]
            animates = animates_ as AnimateDB[]
            var circles: CircleDB[]
            circles = circles_ as CircleDB[]
            var ellipses: EllipseDB[]
            ellipses = ellipses_ as EllipseDB[]
            var lines: LineDB[]
            lines = lines_ as LineDB[]
            var paths: PathDB[]
            paths = paths_ as PathDB[]
            var polygones: PolygoneDB[]
            polygones = polygones_ as PolygoneDB[]
            var polylines: PolylineDB[]
            polylines = polylines_ as PolylineDB[]
            var rects: RectDB[]
            rects = rects_ as RectDB[]
            var svgs: SVGDB[]
            svgs = svgs_ as SVGDB[]
            var texts: TextDB[]
            texts = texts_ as TextDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Animates_array = animates

            // clear the map that counts Animate in the GET
            FrontRepoSingloton.Animates_batch.clear()

            animates.forEach(
              animate => {
                FrontRepoSingloton.Animates.set(animate.ID, animate)
                FrontRepoSingloton.Animates_batch.set(animate.ID, animate)
              }
            )

            // clear animates that are absent from the batch
            FrontRepoSingloton.Animates.forEach(
              animate => {
                if (FrontRepoSingloton.Animates_batch.get(animate.ID) == undefined) {
                  FrontRepoSingloton.Animates.delete(animate.ID)
                }
              }
            )

            // sort Animates_array array
            FrontRepoSingloton.Animates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Circles_array = circles

            // clear the map that counts Circle in the GET
            FrontRepoSingloton.Circles_batch.clear()

            circles.forEach(
              circle => {
                FrontRepoSingloton.Circles.set(circle.ID, circle)
                FrontRepoSingloton.Circles_batch.set(circle.ID, circle)
              }
            )

            // clear circles that are absent from the batch
            FrontRepoSingloton.Circles.forEach(
              circle => {
                if (FrontRepoSingloton.Circles_batch.get(circle.ID) == undefined) {
                  FrontRepoSingloton.Circles.delete(circle.ID)
                }
              }
            )

            // sort Circles_array array
            FrontRepoSingloton.Circles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            FrontRepoSingloton.Ellipses_batch.clear()

            ellipses.forEach(
              ellipse => {
                FrontRepoSingloton.Ellipses.set(ellipse.ID, ellipse)
                FrontRepoSingloton.Ellipses_batch.set(ellipse.ID, ellipse)
              }
            )

            // clear ellipses that are absent from the batch
            FrontRepoSingloton.Ellipses.forEach(
              ellipse => {
                if (FrontRepoSingloton.Ellipses_batch.get(ellipse.ID) == undefined) {
                  FrontRepoSingloton.Ellipses.delete(ellipse.ID)
                }
              }
            )

            // sort Ellipses_array array
            FrontRepoSingloton.Ellipses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Lines_array = lines

            // clear the map that counts Line in the GET
            FrontRepoSingloton.Lines_batch.clear()

            lines.forEach(
              line => {
                FrontRepoSingloton.Lines.set(line.ID, line)
                FrontRepoSingloton.Lines_batch.set(line.ID, line)
              }
            )

            // clear lines that are absent from the batch
            FrontRepoSingloton.Lines.forEach(
              line => {
                if (FrontRepoSingloton.Lines_batch.get(line.ID) == undefined) {
                  FrontRepoSingloton.Lines.delete(line.ID)
                }
              }
            )

            // sort Lines_array array
            FrontRepoSingloton.Lines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Paths_array = paths

            // clear the map that counts Path in the GET
            FrontRepoSingloton.Paths_batch.clear()

            paths.forEach(
              path => {
                FrontRepoSingloton.Paths.set(path.ID, path)
                FrontRepoSingloton.Paths_batch.set(path.ID, path)
              }
            )

            // clear paths that are absent from the batch
            FrontRepoSingloton.Paths.forEach(
              path => {
                if (FrontRepoSingloton.Paths_batch.get(path.ID) == undefined) {
                  FrontRepoSingloton.Paths.delete(path.ID)
                }
              }
            )

            // sort Paths_array array
            FrontRepoSingloton.Paths_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            FrontRepoSingloton.Polygones_batch.clear()

            polygones.forEach(
              polygone => {
                FrontRepoSingloton.Polygones.set(polygone.ID, polygone)
                FrontRepoSingloton.Polygones_batch.set(polygone.ID, polygone)
              }
            )

            // clear polygones that are absent from the batch
            FrontRepoSingloton.Polygones.forEach(
              polygone => {
                if (FrontRepoSingloton.Polygones_batch.get(polygone.ID) == undefined) {
                  FrontRepoSingloton.Polygones.delete(polygone.ID)
                }
              }
            )

            // sort Polygones_array array
            FrontRepoSingloton.Polygones_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            FrontRepoSingloton.Polylines_batch.clear()

            polylines.forEach(
              polyline => {
                FrontRepoSingloton.Polylines.set(polyline.ID, polyline)
                FrontRepoSingloton.Polylines_batch.set(polyline.ID, polyline)
              }
            )

            // clear polylines that are absent from the batch
            FrontRepoSingloton.Polylines.forEach(
              polyline => {
                if (FrontRepoSingloton.Polylines_batch.get(polyline.ID) == undefined) {
                  FrontRepoSingloton.Polylines.delete(polyline.ID)
                }
              }
            )

            // sort Polylines_array array
            FrontRepoSingloton.Polylines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Rects_array = rects

            // clear the map that counts Rect in the GET
            FrontRepoSingloton.Rects_batch.clear()

            rects.forEach(
              rect => {
                FrontRepoSingloton.Rects.set(rect.ID, rect)
                FrontRepoSingloton.Rects_batch.set(rect.ID, rect)
              }
            )

            // clear rects that are absent from the batch
            FrontRepoSingloton.Rects.forEach(
              rect => {
                if (FrontRepoSingloton.Rects_batch.get(rect.ID) == undefined) {
                  FrontRepoSingloton.Rects.delete(rect.ID)
                }
              }
            )

            // sort Rects_array array
            FrontRepoSingloton.Rects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            FrontRepoSingloton.SVGs_batch.clear()

            svgs.forEach(
              svg => {
                FrontRepoSingloton.SVGs.set(svg.ID, svg)
                FrontRepoSingloton.SVGs_batch.set(svg.ID, svg)
              }
            )

            // clear svgs that are absent from the batch
            FrontRepoSingloton.SVGs.forEach(
              svg => {
                if (FrontRepoSingloton.SVGs_batch.get(svg.ID) == undefined) {
                  FrontRepoSingloton.SVGs.delete(svg.ID)
                }
              }
            )

            // sort SVGs_array array
            FrontRepoSingloton.SVGs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Texts_array = texts

            // clear the map that counts Text in the GET
            FrontRepoSingloton.Texts_batch.clear()

            texts.forEach(
              text => {
                FrontRepoSingloton.Texts.set(text.ID, text)
                FrontRepoSingloton.Texts_batch.set(text.ID, text)
              }
            )

            // clear texts that are absent from the batch
            FrontRepoSingloton.Texts.forEach(
              text => {
                if (FrontRepoSingloton.Texts_batch.get(text.ID) == undefined) {
                  FrontRepoSingloton.Texts.delete(text.ID)
                }
              }
            )

            // sort Texts_array array
            FrontRepoSingloton.Texts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            animates.forEach(
              animate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Circle.Animations redeeming
                {
                  let _circle = FrontRepoSingloton.Circles.get(animate.Circle_AnimationsDBID.Int64)
                  if (_circle) {
                    if (_circle.Animations == undefined) {
                      _circle.Animations = new Array<AnimateDB>()
                    }
                    _circle.Animations.push(animate)
                    if (animate.Circle_Animations_reverse == undefined) {
                      animate.Circle_Animations_reverse = _circle
                    }
                  }
                }
                // insertion point for slice of pointer field Ellipse.Animates redeeming
                {
                  let _ellipse = FrontRepoSingloton.Ellipses.get(animate.Ellipse_AnimatesDBID.Int64)
                  if (_ellipse) {
                    if (_ellipse.Animates == undefined) {
                      _ellipse.Animates = new Array<AnimateDB>()
                    }
                    _ellipse.Animates.push(animate)
                    if (animate.Ellipse_Animates_reverse == undefined) {
                      animate.Ellipse_Animates_reverse = _ellipse
                    }
                  }
                }
                // insertion point for slice of pointer field Line.Animates redeeming
                {
                  let _line = FrontRepoSingloton.Lines.get(animate.Line_AnimatesDBID.Int64)
                  if (_line) {
                    if (_line.Animates == undefined) {
                      _line.Animates = new Array<AnimateDB>()
                    }
                    _line.Animates.push(animate)
                    if (animate.Line_Animates_reverse == undefined) {
                      animate.Line_Animates_reverse = _line
                    }
                  }
                }
                // insertion point for slice of pointer field Path.Animates redeeming
                {
                  let _path = FrontRepoSingloton.Paths.get(animate.Path_AnimatesDBID.Int64)
                  if (_path) {
                    if (_path.Animates == undefined) {
                      _path.Animates = new Array<AnimateDB>()
                    }
                    _path.Animates.push(animate)
                    if (animate.Path_Animates_reverse == undefined) {
                      animate.Path_Animates_reverse = _path
                    }
                  }
                }
                // insertion point for slice of pointer field Polygone.Animates redeeming
                {
                  let _polygone = FrontRepoSingloton.Polygones.get(animate.Polygone_AnimatesDBID.Int64)
                  if (_polygone) {
                    if (_polygone.Animates == undefined) {
                      _polygone.Animates = new Array<AnimateDB>()
                    }
                    _polygone.Animates.push(animate)
                    if (animate.Polygone_Animates_reverse == undefined) {
                      animate.Polygone_Animates_reverse = _polygone
                    }
                  }
                }
                // insertion point for slice of pointer field Polyline.Animates redeeming
                {
                  let _polyline = FrontRepoSingloton.Polylines.get(animate.Polyline_AnimatesDBID.Int64)
                  if (_polyline) {
                    if (_polyline.Animates == undefined) {
                      _polyline.Animates = new Array<AnimateDB>()
                    }
                    _polyline.Animates.push(animate)
                    if (animate.Polyline_Animates_reverse == undefined) {
                      animate.Polyline_Animates_reverse = _polyline
                    }
                  }
                }
                // insertion point for slice of pointer field Rect.Animations redeeming
                {
                  let _rect = FrontRepoSingloton.Rects.get(animate.Rect_AnimationsDBID.Int64)
                  if (_rect) {
                    if (_rect.Animations == undefined) {
                      _rect.Animations = new Array<AnimateDB>()
                    }
                    _rect.Animations.push(animate)
                    if (animate.Rect_Animations_reverse == undefined) {
                      animate.Rect_Animations_reverse = _rect
                    }
                  }
                }
                // insertion point for slice of pointer field Text.Animates redeeming
                {
                  let _text = FrontRepoSingloton.Texts.get(animate.Text_AnimatesDBID.Int64)
                  if (_text) {
                    if (_text.Animates == undefined) {
                      _text.Animates = new Array<AnimateDB>()
                    }
                    _text.Animates.push(animate)
                    if (animate.Text_Animates_reverse == undefined) {
                      animate.Text_Animates_reverse = _text
                    }
                  }
                }
              }
            )
            circles.forEach(
              circle => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Circles redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(circle.SVG_CirclesDBID.Int64)
                  if (_svg) {
                    if (_svg.Circles == undefined) {
                      _svg.Circles = new Array<CircleDB>()
                    }
                    _svg.Circles.push(circle)
                    if (circle.SVG_Circles_reverse == undefined) {
                      circle.SVG_Circles_reverse = _svg
                    }
                  }
                }
              }
            )
            ellipses.forEach(
              ellipse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Ellipses redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(ellipse.SVG_EllipsesDBID.Int64)
                  if (_svg) {
                    if (_svg.Ellipses == undefined) {
                      _svg.Ellipses = new Array<EllipseDB>()
                    }
                    _svg.Ellipses.push(ellipse)
                    if (ellipse.SVG_Ellipses_reverse == undefined) {
                      ellipse.SVG_Ellipses_reverse = _svg
                    }
                  }
                }
              }
            )
            lines.forEach(
              line => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Lines redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(line.SVG_LinesDBID.Int64)
                  if (_svg) {
                    if (_svg.Lines == undefined) {
                      _svg.Lines = new Array<LineDB>()
                    }
                    _svg.Lines.push(line)
                    if (line.SVG_Lines_reverse == undefined) {
                      line.SVG_Lines_reverse = _svg
                    }
                  }
                }
              }
            )
            paths.forEach(
              path => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Paths redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(path.SVG_PathsDBID.Int64)
                  if (_svg) {
                    if (_svg.Paths == undefined) {
                      _svg.Paths = new Array<PathDB>()
                    }
                    _svg.Paths.push(path)
                    if (path.SVG_Paths_reverse == undefined) {
                      path.SVG_Paths_reverse = _svg
                    }
                  }
                }
              }
            )
            polygones.forEach(
              polygone => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Polygones redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(polygone.SVG_PolygonesDBID.Int64)
                  if (_svg) {
                    if (_svg.Polygones == undefined) {
                      _svg.Polygones = new Array<PolygoneDB>()
                    }
                    _svg.Polygones.push(polygone)
                    if (polygone.SVG_Polygones_reverse == undefined) {
                      polygone.SVG_Polygones_reverse = _svg
                    }
                  }
                }
              }
            )
            polylines.forEach(
              polyline => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Polylines redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(polyline.SVG_PolylinesDBID.Int64)
                  if (_svg) {
                    if (_svg.Polylines == undefined) {
                      _svg.Polylines = new Array<PolylineDB>()
                    }
                    _svg.Polylines.push(polyline)
                    if (polyline.SVG_Polylines_reverse == undefined) {
                      polyline.SVG_Polylines_reverse = _svg
                    }
                  }
                }
              }
            )
            rects.forEach(
              rect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Rects redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(rect.SVG_RectsDBID.Int64)
                  if (_svg) {
                    if (_svg.Rects == undefined) {
                      _svg.Rects = new Array<RectDB>()
                    }
                    _svg.Rects.push(rect)
                    if (rect.SVG_Rects_reverse == undefined) {
                      rect.SVG_Rects_reverse = _svg
                    }
                  }
                }
              }
            )
            svgs.forEach(
              svg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            texts.forEach(
              text => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Texts redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(text.SVG_TextsDBID.Int64)
                  if (_svg) {
                    if (_svg.Texts == undefined) {
                      _svg.Texts = new Array<TextDB>()
                    }
                    _svg.Texts.push(text)
                    if (text.SVG_Texts_reverse == undefined) {
                      text.SVG_Texts_reverse = _svg
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // AnimatePull performs a GET on Animate of the stack and redeem association pointers 
  AnimatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.animateService.getAnimates()
        ]).subscribe(
          ([ // insertion point sub template 
            animates,
          ]) => {
            // init the array
            FrontRepoSingloton.Animates_array = animates

            // clear the map that counts Animate in the GET
            FrontRepoSingloton.Animates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            animates.forEach(
              animate => {
                FrontRepoSingloton.Animates.set(animate.ID, animate)
                FrontRepoSingloton.Animates_batch.set(animate.ID, animate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Circle.Animations redeeming
                {
                  let _circle = FrontRepoSingloton.Circles.get(animate.Circle_AnimationsDBID.Int64)
                  if (_circle) {
                    if (_circle.Animations == undefined) {
                      _circle.Animations = new Array<AnimateDB>()
                    }
                    _circle.Animations.push(animate)
                    if (animate.Circle_Animations_reverse == undefined) {
                      animate.Circle_Animations_reverse = _circle
                    }
                  }
                }
                // insertion point for slice of pointer field Ellipse.Animates redeeming
                {
                  let _ellipse = FrontRepoSingloton.Ellipses.get(animate.Ellipse_AnimatesDBID.Int64)
                  if (_ellipse) {
                    if (_ellipse.Animates == undefined) {
                      _ellipse.Animates = new Array<AnimateDB>()
                    }
                    _ellipse.Animates.push(animate)
                    if (animate.Ellipse_Animates_reverse == undefined) {
                      animate.Ellipse_Animates_reverse = _ellipse
                    }
                  }
                }
                // insertion point for slice of pointer field Line.Animates redeeming
                {
                  let _line = FrontRepoSingloton.Lines.get(animate.Line_AnimatesDBID.Int64)
                  if (_line) {
                    if (_line.Animates == undefined) {
                      _line.Animates = new Array<AnimateDB>()
                    }
                    _line.Animates.push(animate)
                    if (animate.Line_Animates_reverse == undefined) {
                      animate.Line_Animates_reverse = _line
                    }
                  }
                }
                // insertion point for slice of pointer field Path.Animates redeeming
                {
                  let _path = FrontRepoSingloton.Paths.get(animate.Path_AnimatesDBID.Int64)
                  if (_path) {
                    if (_path.Animates == undefined) {
                      _path.Animates = new Array<AnimateDB>()
                    }
                    _path.Animates.push(animate)
                    if (animate.Path_Animates_reverse == undefined) {
                      animate.Path_Animates_reverse = _path
                    }
                  }
                }
                // insertion point for slice of pointer field Polygone.Animates redeeming
                {
                  let _polygone = FrontRepoSingloton.Polygones.get(animate.Polygone_AnimatesDBID.Int64)
                  if (_polygone) {
                    if (_polygone.Animates == undefined) {
                      _polygone.Animates = new Array<AnimateDB>()
                    }
                    _polygone.Animates.push(animate)
                    if (animate.Polygone_Animates_reverse == undefined) {
                      animate.Polygone_Animates_reverse = _polygone
                    }
                  }
                }
                // insertion point for slice of pointer field Polyline.Animates redeeming
                {
                  let _polyline = FrontRepoSingloton.Polylines.get(animate.Polyline_AnimatesDBID.Int64)
                  if (_polyline) {
                    if (_polyline.Animates == undefined) {
                      _polyline.Animates = new Array<AnimateDB>()
                    }
                    _polyline.Animates.push(animate)
                    if (animate.Polyline_Animates_reverse == undefined) {
                      animate.Polyline_Animates_reverse = _polyline
                    }
                  }
                }
                // insertion point for slice of pointer field Rect.Animations redeeming
                {
                  let _rect = FrontRepoSingloton.Rects.get(animate.Rect_AnimationsDBID.Int64)
                  if (_rect) {
                    if (_rect.Animations == undefined) {
                      _rect.Animations = new Array<AnimateDB>()
                    }
                    _rect.Animations.push(animate)
                    if (animate.Rect_Animations_reverse == undefined) {
                      animate.Rect_Animations_reverse = _rect
                    }
                  }
                }
                // insertion point for slice of pointer field Text.Animates redeeming
                {
                  let _text = FrontRepoSingloton.Texts.get(animate.Text_AnimatesDBID.Int64)
                  if (_text) {
                    if (_text.Animates == undefined) {
                      _text.Animates = new Array<AnimateDB>()
                    }
                    _text.Animates.push(animate)
                    if (animate.Text_Animates_reverse == undefined) {
                      animate.Text_Animates_reverse = _text
                    }
                  }
                }
              }
            )

            // clear animates that are absent from the GET
            FrontRepoSingloton.Animates.forEach(
              animate => {
                if (FrontRepoSingloton.Animates_batch.get(animate.ID) == undefined) {
                  FrontRepoSingloton.Animates.delete(animate.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // CirclePull performs a GET on Circle of the stack and redeem association pointers 
  CirclePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.circleService.getCircles()
        ]).subscribe(
          ([ // insertion point sub template 
            circles,
          ]) => {
            // init the array
            FrontRepoSingloton.Circles_array = circles

            // clear the map that counts Circle in the GET
            FrontRepoSingloton.Circles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            circles.forEach(
              circle => {
                FrontRepoSingloton.Circles.set(circle.ID, circle)
                FrontRepoSingloton.Circles_batch.set(circle.ID, circle)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Circles redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(circle.SVG_CirclesDBID.Int64)
                  if (_svg) {
                    if (_svg.Circles == undefined) {
                      _svg.Circles = new Array<CircleDB>()
                    }
                    _svg.Circles.push(circle)
                    if (circle.SVG_Circles_reverse == undefined) {
                      circle.SVG_Circles_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear circles that are absent from the GET
            FrontRepoSingloton.Circles.forEach(
              circle => {
                if (FrontRepoSingloton.Circles_batch.get(circle.ID) == undefined) {
                  FrontRepoSingloton.Circles.delete(circle.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // EllipsePull performs a GET on Ellipse of the stack and redeem association pointers 
  EllipsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ellipseService.getEllipses()
        ]).subscribe(
          ([ // insertion point sub template 
            ellipses,
          ]) => {
            // init the array
            FrontRepoSingloton.Ellipses_array = ellipses

            // clear the map that counts Ellipse in the GET
            FrontRepoSingloton.Ellipses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ellipses.forEach(
              ellipse => {
                FrontRepoSingloton.Ellipses.set(ellipse.ID, ellipse)
                FrontRepoSingloton.Ellipses_batch.set(ellipse.ID, ellipse)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Ellipses redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(ellipse.SVG_EllipsesDBID.Int64)
                  if (_svg) {
                    if (_svg.Ellipses == undefined) {
                      _svg.Ellipses = new Array<EllipseDB>()
                    }
                    _svg.Ellipses.push(ellipse)
                    if (ellipse.SVG_Ellipses_reverse == undefined) {
                      ellipse.SVG_Ellipses_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear ellipses that are absent from the GET
            FrontRepoSingloton.Ellipses.forEach(
              ellipse => {
                if (FrontRepoSingloton.Ellipses_batch.get(ellipse.ID) == undefined) {
                  FrontRepoSingloton.Ellipses.delete(ellipse.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // LinePull performs a GET on Line of the stack and redeem association pointers 
  LinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.lineService.getLines()
        ]).subscribe(
          ([ // insertion point sub template 
            lines,
          ]) => {
            // init the array
            FrontRepoSingloton.Lines_array = lines

            // clear the map that counts Line in the GET
            FrontRepoSingloton.Lines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            lines.forEach(
              line => {
                FrontRepoSingloton.Lines.set(line.ID, line)
                FrontRepoSingloton.Lines_batch.set(line.ID, line)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Lines redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(line.SVG_LinesDBID.Int64)
                  if (_svg) {
                    if (_svg.Lines == undefined) {
                      _svg.Lines = new Array<LineDB>()
                    }
                    _svg.Lines.push(line)
                    if (line.SVG_Lines_reverse == undefined) {
                      line.SVG_Lines_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear lines that are absent from the GET
            FrontRepoSingloton.Lines.forEach(
              line => {
                if (FrontRepoSingloton.Lines_batch.get(line.ID) == undefined) {
                  FrontRepoSingloton.Lines.delete(line.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // PathPull performs a GET on Path of the stack and redeem association pointers 
  PathPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pathService.getPaths()
        ]).subscribe(
          ([ // insertion point sub template 
            paths,
          ]) => {
            // init the array
            FrontRepoSingloton.Paths_array = paths

            // clear the map that counts Path in the GET
            FrontRepoSingloton.Paths_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paths.forEach(
              path => {
                FrontRepoSingloton.Paths.set(path.ID, path)
                FrontRepoSingloton.Paths_batch.set(path.ID, path)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Paths redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(path.SVG_PathsDBID.Int64)
                  if (_svg) {
                    if (_svg.Paths == undefined) {
                      _svg.Paths = new Array<PathDB>()
                    }
                    _svg.Paths.push(path)
                    if (path.SVG_Paths_reverse == undefined) {
                      path.SVG_Paths_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear paths that are absent from the GET
            FrontRepoSingloton.Paths.forEach(
              path => {
                if (FrontRepoSingloton.Paths_batch.get(path.ID) == undefined) {
                  FrontRepoSingloton.Paths.delete(path.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // PolygonePull performs a GET on Polygone of the stack and redeem association pointers 
  PolygonePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polygoneService.getPolygones()
        ]).subscribe(
          ([ // insertion point sub template 
            polygones,
          ]) => {
            // init the array
            FrontRepoSingloton.Polygones_array = polygones

            // clear the map that counts Polygone in the GET
            FrontRepoSingloton.Polygones_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polygones.forEach(
              polygone => {
                FrontRepoSingloton.Polygones.set(polygone.ID, polygone)
                FrontRepoSingloton.Polygones_batch.set(polygone.ID, polygone)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Polygones redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(polygone.SVG_PolygonesDBID.Int64)
                  if (_svg) {
                    if (_svg.Polygones == undefined) {
                      _svg.Polygones = new Array<PolygoneDB>()
                    }
                    _svg.Polygones.push(polygone)
                    if (polygone.SVG_Polygones_reverse == undefined) {
                      polygone.SVG_Polygones_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear polygones that are absent from the GET
            FrontRepoSingloton.Polygones.forEach(
              polygone => {
                if (FrontRepoSingloton.Polygones_batch.get(polygone.ID) == undefined) {
                  FrontRepoSingloton.Polygones.delete(polygone.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // PolylinePull performs a GET on Polyline of the stack and redeem association pointers 
  PolylinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.polylineService.getPolylines()
        ]).subscribe(
          ([ // insertion point sub template 
            polylines,
          ]) => {
            // init the array
            FrontRepoSingloton.Polylines_array = polylines

            // clear the map that counts Polyline in the GET
            FrontRepoSingloton.Polylines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            polylines.forEach(
              polyline => {
                FrontRepoSingloton.Polylines.set(polyline.ID, polyline)
                FrontRepoSingloton.Polylines_batch.set(polyline.ID, polyline)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Polylines redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(polyline.SVG_PolylinesDBID.Int64)
                  if (_svg) {
                    if (_svg.Polylines == undefined) {
                      _svg.Polylines = new Array<PolylineDB>()
                    }
                    _svg.Polylines.push(polyline)
                    if (polyline.SVG_Polylines_reverse == undefined) {
                      polyline.SVG_Polylines_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear polylines that are absent from the GET
            FrontRepoSingloton.Polylines.forEach(
              polyline => {
                if (FrontRepoSingloton.Polylines_batch.get(polyline.ID) == undefined) {
                  FrontRepoSingloton.Polylines.delete(polyline.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // RectPull performs a GET on Rect of the stack and redeem association pointers 
  RectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectService.getRects()
        ]).subscribe(
          ([ // insertion point sub template 
            rects,
          ]) => {
            // init the array
            FrontRepoSingloton.Rects_array = rects

            // clear the map that counts Rect in the GET
            FrontRepoSingloton.Rects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rects.forEach(
              rect => {
                FrontRepoSingloton.Rects.set(rect.ID, rect)
                FrontRepoSingloton.Rects_batch.set(rect.ID, rect)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Rects redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(rect.SVG_RectsDBID.Int64)
                  if (_svg) {
                    if (_svg.Rects == undefined) {
                      _svg.Rects = new Array<RectDB>()
                    }
                    _svg.Rects.push(rect)
                    if (rect.SVG_Rects_reverse == undefined) {
                      rect.SVG_Rects_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear rects that are absent from the GET
            FrontRepoSingloton.Rects.forEach(
              rect => {
                if (FrontRepoSingloton.Rects_batch.get(rect.ID) == undefined) {
                  FrontRepoSingloton.Rects.delete(rect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // SVGPull performs a GET on SVG of the stack and redeem association pointers 
  SVGPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.svgService.getSVGs()
        ]).subscribe(
          ([ // insertion point sub template 
            svgs,
          ]) => {
            // init the array
            FrontRepoSingloton.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            FrontRepoSingloton.SVGs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            svgs.forEach(
              svg => {
                FrontRepoSingloton.SVGs.set(svg.ID, svg)
                FrontRepoSingloton.SVGs_batch.set(svg.ID, svg)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear svgs that are absent from the GET
            FrontRepoSingloton.SVGs.forEach(
              svg => {
                if (FrontRepoSingloton.SVGs_batch.get(svg.ID) == undefined) {
                  FrontRepoSingloton.SVGs.delete(svg.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // TextPull performs a GET on Text of the stack and redeem association pointers 
  TextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.textService.getTexts()
        ]).subscribe(
          ([ // insertion point sub template 
            texts,
          ]) => {
            // init the array
            FrontRepoSingloton.Texts_array = texts

            // clear the map that counts Text in the GET
            FrontRepoSingloton.Texts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            texts.forEach(
              text => {
                FrontRepoSingloton.Texts.set(text.ID, text)
                FrontRepoSingloton.Texts_batch.set(text.ID, text)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Texts redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(text.SVG_TextsDBID.Int64)
                  if (_svg) {
                    if (_svg.Texts == undefined) {
                      _svg.Texts = new Array<TextDB>()
                    }
                    _svg.Texts.push(text)
                    if (text.SVG_Texts_reverse == undefined) {
                      text.SVG_Texts_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear texts that are absent from the GET
            FrontRepoSingloton.Texts.forEach(
              text => {
                if (FrontRepoSingloton.Texts_batch.get(text.ID) == undefined) {
                  FrontRepoSingloton.Texts.delete(text.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getAnimateUniqueID(id: number): number {
  return 31 * id
}
export function getCircleUniqueID(id: number): number {
  return 37 * id
}
export function getEllipseUniqueID(id: number): number {
  return 41 * id
}
export function getLineUniqueID(id: number): number {
  return 43 * id
}
export function getPathUniqueID(id: number): number {
  return 47 * id
}
export function getPolygoneUniqueID(id: number): number {
  return 53 * id
}
export function getPolylineUniqueID(id: number): number {
  return 59 * id
}
export function getRectUniqueID(id: number): number {
  return 61 * id
}
export function getSVGUniqueID(id: number): number {
  return 67 * id
}
export function getTextUniqueID(id: number): number {
  return 71 * id
}
