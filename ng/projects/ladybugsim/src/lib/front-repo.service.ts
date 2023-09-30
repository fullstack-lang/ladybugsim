import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { LadybugDB } from './ladybug-db'
import { LadybugService } from './ladybug.service'

import { LadybugSimulationDB } from './ladybugsimulation-db'
import { LadybugSimulationService } from './ladybugsimulation.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Ladybugs_array = new Array<LadybugDB>() // array of repo instances
  Ladybugs = new Map<number, LadybugDB>() // map of repo instances
  Ladybugs_batch = new Map<number, LadybugDB>() // same but only in last GET (for finding repo instances to delete)

  LadybugSimulations_array = new Array<LadybugSimulationDB>() // array of repo instances
  LadybugSimulations = new Map<number, LadybugSimulationDB>() // map of repo instances
  LadybugSimulations_batch = new Map<number, LadybugSimulationDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Ladybug':
        return this.Ladybugs_array as unknown as Array<Type>
      case 'LadybugSimulation':
        return this.LadybugSimulations_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Ladybug':
        return this.Ladybugs_array as unknown as Map<number, Type>
      case 'LadybugSimulation':
        return this.LadybugSimulations_array as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
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

  GONG__StackPath: string = ""
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

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private ladybugService: LadybugService,
    private ladybugsimulationService: LadybugSimulationService,
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
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template 
    Observable<LadybugDB[]>,
    Observable<LadybugSimulationDB[]>,
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template
      this.ladybugService.getLadybugs(this.GONG__StackPath),
      this.ladybugsimulationService.getLadybugSimulations(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template
      this.ladybugService.getLadybugs(this.GONG__StackPath),
      this.ladybugsimulationService.getLadybugSimulations(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            ladybugs_,
            ladybugsimulations_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var ladybugs: LadybugDB[]
            ladybugs = ladybugs_ as LadybugDB[]
            var ladybugsimulations: LadybugSimulationDB[]
            ladybugsimulations = ladybugsimulations_ as LadybugSimulationDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Ladybugs_array = ladybugs

            // clear the map that counts Ladybug in the GET
            this.frontRepo.Ladybugs_batch.clear()

            ladybugs.forEach(
              ladybug => {
                this.frontRepo.Ladybugs.set(ladybug.ID, ladybug)
                this.frontRepo.Ladybugs_batch.set(ladybug.ID, ladybug)
              }
            )

            // clear ladybugs that are absent from the batch
            this.frontRepo.Ladybugs.forEach(
              ladybug => {
                if (this.frontRepo.Ladybugs_batch.get(ladybug.ID) == undefined) {
                  this.frontRepo.Ladybugs.delete(ladybug.ID)
                }
              }
            )

            // sort Ladybugs_array array
            this.frontRepo.Ladybugs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.LadybugSimulations_array = ladybugsimulations

            // clear the map that counts LadybugSimulation in the GET
            this.frontRepo.LadybugSimulations_batch.clear()

            ladybugsimulations.forEach(
              ladybugsimulation => {
                this.frontRepo.LadybugSimulations.set(ladybugsimulation.ID, ladybugsimulation)
                this.frontRepo.LadybugSimulations_batch.set(ladybugsimulation.ID, ladybugsimulation)
              }
            )

            // clear ladybugsimulations that are absent from the batch
            this.frontRepo.LadybugSimulations.forEach(
              ladybugsimulation => {
                if (this.frontRepo.LadybugSimulations_batch.get(ladybugsimulation.ID) == undefined) {
                  this.frontRepo.LadybugSimulations.delete(ladybugsimulation.ID)
                }
              }
            )

            // sort LadybugSimulations_array array
            this.frontRepo.LadybugSimulations_array.sort((t1, t2) => {
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
                  let _ladybugsimulation = this.frontRepo.LadybugSimulations.get(ladybug.LadybugSimulation_LadybugsDBID.Int64)
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

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            ladybugs.forEach(
              ladybug => {
                // insertion point for sorting
              }
            )
            ladybugsimulations.forEach(
              ladybugsimulation => {
                // insertion point for sorting
                ladybugsimulation.Ladybugs?.sort((t1, t2) => {
                  if (t1.LadybugSimulation_LadybugsDBID_Index.Int64 > t2.LadybugSimulation_LadybugsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.LadybugSimulation_LadybugsDBID_Index.Int64 < t2.LadybugSimulation_LadybugsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
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
          this.ladybugService.getLadybugs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            ladybugs,
          ]) => {
            // init the array
            this.frontRepo.Ladybugs_array = ladybugs

            // clear the map that counts Ladybug in the GET
            this.frontRepo.Ladybugs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ladybugs.forEach(
              ladybug => {
                this.frontRepo.Ladybugs.set(ladybug.ID, ladybug)
                this.frontRepo.Ladybugs_batch.set(ladybug.ID, ladybug)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field LadybugSimulation.Ladybugs redeeming
                {
                  let _ladybugsimulation = this.frontRepo.LadybugSimulations.get(ladybug.LadybugSimulation_LadybugsDBID.Int64)
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
            this.frontRepo.Ladybugs.forEach(
              ladybug => {
                if (this.frontRepo.Ladybugs_batch.get(ladybug.ID) == undefined) {
                  this.frontRepo.Ladybugs.delete(ladybug.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
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
          this.ladybugsimulationService.getLadybugSimulations(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            ladybugsimulations,
          ]) => {
            // init the array
            this.frontRepo.LadybugSimulations_array = ladybugsimulations

            // clear the map that counts LadybugSimulation in the GET
            this.frontRepo.LadybugSimulations_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            ladybugsimulations.forEach(
              ladybugsimulation => {
                this.frontRepo.LadybugSimulations.set(ladybugsimulation.ID, ladybugsimulation)
                this.frontRepo.LadybugSimulations_batch.set(ladybugsimulation.ID, ladybugsimulation)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear ladybugsimulations that are absent from the GET
            this.frontRepo.LadybugSimulations.forEach(
              ladybugsimulation => {
                if (this.frontRepo.LadybugSimulations_batch.get(ladybugsimulation.ID) == undefined) {
                  this.frontRepo.LadybugSimulations.delete(ladybugsimulation.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
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
