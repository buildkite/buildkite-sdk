import { VersionActions } from "nx/release";
import { basename, dirname } from "node:path/posix";
import type { Tree } from "nx/src/generators/tree";

const PATTERNS: { match: RegExp; version: RegExp }[] = [
    // Scoped to [project], so a version key in another table cannot win.
    {
        match: /\.toml$/,
        version: /(\[project\][\s\S]*?^version\s*=\s*")([^"]+)(")/m,
    },
    { match: /\.rb$/, version: /(VERSION\s*=\s*")([^"]+)(")/ },
    { match: /\.csproj$/, version: /(<Version>)([^<]+)(<\/Version>)/ },
];

function patternFor(file: string): RegExp {
    const entry = PATTERNS.find((p) => p.match.test(file));
    if (!entry) {
        throw new Error(`No version pattern known for ${file}`);
    }
    return entry.version;
}

export default class ManifestVersionActions extends VersionActions {
    private get file(): string | null {
        const file = this.finalConfigForProject.versionActionsOptions?.file;
        return typeof file === "string" ? file : null;
    }

    get validManifestFilenames(): string[] | null {
        const file = this.file;
        return file ? [basename(file)] : null;
    }

    async init(tree: Tree) {
        const file = this.file;
        if (
            file &&
            this.finalConfigForProject.manifestRootsToUpdate.length === 0
        ) {
            this.finalConfigForProject.manifestRootsToUpdate.push({
                path: dirname(file),
                preserveLocalDependencyProtocols:
                    this.finalConfigForProject.preserveLocalDependencyProtocols,
            });
        }
        await super.init(tree);
    }

    async readCurrentVersionFromSourceManifest(tree: Tree) {
        const file = this.file;
        if (!file) {
            return null;
        }

        const contents = tree.read(file, "utf-8");
        if (!contents) {
            throw new Error(`Could not read ${file}`);
        }

        const match = patternFor(file).exec(contents);
        if (!match) {
            throw new Error(`No version found in ${file}`);
        }

        return { currentVersion: match[2], manifestPath: file };
    }

    async readCurrentVersionFromRegistry() {
        return null;
    }

    async readCurrentVersionOfDependency() {
        return { currentVersion: null, dependencyCollection: null };
    }

    async updateProjectVersion(tree: Tree, newVersion: string) {
        const file = this.file;
        if (!file) {
            return [`Version ${newVersion} is carried by the git tag only`];
        }

        const contents = tree.read(file, "utf-8");
        if (!contents) {
            throw new Error(`Could not read ${file}`);
        }

        const pattern = patternFor(file);
        if (!pattern.test(contents)) {
            throw new Error(`No version found in ${file}`);
        }

        tree.write(file, contents.replace(pattern, `$1${newVersion}$3`));
        return [`Set version to ${newVersion} in ${file}`];
    }

    async updateProjectDependencies() {
        return [];
    }
}
