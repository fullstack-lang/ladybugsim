// insertion point sub template for components imports 
  import { LadybugsTableComponent } from './ladybugs-table/ladybugs-table.component'
  import { LadybugSortingComponent } from './ladybug-sorting/ladybug-sorting.component'
  import { LadybugSimulationsTableComponent } from './ladybugsimulations-table/ladybugsimulations-table.component'
  import { LadybugSimulationSortingComponent } from './ladybugsimulation-sorting/ladybugsimulation-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfLadybugsComponents: Map<string, any> = new Map([["LadybugsTableComponent", LadybugsTableComponent],])
  export const MapOfLadybugSortingComponents: Map<string, any> = new Map([["LadybugSortingComponent", LadybugSortingComponent],])
  export const MapOfLadybugSimulationsComponents: Map<string, any> = new Map([["LadybugSimulationsTableComponent", LadybugSimulationsTableComponent],])
  export const MapOfLadybugSimulationSortingComponents: Map<string, any> = new Map([["LadybugSimulationSortingComponent", LadybugSimulationSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Ladybug", MapOfLadybugsComponents],
      ["LadybugSimulation", MapOfLadybugSimulationsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Ladybug", MapOfLadybugSortingComponents],
      ["LadybugSimulation", MapOfLadybugSimulationSortingComponents],
    ]
  )
