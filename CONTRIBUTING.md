# Contributing

## Commit Standard

Use Conventional Commits:

```text
type(scope): short description
```

- Allowed types: `feat`, `fix`, `chore`, `docs`, `refactor`, `test`, `ci`
- Common scopes: `go`, `docs`
- Keep the subject imperative, under 72 characters, and without a trailing period
- Do not add AI attribution lines

## Branching

- `main` is the stable branch
- `dev` is the shared integration and release-prep branch

## Wrapper Regeneration

This repo is generated output. Regenerate it from `sc-cdk-service-host-module`:

1. update the TypeScript source repo
2. run `npm run package:go`
3. sync `dist/go/cdkservicehostmodule/` into this repo
4. run `go mod tidy` in `cdkservicehostmodule/`
5. verify with `go build ./...`
