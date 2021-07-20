import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { LadybugDB } from './ladybug-db'
import { LadybugService } from './ladybug.service'

import { LadybugSimulationDB } from './ladybugsimulation-db'
import { LadybugSimulationService } from './ladybugsimulation.service'

import { UpdatePositionEventDB } from './updatepositionevent-db'
import { UpdatePositionEventService } from './updatepositionevent.service'

import { UpdateSpeedEventDB } from './updatespeedevent-db'
import { UpdateSpeedEventService } from './updatespeedevent.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Ladybugs_array = new Array<LadybugDB>(); // array of repo instances
  Ladybugs = new Map<number, LadybugDB>(); // map of repo instances
  Ladybugs_batch = new Map<number, LadybugDB>(); // same but only in last GET (for finding repo instances to delete)
  LadybugSimulations_array = new Array<LadybugSimulationDB>(); // array of repo instances
  LadybugSimulations = new Map<number, LadybugSimulationDB>(); // map of repo instances
  LadybugSimulations_batch = new Map<number, LadybugSimulationDB>(); // same but only in last GET (for finding repo instances to delete)
  UpdatePositionEvents_array = new Array<UpdatePositionEventDB>(); // array of repo instances
  UpdatePositionEvents = new Map<number, UpdatePositionEventDB>(); // map of repo instances
  UpdatePositionEvents_batch = new Map<number, UpdatePositionEventDB>(); // same but only in last GET (for finding repo instances to delete)
  UpdateSpeedEvents_array = new Array<UpdateSpeedEventDB>(); // array of repo instances
  UpdateSpeedEvents = new Map<number, UpdateSpeedEventDB>(); // map of repo instances
  UpdateSpeedEvents_batch = new Map<number, UpdateSpeedEventDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
  Int64: number
  Valid: boolean
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number; // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean; // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode;

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string;  // The "Aclass"
  SourceField: string; // the "AnarrayofbUse"
  IntermediateStruct: string; // the "AclassBclassUse" 
  IntermediateStructField: string; // the "Bclass" as field
  NextAssociationStruct: string; // the "Bclass"
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
    private ladybugService: LadybugService,
    private ladybugsimulationService: LadybugSimulationService,
    private updatepositioneventService: UpdatePositionEventService,
    private updatespeedeventService: UpdateSpeedEventService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["post" + structName](instanceToBePosted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["delete" + structName](instanceToBeDeleted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<LadybugDB[]>,
    Observable<LadybugSimulationDB[]>,
    Observable<UpdatePositionEventDB[]>,
    Observable<UpdateSpeedEventDB[]>,
  ] = [ // insertion point sub template 
      this.ladybugService.getLadybugs(),
      this.ladybugsimulationService.getLadybugSimulations(),
      this.updatepositioneventService.getUpdatePositionEvents(),
      this.updatespeedeventService.getUpdateSpeedEvents(),
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
            ladybugs_,
            ladybugsimulations_,
            updatepositionevents_,
            updatespeedevents_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var ladybugs: LadybugDB[]
            ladybugs = ladybugs_
            var ladybugsimulations: LadybugSimulationDB[]
            ladybugsimulations = ladybugsimulations_
            var updatepositionevents: UpdatePositionEventDB[]
            updatepositionevents = updatepositionevents_
            var updatespeedevents: UpdateSpeedEventDB[]
            updatespeedevents = updatespeedevents_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Ladybugs_array = ladybugs

            // clear the map that counts Ladybug in the GET
            FrontRepoSingloton.Ladybugs_batch.clear()

            ladybugs.forEach(
              ladybug => {
                FrontRepoSingloton.Ladybugs.set(ladybug.ID, ladybug)
                FrontRepoSingloton.Ladybugs_batch.set(ladybug.ID, ladybug)
              }
            )

            // clear ladybugs that are absent from the batch
            FrontRepoSingloton.Ladybugs.forEach(
              ladybug => {
                if (FrontRepoSingloton.Ladybugs_batch.get(ladybug.ID) == undefined) {
                  FrontRepoSingloton.Ladybugs.delete(ladybug.ID)
                }
              }
            )

            // sort Ladybugs_array array
            FrontRepoSingloton.Ladybugs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.LadybugSimulations_array = ladybugsimulations

            // clear the map that counts LadybugSimulation in the GET
            FrontRepoSingloton.LadybugSimulations_batch.clear()

            ladybugsimulations.forEach(
              ladybugsimulation => {
                FrontRepoSingloton.LadybugSimulations.set(ladybugsimulation.ID, ladybugsimulation)
                FrontRepoSingloton.LadybugSimulations_batch.set(ladybugsimulation.ID, ladybugsimulation)
              }
            )

            // clear ladybugsimulations that are absent from the batch
            FrontRepoSingloton.LadybugSimulations.forEach(
              ladybugsimulation => {
                if (FrontRepoSingloton.LadybugSimulations_batch.get(ladybugsimulation.ID) == undefined) {
                  FrontRepoSingloton.LadybugSimulations.delete(ladybugsimulation.ID)
                }
              }
            )

            // sort LadybugSimulations_array array
            FrontRepoSingloton.LadybugSimulations_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.UpdatePositionEvents_array = updatepositionevents

            // clear the map that counts UpdatePositionEvent in the GET
            FrontRepoSingloton.UpdatePositionEvents_batch.clear()

            updatepositionevents.forEach(
              updatepositionevent => {
                FrontRepoSingloton.UpdatePositionEvents.set(updatepositionevent.ID, updatepositionevent)
                FrontRepoSingloton.UpdatePositionEvents_batch.set(updatepositionevent.ID, updatepositionevent)
              }
            )

            // clear updatepositionevents that are absent from the batch
            FrontRepoSingloton.UpdatePositionEvents.forEach(
              updatepositionevent => {
                if (FrontRepoSingloton.UpdatePositionEvents_batch.get(updatepositionevent.ID) == undefined) {
                  FrontRepoSingloton.UpdatePositionEvents.delete(updatepositionevent.ID)
                }
              }
            )

            // sort UpdatePositionEvents_array array
            FrontRepoSingloton.UpdatePositionEvents_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.UpdateSpeedEvents_array = updatespeedevents

            // clear the map that counts UpdateSpeedEvent in the GET
            FrontRepoSingloton.UpdateSpeedEvents_batch.clear()

            updatespeedevents.forEach(
              updatespeedevent => {
                FrontRepoSingloton.UpdateSpeedEvents.set(updatespeedevent.ID, updatespeedevent)
                FrontRepoSingloton.UpdateSpeedEvents_batch.set(updatespeedevent.ID, updatespeedevent)
              }
            )

            // clear updatespeedevents that are absent from the batch
            FrontRepoSingloton.UpdateSpeedEvents.forEach(
              updatespeedevent => {
                if (FrontRepoSingloton.UpdateSpeedEvents_batch.get(updatespeedevent.ID) == undefined) {
                  FrontRepoSingloton.UpdateSpeedEvents.delete(updatespeedevent.ID)
                }
              }
            )

            // sort UpdateSpeedEvents_array array
            FrontRepoSingloton.UpdateSpeedEvents_array.sort((t1, t2) => {
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
            ladybugs.forEach(
              ladybug => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field LadybugSimulation.Ladybugs redeeming
                {
                  let _ladybugsimulation = FrontRepoSingloton.LadybugSimulations.get(ladybug.LadybugSimulation_LadybugsDBID.Int64)
                  if (_ladybugsimulation) {
                    if (_ladybugsimulation.Ladybugs == undefined) {
                      _ladybugsimulation.Ladybugs = new Array<LadybugDB>()
                    }
                    _ladybugsimulation.Ladybugs.push(ladybug)
                    if (ladybug.LadybugSimulation_Ladybugs_reverse == undefined) {
                      ladybug.LadybugSimulation_Ladybugs_reverse = _ladybugsimulation
                    }
                  }
                }
              }
            )
            ladybugsimulations.forEach(
              ladybugsimulation => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            updatepositionevents.forEach(
              updatepositionevent => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            updatespeedevents.forEach(
              updatespeedevent => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
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

  // LadybugPull performs a GET on Ladybug of the stack and redeem association pointers 
  LadybugPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ladybugService.getLadybugs()
        ]).subscribe(
          ([ // insertion point sub template 
            ladybugs,
          ]) => {
            // init the array
            FrontRepoSingloton.Ladybugs_array = ladybugs

            // clear the map that counts Ladybug in the GET
            FrontRepoSingloton.Ladybugs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ladybugs.forEach(
              ladybug => {
                FrontRepoSingloton.Ladybugs.set(ladybug.ID, ladybug)
                FrontRepoSingloton.Ladybugs_batch.set(ladybug.ID, ladybug)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field LadybugSimulation.Ladybugs redeeming
                {
                  let _ladybugsimulation = FrontRepoSingloton.LadybugSimulations.get(ladybug.LadybugSimulation_LadybugsDBID.Int64)
                  if (_ladybugsimulation) {
                    if (_ladybugsimulation.Ladybugs == undefined) {
                      _ladybugsimulation.Ladybugs = new Array<LadybugDB>()
                    }
                    _ladybugsimulation.Ladybugs.push(ladybug)
                    if (ladybug.LadybugSimulation_Ladybugs_reverse == undefined) {
                      ladybug.LadybugSimulation_Ladybugs_reverse = _ladybugsimulation
                    }
                  }
                }
              }
            )

            // clear ladybugs that are absent from the GET
            FrontRepoSingloton.Ladybugs.forEach(
              ladybug => {
                if (FrontRepoSingloton.Ladybugs_batch.get(ladybug.ID) == undefined) {
                  FrontRepoSingloton.Ladybugs.delete(ladybug.ID)
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

  // LadybugSimulationPull performs a GET on LadybugSimulation of the stack and redeem association pointers 
  LadybugSimulationPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ladybugsimulationService.getLadybugSimulations()
        ]).subscribe(
          ([ // insertion point sub template 
            ladybugsimulations,
          ]) => {
            // init the array
            FrontRepoSingloton.LadybugSimulations_array = ladybugsimulations

            // clear the map that counts LadybugSimulation in the GET
            FrontRepoSingloton.LadybugSimulations_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ladybugsimulations.forEach(
              ladybugsimulation => {
                FrontRepoSingloton.LadybugSimulations.set(ladybugsimulation.ID, ladybugsimulation)
                FrontRepoSingloton.LadybugSimulations_batch.set(ladybugsimulation.ID, ladybugsimulation)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear ladybugsimulations that are absent from the GET
            FrontRepoSingloton.LadybugSimulations.forEach(
              ladybugsimulation => {
                if (FrontRepoSingloton.LadybugSimulations_batch.get(ladybugsimulation.ID) == undefined) {
                  FrontRepoSingloton.LadybugSimulations.delete(ladybugsimulation.ID)
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

  // UpdatePositionEventPull performs a GET on UpdatePositionEvent of the stack and redeem association pointers 
  UpdatePositionEventPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.updatepositioneventService.getUpdatePositionEvents()
        ]).subscribe(
          ([ // insertion point sub template 
            updatepositionevents,
          ]) => {
            // init the array
            FrontRepoSingloton.UpdatePositionEvents_array = updatepositionevents

            // clear the map that counts UpdatePositionEvent in the GET
            FrontRepoSingloton.UpdatePositionEvents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            updatepositionevents.forEach(
              updatepositionevent => {
                FrontRepoSingloton.UpdatePositionEvents.set(updatepositionevent.ID, updatepositionevent)
                FrontRepoSingloton.UpdatePositionEvents_batch.set(updatepositionevent.ID, updatepositionevent)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear updatepositionevents that are absent from the GET
            FrontRepoSingloton.UpdatePositionEvents.forEach(
              updatepositionevent => {
                if (FrontRepoSingloton.UpdatePositionEvents_batch.get(updatepositionevent.ID) == undefined) {
                  FrontRepoSingloton.UpdatePositionEvents.delete(updatepositionevent.ID)
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

  // UpdateSpeedEventPull performs a GET on UpdateSpeedEvent of the stack and redeem association pointers 
  UpdateSpeedEventPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.updatespeedeventService.getUpdateSpeedEvents()
        ]).subscribe(
          ([ // insertion point sub template 
            updatespeedevents,
          ]) => {
            // init the array
            FrontRepoSingloton.UpdateSpeedEvents_array = updatespeedevents

            // clear the map that counts UpdateSpeedEvent in the GET
            FrontRepoSingloton.UpdateSpeedEvents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            updatespeedevents.forEach(
              updatespeedevent => {
                FrontRepoSingloton.UpdateSpeedEvents.set(updatespeedevent.ID, updatespeedevent)
                FrontRepoSingloton.UpdateSpeedEvents_batch.set(updatespeedevent.ID, updatespeedevent)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear updatespeedevents that are absent from the GET
            FrontRepoSingloton.UpdateSpeedEvents.forEach(
              updatespeedevent => {
                if (FrontRepoSingloton.UpdateSpeedEvents_batch.get(updatespeedevent.ID) == undefined) {
                  FrontRepoSingloton.UpdateSpeedEvents.delete(updatespeedevent.ID)
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
export function getLadybugUniqueID(id: number): number {
  return 31 * id
}
export function getLadybugSimulationUniqueID(id: number): number {
  return 37 * id
}
export function getUpdatePositionEventUniqueID(id: number): number {
  return 41 * id
}
export function getUpdateSpeedEventUniqueID(id: number): number {
  return 43 * id
}
