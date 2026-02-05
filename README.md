# Simple Serialize Implementation

An educational implementation of the Simple Serialize (SSZ) specification, focusing on the core principles used in the Ethereum Lean Consensus.

## Features
* **Little-Endian Serialization**: Byte-perfect encoding for basic types.
* **Offset Management**: Proper handling of variable-length lists within containers.
* **Merkleization Engine**: 
  * Pre-computed Zero-Hash Ladder for performance.
  * Recursive `HashLayers` for calculating the Merkle Root.
  * `Mix-In-Length` logic for list integrity.
* **Packing**: Efficient 32-byte chunking of serialized data.

## Implementation Progress
- [x] Basic Types (uint32, uint64)
- [x] Fixed-size Containers
- [x] Variable-size Lists
- [x] Merkle Tree reduction
- [x] Root generation for nested structs