import {
    CacheObject,
    ConcurrencyMethod,
    DependsOn,
    MatrixObject,
    Notify,
    NotifyEnum,
    Retry,
    Signature,
    SoftFail,
} from './types'

interface CommandStepOptionalAttributes {
    agents?: string[] | { [key: string]: any };
    allow_dependency_failure?: boolean;
    artifact_paths?: string[] | string;
    branches?: string[] | string;
    cache?: string | string[] | CacheObject;
    cancel_on_build_failing?: boolean;
    concurrency?: number;
    concurrency_group?: string;
    concurrency_method?: ConcurrencyMethod;
    depends_on?: Array<DependsOn | string> | null | string;
    env?: { [key: string]: any };
    id?: string;
    identifier?: string;
    if?: string;
    key?: string;
    label?: string;
    matrix?: Array<boolean | number | string> | MatrixObject;
    name?: string;
    notify?: Array<Notify | NotifyEnum>;
    parallelism?: number;
    plugins?: Array<{ [key: string]: any } | string> | { [key: string]: any };
    priority?: number;
    retry?: Retry;
    signature?: Signature;
    skip?: boolean | string;
    soft_fail?: SoftFail[] | boolean;
    timeout_in_minutes?: number;
}

export interface SingleCommandStep extends CommandStepOptionalAttributes {
    command: string
}

export interface MultipleCommandStep extends CommandStepOptionalAttributes {
    commands: string[]
}

export type CommandStep = SingleCommandStep | MultipleCommandStep
