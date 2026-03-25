# cdk-ec2-service-module-go

This repository contains the generated Go bindings for `cdk-ec2-service-module`.

Use this repo when:

- a Go CDK consumer needs the published module
- the service repo wants to import the shared EC2 service constructs from Go

Do not hand-maintain the implementation here. The source of truth remains:

- `https://github.com/Bh-an/cdk-ec2-service-module`

## Published Module

```text
github.com/Bh-an/cdk-ec2-service-module-go/cdkec2servicemodule
```

The generated Go module lives under:

```text
cdkec2servicemodule/
```

## Current Release

`v0.1.0`

## Contents

The generated package exposes the shared service contracts and the two concrete constructs:

- `Ec2DockerService`
- `PrivateEc2DockerService`

These bindings are generated from the TypeScript source repo and should be refreshed from there for each new release.

## Documentation Governance

Tracked docs in this repo must stay minimal and release-oriented.

- this repo documents the published Go module path and release line
- implementation detail and design rationale stay in the TypeScript source repo
- local operational notes belong under `project_docs/` and are not part of the published surface
