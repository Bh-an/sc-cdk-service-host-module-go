# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Wrapper releases track the source repo [`sc-cdk-service-host-module`](https://github.com/Bh-an/sc-cdk-service-host-module). Each entry here reflects a regeneration from the corresponding source version.

## [0.3.0] - 2026-03-15

### Changed
- **Breaking:** regenerated bindings for renamed constructs (`PublicServiceHost`, `PrivateServiceHost`)
- Automated release workflow triggered by source repo tagging

### Added
- Rerun-safe release pipeline (handles pre-existing tags)

## [0.2.0] - 2026-03-10

### Changed
- Regenerated wrapper for `PrivateServiceHost` and exposure control additions

## [0.1.1] - 2026-03-05

### Fixed
- Corrected module path alignment with `sc-` prefixed repo name
- Refreshed wrapper checksums

## [0.1.0] - 2026-03-01

### Added
- Initial generated Go wrapper for `cdkservicehostmodule`
- Published as `github.com/Bh-an/sc-cdk-service-host-module-go/cdkservicehostmodule`
