import { BlockStep } from './blockStep'
import { CommandStep } from './commandStep'
import { InputStep } from './inputStep'
import { TriggerStep } from './triggerStep'
import { WaitStep } from './waitStep'
import { DependsOn, Notify, NotifyEnum } from './types'

export interface GroupStep {
    group: string;
    steps: Array<BlockStep | CommandStep | InputStep | TriggerStep | WaitStep>;
    allow_dependency_failure?: boolean;
    depends_on?: Array<DependsOn | string> | null | string;
    if?: string;
    key?: string;
    label?: string;
    notify?: Array<Notify | NotifyEnum>;
    skip?: boolean | string;
}
