import { BlockedState, DependsOnClass, Field } from './types'

export interface BlockStep {
    block: string;

    allow_dependency_failure?: boolean;
    blocked_state?: BlockedState;
    branches?: string[] | string;
    depends_on?: Array<DependsOnClass | string> | null | string;
    fields?: Field[];
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    prompt?: string;
}
