# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].

<!--
## Unreleased

### Added
### Changed
### Removed
-->

## [0.1.3][] - 2025-01-17

### Removed

* `utils/notify` to separate project
  [github.com/woozymasta/notify](https://github.com/WoozyMasta/notify)

[0.1.3]: https://github.com/WoozyMasta/steam/compare/v0.1.2...v0.1.3

## [0.1.2][] - 2025-01-12

### Added

* `utils/appid`
* `filedetails` support for splitting requested IDs into chunks to avoid
  exceeding the URI request length limit
* `filedetails` new `GetConcurrent` methods for `Query` which allows you to
  execute multiple API requests concurrently
* `filedetails` for `Query` added `SetConcurrency` and `SetChunkMax` methods

### Changed

* `serverlist` `Server.Region` change type to `int` to prevent overflow
* `serverlist` Extend tests
* `filedetails` Extend tests

[0.1.2]: https://github.com/WoozyMasta/steam/compare/v0.1.1...v0.1.2

## [0.1.1][] - 2025-01-10

### Added

* `utils/notify`

### Changed

* in `utils/latest` function `CompareVersions` return `int8` now

[0.1.1]: https://github.com/WoozyMasta/steam/compare/v0.1.0...v0.1.1

## [0.1.0][] - 2025-01-09

First public release

### Added

* `filedetails`
* `serverlist`
* `utils/latest`

[0.1.0]: https://github.com/WoozyMasta/steam/tree/v0.1.0

<!--links-->
[Keep a Changelog]: https://keepachangelog.com/en/1.1.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
