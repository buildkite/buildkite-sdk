import { DependsOn } from "./types";

export interface WaitStep {
    allow_dependency_failure?: boolean;
    branches?: string[] | string;
    continue_on_failure?: boolean;
    depends_on?: Array<DependsOn | string> | null | string;
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    name?: string;
    wait: string;
}
