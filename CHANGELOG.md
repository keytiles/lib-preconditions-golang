# Versioning policy

We are following [Semantic versioning](https://semver.org/) in this library

# Changes in releases

## pre-release 0.3.1

Fixes:

- In v0.2.1 signature of `EnsureParamNonNil()` and `EnsureParamNonNilPub()` methods were changed but usage revealed bug. It did not work. Now test case
  was added, and implementation was changed so it works now.

## pre-release 0.3.0

New features:
- Added `EnsureParamState()` and `EnsureParamStatePub()` methods

## pre-release 0.2.1

Fixes:

- Fixing signature of `EnsureParamNonNil()` and `EnsureParamNonNilPub()` methods so it can also be used for not just pointers but function variables
  or anything which can be Nil really.

## pre-release 0.2.0

Fixes / Breaking (early):

- All `kt_precond` methods return `*kt_errors.FaultBuilder` instead of a `kt_errors.Fault` so we can tweak further the Fault before we build and throw it.

New features:

- Added `EnsureState()` and `EnsureStatePub` methods to check states and throw `kt_errors.IllegalStateFault` if condition not met.
- Added quick example code to README.md

## pre-release 0.1.0

Initial commit - quick release to start tasting it and validate if it is helpful.
