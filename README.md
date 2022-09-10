This module provides utilities to be used with [cuelang.org/go/cue].

Currently, it defines a type `Overlay` that allows feeding io/fs based file
systems, or files embedded using `//go:embed` to a CUE loader.
An Overlay value may be assigned to the Overlay field of a [load.Config].

[cuelang.org/go/cue]: https://pkg.go.dev/cuelang.org/go/cue
[load.Config]: https://pkg.go.dev/cuelang.org/go@v0.4.3/cue/load#Config
