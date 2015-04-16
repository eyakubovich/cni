# cni - the Container Network Interface

## What is CNI?

CNI, the _Container Network Interface_, is a proposed standard for configuring network interfaces for Linux application containers.
The standard consists of a simple specification for how executable plugins can be used to configure network namespaces.
The specification itself is contained in [SPEC.md](SPEC.md)

## Why develop CNI?

Application containers on Linux are a rapidly evolving area, and within this space networking is a particularly unsolved problem, as it is highly environment-specific.
We believe that every container runtime will seek to solve the same problem of making the network layer pluggable.
In order to avoid duplication, we think it is prudent to define a common interface between the network plugins and container execution.
Hence we are proposing this specification, along with an initial set of plugins that can be used by different container runtime systems.

## How do I use CNI?

TODO...
