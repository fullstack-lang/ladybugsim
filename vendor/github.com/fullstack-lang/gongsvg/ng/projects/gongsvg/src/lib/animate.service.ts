// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { AnimateDB } from './animate-db';

// insertion point for imports
import { CircleDB } from './circle-db'
import { EllipseDB } from './ellipse-db'
import { LineDB } from './line-db'
import { PathDB } from './path-db'
import { PolygoneDB } from './polygone-db'
import { PolylineDB } from './polyline-db'
import { RectDB } from './rect-db'
import { TextDB } from './text-db'

@Injectable({
  providedIn: 'root'
})
export class AnimateService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  AnimateServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private animatesUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.animatesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/animates';
  }

  /** GET animates from the server */
  getAnimates(): Observable<AnimateDB[]> {
    return this.http.get<AnimateDB[]>(this.animatesUrl)
      .pipe(
        tap(_ => this.log('fetched animates')),
        catchError(this.handleError<AnimateDB[]>('getAnimates', []))
      );
  }

  /** GET animate by id. Will 404 if id not found */
  getAnimate(id: number): Observable<AnimateDB> {
    const url = `${this.animatesUrl}/${id}`;
    return this.http.get<AnimateDB>(url).pipe(
      tap(_ => this.log(`fetched animate id=${id}`)),
      catchError(this.handleError<AnimateDB>(`getAnimate id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new animate to the server */
  postAnimate(animatedb: AnimateDB): Observable<AnimateDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Circle_Animations_reverse = animatedb.Circle_Animations_reverse
    animatedb.Circle_Animations_reverse = new CircleDB
    let _Ellipse_Animates_reverse = animatedb.Ellipse_Animates_reverse
    animatedb.Ellipse_Animates_reverse = new EllipseDB
    let _Line_Animates_reverse = animatedb.Line_Animates_reverse
    animatedb.Line_Animates_reverse = new LineDB
    let _Path_Animates_reverse = animatedb.Path_Animates_reverse
    animatedb.Path_Animates_reverse = new PathDB
    let _Polygone_Animates_reverse = animatedb.Polygone_Animates_reverse
    animatedb.Polygone_Animates_reverse = new PolygoneDB
    let _Polyline_Animates_reverse = animatedb.Polyline_Animates_reverse
    animatedb.Polyline_Animates_reverse = new PolylineDB
    let _Rect_Animations_reverse = animatedb.Rect_Animations_reverse
    animatedb.Rect_Animations_reverse = new RectDB
    let _Text_Animates_reverse = animatedb.Text_Animates_reverse
    animatedb.Text_Animates_reverse = new TextDB

    return this.http.post<AnimateDB>(this.animatesUrl, animatedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        animatedb.Circle_Animations_reverse = _Circle_Animations_reverse
        animatedb.Ellipse_Animates_reverse = _Ellipse_Animates_reverse
        animatedb.Line_Animates_reverse = _Line_Animates_reverse
        animatedb.Path_Animates_reverse = _Path_Animates_reverse
        animatedb.Polygone_Animates_reverse = _Polygone_Animates_reverse
        animatedb.Polyline_Animates_reverse = _Polyline_Animates_reverse
        animatedb.Rect_Animations_reverse = _Rect_Animations_reverse
        animatedb.Text_Animates_reverse = _Text_Animates_reverse
        this.log(`posted animatedb id=${animatedb.ID}`)
      }),
      catchError(this.handleError<AnimateDB>('postAnimate'))
    );
  }

  /** DELETE: delete the animatedb from the server */
  deleteAnimate(animatedb: AnimateDB | number): Observable<AnimateDB> {
    const id = typeof animatedb === 'number' ? animatedb : animatedb.ID;
    const url = `${this.animatesUrl}/${id}`;

    return this.http.delete<AnimateDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted animatedb id=${id}`)),
      catchError(this.handleError<AnimateDB>('deleteAnimate'))
    );
  }

  /** PUT: update the animatedb on the server */
  updateAnimate(animatedb: AnimateDB): Observable<AnimateDB> {
    const id = typeof animatedb === 'number' ? animatedb : animatedb.ID;
    const url = `${this.animatesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Circle_Animations_reverse = animatedb.Circle_Animations_reverse
    animatedb.Circle_Animations_reverse = new CircleDB
    let _Ellipse_Animates_reverse = animatedb.Ellipse_Animates_reverse
    animatedb.Ellipse_Animates_reverse = new EllipseDB
    let _Line_Animates_reverse = animatedb.Line_Animates_reverse
    animatedb.Line_Animates_reverse = new LineDB
    let _Path_Animates_reverse = animatedb.Path_Animates_reverse
    animatedb.Path_Animates_reverse = new PathDB
    let _Polygone_Animates_reverse = animatedb.Polygone_Animates_reverse
    animatedb.Polygone_Animates_reverse = new PolygoneDB
    let _Polyline_Animates_reverse = animatedb.Polyline_Animates_reverse
    animatedb.Polyline_Animates_reverse = new PolylineDB
    let _Rect_Animations_reverse = animatedb.Rect_Animations_reverse
    animatedb.Rect_Animations_reverse = new RectDB
    let _Text_Animates_reverse = animatedb.Text_Animates_reverse
    animatedb.Text_Animates_reverse = new TextDB

    return this.http.put<AnimateDB>(url, animatedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        animatedb.Circle_Animations_reverse = _Circle_Animations_reverse
        animatedb.Ellipse_Animates_reverse = _Ellipse_Animates_reverse
        animatedb.Line_Animates_reverse = _Line_Animates_reverse
        animatedb.Path_Animates_reverse = _Path_Animates_reverse
        animatedb.Polygone_Animates_reverse = _Polygone_Animates_reverse
        animatedb.Polyline_Animates_reverse = _Polyline_Animates_reverse
        animatedb.Rect_Animations_reverse = _Rect_Animations_reverse
        animatedb.Text_Animates_reverse = _Text_Animates_reverse
        this.log(`updated animatedb id=${animatedb.ID}`)
      }),
      catchError(this.handleError<AnimateDB>('updateAnimate'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
