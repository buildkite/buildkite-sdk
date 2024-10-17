// This file is auto-generated. Do not edit.
import * as fs from "fs";
import * as types from "./types";

class StepBuilder {
	private steps: any[] = [];

	public write() {
		fs.writeFileSync("pipeline.json", JSON.stringify({ steps: this.steps }, null, 4));
	}
    // A block step is used to pause the execution of a build and wait on a team member to unblock it using the web or the API.
    public addBlockStep(args: types.Block): this {
        this.steps.push({ ...args });
        return this;
    }
    // A command step runs one or more shell commands on one or more agents.
    public addCommandStep(args: types.Command): this {
        this.steps.push({ ...args });
        return this;
    }
    // A group step can contain various sub-steps, and display them in a single logical group on the Build page.
    public addGroupStep(args: types.Group): this {
        this.steps.push({ ...args });
        return this;
    }
    // An input step is used to collect information from a user.
    public addInputStep(args: types.Input): this {
        this.steps.push({ ...args });
        return this;
    }
    // A trigger step creates a build on another pipeline.
    public addTriggerStep(args: types.Trigger): this {
        this.steps.push({ ...args });
        return this;
    }
    // A wait step waits for all previous steps to have successfully completed before allowing following jobs to continue.
    public addWaitStep(args: types.Wait): this {
        this.steps.push({ ...args });
        return this;
    }

}
export default StepBuilder;