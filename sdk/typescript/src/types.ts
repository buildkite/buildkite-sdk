/**
 * The state that the build is set to when the build is blocked by this block step
 */
export enum BlockedState {
    Failed = "failed",
    Passed = "passed",
    Running = "running",
}

export interface DependsOnClass {
    allow_failure?: boolean;
    step?: string;
}

/**
 * A list of input fields required to be filled out before unblocking the step
 */
export interface Field {
    /**
     * The value that is pre-filled in the text field
     *
     * The value of the option(s) that will be pre-selected in the dropdown
     */
    default?: string[] | string;
    /**
     * The format must be a regular expression implicitly anchored to the beginning and end of
     * the input and is functionally equivalent to the HTML5 pattern attribute.
     */
    format?: string;
    /**
     * The explanatory text that is shown after the label
     */
    hint?: string;
    /**
     * The meta-data key that stores the field's input
     */
    key: string;
    /**
     * Whether the field is required for form submission
     */
    required?: boolean;
    /**
     * The text input name
     */
    text?: string;
    /**
     * Whether more than one option may be selected
     */
    multiple?: boolean;
    options?: Option[];
    /**
     * The text input name
     */
    select?: string;
}

export interface Option {
    /**
     * The text displayed directly under the select field’s label
     */
    hint?: string;
    /**
     * The text displayed on the select list item
     */
    label: string;
    /**
     * Whether the field is required for form submission
     */
    required?: boolean;
    /**
     * The value to be stored as meta-data
     */
    value: string;
}
