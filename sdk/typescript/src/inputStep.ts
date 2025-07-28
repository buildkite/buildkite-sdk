import { DependsOn, Field } from './types'

export interface InputStep {
    input: string;
    fields: Field[];
    allow_dependency_failure?: boolean;
    branches?: string[] | string;
    depends_on?: Array<DependsOn | string> | null | string;
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    prompt?: string;
    blocked_state: 'passed' | 'failed' | 'running'
}
