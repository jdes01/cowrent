import { DeepReadonly } from 'ts-essentials'


export type WorkspaceResponse = DeepReadonly<{
    uuid: string
    name: string
    seats: number
}>


export type CoworkingResponse = DeepReadonly<{
    workspaces: WorkspaceResponse[]
    uuid: string
    name: string
}>
