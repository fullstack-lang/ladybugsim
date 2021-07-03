// generated by MultiCodeGeneratorNgService
import { Injectable } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { LadybugSimulationAPI } from './ladybugsimulation-api';
import { LadybugSimulationDB } from './ladybugsimulation-db';

@Injectable({
  providedIn: 'root'
})
export class LadybugSimulationService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LadybugSimulationServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private ladybugsimulationsUrl = 'http://localhost:8080/api/github.com/fullstack-lang/ladybugsim/go/v1/ladybugsimulations';

  constructor(
    private http: HttpClient
  ) { }

  /** GET ladybugsimulations from the server */
  getLadybugSimulations(): Observable<LadybugSimulationDB[]> {
    return this.http.get<LadybugSimulationDB[]>(this.ladybugsimulationsUrl)
      .pipe(
        tap(_ => this.log('fetched ladybugsimulations')),
        catchError(this.handleError<LadybugSimulationDB[]>('getLadybugSimulations', []))
      );
  }

  /** GET ladybugsimulation by id. Will 404 if id not found */
  getLadybugSimulation(id: number): Observable<LadybugSimulationDB> {
    const url = `${this.ladybugsimulationsUrl}/${id}`;
    return this.http.get<LadybugSimulationDB>(url).pipe(
      tap(_ => this.log(`fetched ladybugsimulation id=${id}`)),
      catchError(this.handleError<LadybugSimulationDB>(`getLadybugSimulation id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new ladybugsimulation to the server */
  postLadybugSimulation(ladybugsimulationdb: LadybugSimulationDB): Observable<LadybugSimulationDB> {

		// insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    ladybugsimulationdb.Ladybugs = []

		return this.http.post<LadybugSimulationDB>(this.ladybugsimulationsUrl, ladybugsimulationdb, this.httpOptions).pipe(
			tap(_ => {
				// insertion point for restoration of reverse pointers
				this.log(`posted ladybugsimulationdb id=${ladybugsimulationdb.ID}`)
			}),
			catchError(this.handleError<LadybugSimulationDB>('postLadybugSimulation'))
		);
  }

  /** DELETE: delete the ladybugsimulationdb from the server */
  deleteLadybugSimulation(ladybugsimulationdb: LadybugSimulationDB | number): Observable<LadybugSimulationDB> {
    const id = typeof ladybugsimulationdb === 'number' ? ladybugsimulationdb : ladybugsimulationdb.ID;
    const url = `${this.ladybugsimulationsUrl}/${id}`;

    return this.http.delete<LadybugSimulationDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted ladybugsimulationdb id=${id}`)),
      catchError(this.handleError<LadybugSimulationDB>('deleteLadybugSimulation'))
    );
  }

  /** PUT: update the ladybugsimulationdb on the server */
  updateLadybugSimulation(ladybugsimulationdb: LadybugSimulationDB): Observable<LadybugSimulationDB> {
    const id = typeof ladybugsimulationdb === 'number' ? ladybugsimulationdb : ladybugsimulationdb.ID;
    const url = `${this.ladybugsimulationsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    ladybugsimulationdb.Ladybugs = []

    return this.http.put(url, ladybugsimulationdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated ladybugsimulationdb id=${ladybugsimulationdb.ID}`)
      }),
      catchError(this.handleError<LadybugSimulationDB>('updateLadybugSimulation'))
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