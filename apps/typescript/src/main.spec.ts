import { Pipeline } from "@buildkite/buildkite-sdk";

describe("toJSON()", () => {
    it("should work", () => {
        expect(new Pipeline().toJSON()).toBe(
            JSON.stringify({}, null, 4)
        );
    });
});
