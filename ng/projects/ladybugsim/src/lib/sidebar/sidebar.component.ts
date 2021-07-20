import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { LadybugService } from '../ladybug.service'
import { getLadybugUniqueID } from '../front-repo.service'
import { LadybugSimulationService } from '../ladybugsimulation.service'
import { getLadybugSimulationUniqueID } from '../front-repo.service'
import { UpdatePositionEventService } from '../updatepositionevent.service'
import { getUpdatePositionEventUniqueID } from '../front-repo.service'
import { UpdateSpeedEventService } from '../updatespeedevent.service'
import { getUpdateSpeedEventUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-ladybugsim-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo
  commitNb: number

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,

    // insertion point for per struct service declaration
    private ladybugService: LadybugService,
    private ladybugsimulationService: LadybugSimulationService,
    private updatepositioneventService: UpdatePositionEventService,
    private updatespeedeventService: UpdateSpeedEventService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.ladybugService.LadybugServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.ladybugsimulationService.LadybugSimulationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.updatepositioneventService.UpdatePositionEventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.updatespeedeventService.UpdateSpeedEventServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = true
            } else {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Ladybug part of the mat tree
      */
      let ladybugGongNodeStruct: GongNode = {
        name: "Ladybug",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Ladybug",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(ladybugGongNodeStruct)

      this.frontRepo.Ladybugs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Ladybugs_array.forEach(
        ladybugDB => {
          let ladybugGongNodeInstance: GongNode = {
            name: ladybugDB.Name,
            type: GongNodeType.INSTANCE,
            id: ladybugDB.ID,
            uniqueIdPerStack: getLadybugUniqueID(ladybugDB.ID),
            structName: "Ladybug",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          ladybugGongNodeStruct.children.push(ladybugGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the LadybugSimulation part of the mat tree
      */
      let ladybugsimulationGongNodeStruct: GongNode = {
        name: "LadybugSimulation",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "LadybugSimulation",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(ladybugsimulationGongNodeStruct)

      this.frontRepo.LadybugSimulations_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.LadybugSimulations_array.forEach(
        ladybugsimulationDB => {
          let ladybugsimulationGongNodeInstance: GongNode = {
            name: ladybugsimulationDB.Name,
            type: GongNodeType.INSTANCE,
            id: ladybugsimulationDB.ID,
            uniqueIdPerStack: getLadybugSimulationUniqueID(ladybugsimulationDB.ID),
            structName: "LadybugSimulation",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          ladybugsimulationGongNodeStruct.children.push(ladybugsimulationGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Ladybugs
          */
          let LadybugsGongNodeAssociation: GongNode = {
            name: "(Ladybug) Ladybugs",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ladybugsimulationDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "LadybugSimulation",
            associationField: "Ladybugs",
            associatedStructName: "Ladybug",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ladybugsimulationGongNodeInstance.children.push(LadybugsGongNodeAssociation)

          ladybugsimulationDB.Ladybugs?.forEach(ladybugDB => {
            let ladybugNode: GongNode = {
              name: ladybugDB.Name,
              type: GongNodeType.INSTANCE,
              id: ladybugDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLadybugSimulationUniqueID(ladybugsimulationDB.ID)
                + 11 * getLadybugUniqueID(ladybugDB.ID),
              structName: "Ladybug",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LadybugsGongNodeAssociation.children.push(ladybugNode)
          })

        }
      )

      /**
      * fill up the UpdatePositionEvent part of the mat tree
      */
      let updatepositioneventGongNodeStruct: GongNode = {
        name: "UpdatePositionEvent",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "UpdatePositionEvent",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(updatepositioneventGongNodeStruct)

      this.frontRepo.UpdatePositionEvents_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.UpdatePositionEvents_array.forEach(
        updatepositioneventDB => {
          let updatepositioneventGongNodeInstance: GongNode = {
            name: updatepositioneventDB.Name,
            type: GongNodeType.INSTANCE,
            id: updatepositioneventDB.ID,
            uniqueIdPerStack: getUpdatePositionEventUniqueID(updatepositioneventDB.ID),
            structName: "UpdatePositionEvent",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          updatepositioneventGongNodeStruct.children.push(updatepositioneventGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the UpdateSpeedEvent part of the mat tree
      */
      let updatespeedeventGongNodeStruct: GongNode = {
        name: "UpdateSpeedEvent",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "UpdateSpeedEvent",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(updatespeedeventGongNodeStruct)

      this.frontRepo.UpdateSpeedEvents_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.UpdateSpeedEvents_array.forEach(
        updatespeedeventDB => {
          let updatespeedeventGongNodeInstance: GongNode = {
            name: updatespeedeventDB.Name,
            type: GongNodeType.INSTANCE,
            id: updatespeedeventDB.ID,
            uniqueIdPerStack: getUpdateSpeedEventUniqueID(updatespeedeventDB.ID),
            structName: "UpdateSpeedEvent",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          updatespeedeventGongNodeStruct.children.push(updatespeedeventGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.uniqueIdPerStack] != undefined) {
              if (memoryOfExpandedNodes[node.uniqueIdPerStack]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_table: ["github_com_fullstack_lang_ladybugsim_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_ladybugsim_go_table: ["github_com_fullstack_lang_ladybugsim_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_ladybugsim_go_presentation: ["github_com_fullstack_lang_ladybugsim_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet( node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_ladybugsim_go_editor: ["github_com_fullstack_lang_ladybugsim_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
