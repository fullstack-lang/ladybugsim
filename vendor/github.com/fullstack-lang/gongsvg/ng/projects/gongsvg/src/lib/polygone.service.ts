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

import { PolygoneDB } from './polygone-db';

// insertion point for imports
import { SVGDB } from './svg-db'

@Injectable({
  providedIn: 'root'
})
export class PolygoneService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  PolygoneServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private polygonesUrl: string

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
    this.polygonesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/polygones';
  }

  /** GET polygones from the server */
  getPolygones(): Observable<PolygoneDB[]> {
    return this.http.get<PolygoneDB[]>(this.polygonesUrl)
      .pipe(
        tap(_ => this.log('fetched polygones')),
        catchError(this.handleError<PolygoneDB[]>('getPolygones', []))
      );
  }

  /** GET polygone by id. Will 404 if id not found */
  getPolygone(id: number): Observable<PolygoneDB> {
    const url = `${this.polygonesUrl}/${id}`;
    return this.http.get<PolygoneDB>(url).pipe(
      tap(_ => this.log(`fetched polygone id=${id}`)),
      catchError(this.handleError<PolygoneDB>(`getPolygone id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new polygone to the server */
  postPolygone(polygonedb: PolygoneDB): Observable<PolygoneDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    polygonedb.Animates = []
    let _SVG_Polygones_reverse = polygonedb.SVG_Polygones_reverse
    polygonedb.SVG_Polygones_reverse = new SVGDB

    return this.http.post<PolygoneDB>(this.polygonesUrl, polygonedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        polygonedb.SVG_Polygones_reverse = _SVG_Polygones_reverse
        this.log(`posted polygonedb id=${polygonedb.ID}`)
      }),
      catchError(this.handleError<PolygoneDB>('postPolygone'))
    );
  }

  /** DELETE: delete the polygonedb from the server */
  deletePolygone(polygonedb: PolygoneDB | number): Observable<PolygoneDB> {
    const id = typeof polygonedb === 'number' ? polygonedb : polygonedb.ID;
    const url = `${this.polygonesUrl}/${id}`;

    return this.http.delete<PolygoneDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted polygonedb id=${id}`)),
      catchError(this.handleError<PolygoneDB>('deletePolygone'))
    );
  }

  /** PUT: update the polygonedb on the server */
  updatePolygone(polygonedb: PolygoneDB): Observable<PolygoneDB> {
    const id = typeof polygonedb === 'number' ? polygonedb : polygonedb.ID;
    const url = `${this.polygonesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    polygonedb.Animates = []
    let _SVG_Polygones_reverse = polygonedb.SVG_Polygones_reverse
    polygonedb.SVG_Polygones_reverse = new SVGDB

    return this.http.put<PolygoneDB>(url, polygonedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        polygonedb.SVG_Polygones_reverse = _SVG_Polygones_reverse
        this.log(`updated polygonedb id=${polygonedb.ID}`)
      }),
      catchError(this.handleError<PolygoneDB>('updatePolygone'))
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