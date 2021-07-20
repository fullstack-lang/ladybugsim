// insertion point sub template for components imports 
  import { LadybugsTableComponent } from './ladybugs-table/ladybugs-table.component'
  import { LadybugSortingComponent } from './ladybug-sorting/ladybug-sorting.component'
  import { LadybugSimulationsTableComponent } from './ladybugsimulations-table/ladybugsimulations-table.component'
  import { LadybugSimulationSortingComponent } from './ladybugsimulation-sorting/ladybugsimulation-sorting.component'
  import { UpdatePositionEventsTableComponent } from './updatepositionevents-table/updatepositionevents-table.component'
  import { UpdatePositionEventSortingComponent } from './updatepositionevent-sorting/updatepositionevent-sorting.component'
  import { UpdateSpeedEventsTableComponent } from './updatespeedevents-table/updatespeedevents-table.component'
  import { UpdateSpeedEventSortingComponent } from './updatespeedevent-sorting/updatespeedevent-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfLadybugsComponents: Map<string, any> = new Map([["LadybugsTableComponent", LadybugsTableComponent],])
  export const MapOfLadybugSortingComponents: Map<string, any> = new Map([["LadybugSortingComponent", LadybugSortingComponent],])
  export const MapOfLadybugSimulationsComponents: Map<string, any> = new Map([["LadybugSimulationsTableComponent", LadybugSimulationsTableComponent],])
  export const MapOfLadybugSimulationSortingComponents: Map<string, any> = new Map([["LadybugSimulationSortingComponent", LadybugSimulationSortingComponent],])
  export const MapOfUpdatePositionEventsComponents: Map<string, any> = new Map([["UpdatePositionEventsTableComponent", UpdatePositionEventsTableComponent],])
  export const MapOfUpdatePositionEventSortingComponents: Map<string, any> = new Map([["UpdatePositionEventSortingComponent", UpdatePositionEventSortingComponent],])
  export const MapOfUpdateSpeedEventsComponents: Map<string, any> = new Map([["UpdateSpeedEventsTableComponent", UpdateSpeedEventsTableComponent],])
  export const MapOfUpdateSpeedEventSortingComponents: Map<string, any> = new Map([["UpdateSpeedEventSortingComponent", UpdateSpeedEventSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Ladybug", MapOfLadybugsComponents],
      ["LadybugSimulation", MapOfLadybugSimulationsComponents],
      ["UpdatePositionEvent", MapOfUpdatePositionEventsComponents],
      ["UpdateSpeedEvent", MapOfUpdateSpeedEventsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Ladybug", MapOfLadybugSortingComponents],
      ["LadybugSimulation", MapOfLadybugSimulationSortingComponents],
      ["UpdatePositionEvent", MapOfUpdatePositionEventSortingComponents],
      ["UpdateSpeedEvent", MapOfUpdateSpeedEventSortingComponents],
    ]
  )
