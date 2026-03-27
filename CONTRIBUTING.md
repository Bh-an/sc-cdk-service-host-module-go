# Contributing

## Important: This Repo Is Generated

The Go files under `cdkservicehostmodule/` are produced by `jsii-pacmak`. Do not hand-edit them. To change the constructs, work in [`sc-cdk-service-host-module`](https://github.com/Bh-an/sc-cdk-service-host-module) instead.

## Commits

[Conventional Commits](https://www.conventionalcommits.org/) format:

```
type(scope): short description
```

| Rule | Detail |
|------|--------|
| Types | `feat`, `fix`, `chore`, `docs`, `refactor`, `test`, `ci` |
| Scopes | `go`, `docs` |
| Subject | Imperative mood, under 72 chars, no trailing period |
| Granularity | One logical change per commit |

## Branches

| Branch | Purpose |
|--------|---------|
| `main` | Tagged releases |
| `dev` | Integration and release prep |
| `ci-cd` | Workflow and automation changes only |

## Pending Branch Protection Plan

This is the intended GitHub policy, but it is not enforced yet:

- `main` should require pull requests
- direct pushes to `main` should be blocked
- required checks should pass before merge
- `dev` should at least block force-push and deletion

Until GitHub protection is enabled, contributors should treat PR-based merges into `main` as mandatory.

## Before You Commit

```bash
cd cdkservicehostmodule && go build ./...
```

## Wrapper Regeneration

1. Make changes in the TypeScript source repo
2. Run `npm run package:go` there
3. Replace `cdkservicehostmodule/` with the output from `dist/go/cdkservicehostmodule/`
4. Run `go mod tidy` in `cdkservicehostmodule/`
5. Verify with `go build ./...`
6. Commit with `chore(go): regenerate wrapper for vX.Y.Z`

The release workflow automates this sequence from the tagged source repo.

## Releases

Wrapper releases are triggered by the source repo's release workflow — not by manual tagging here. The workflow receives a `version` and `source_tag` input and:

1. Checks out the source at the given tag
2. Regenerates the Go bindings
3. Commits, tags (both repo-level and module-level), and pushes
4. Creates a GitHub Release

The `RELEASE_REPO_TOKEN` secret is required for cross-repo checkout.

Reruns are safe when tags already exist remotely.
