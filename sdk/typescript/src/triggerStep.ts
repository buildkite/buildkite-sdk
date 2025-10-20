import { Build, DependsOn, SoftFail } from './types'

export interface TriggerStep {
    allow_dependency_failure?: boolean;
    async?: boolean;
    branches?: string[] | string;
    build?: Build;
    depends_on?: Array<DependsOn | string> | null | string;
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    skip?: boolean | string;
    soft_fail?: SoftFail[] | boolean;
    trigger: string;
}
