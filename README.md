# Large ECDSA Collider
## What is Large ECDSA Collider?
The "Large ECDSA Collider" (inspired by Large Bitcoin Collider) is an experiment to find at least one collision of private ECDSA keys by creating addresses to private keys in a continuous 2^256-bit range. These are checked against the list of known Ethereum addresses with funds on them through an active Ethereum Node. In the rare event of a collision, the assets on the address in question would become accessible to the collision finder.

# Development Intent
Large ECDSA Collider takes inspiration from Large Bitcoin Collider to carry out an experiment in testing the claim of possibility vs randomness with a real-world implementation. It is not the intent of this project to claim possession of assets found in collision - which, depending on your place of residence, may even be illegal. However, it is not illegal to search the identity space for collisions.

## Goal and Working Principle
This repository is a proof-of-concept implementation of a worker in a distributed network of master and worker nodes. The idea is for a worker to register itself to a worker pool accounted by a discovery agent. Master nodes assign large contiguous subsets of generated identities, also known as 'wallets,' to worker nodes. Worker nodes operate asynchronously, attempting identity collisions on a given dataset. The dataset must match the 2^256-bit range for collisions to work. It is possible to use a real-time dataset, such as an active blockchain node.

## Current Implementation
This implementation of the ECDSA collider is self-sufficient in generating and performing identity collisions, and thanks to Golang's concurrency, its performance is noticeably faster than the original Large Bitcoin Collider client implementation. However, it should be noted that I have not run any standard benchmarks to guarantee the above claim on its performance. Since this implementation can only run self-sufficiently, it performs generation and collisions in parallel, depending on the processing cores available in its runtime environment. Also, collision search in this implementation is ephemeral, and duplicate entries in the collision pool may exist. Given its ephemeral nature, this implementation uses filesystem logs to record discovered collisions.

## Impact of Collision
In the event of a collision, the identity is compromised. It should be noted that since Ethereum is a popular choice for building layer-2 and layer-3 blockchain networks, the compromised address affects all the layers. Layer-1 blockchain networks, which define addresses within the fixed range of 2^256, are also susceptible to compromise.

## How should you proceed in the event of collision discovery?
A proper protocol remains unclear at this point. Here are two options that I consider to be the best available:

- Transfer the assets to a custodial address. Publishing the details of the discovery without the private key may help reach the rightful owner. It is possible for the original owner to provide proof of ownership through zero-knowledge protocol implementations, such as signature validation.

- Ignore the discovery and destroy the filesystem logs and all traces of the private key. It is still recommended to publish the details of the discovery on popular forums, depending on the network being searched, to inform the original owner of the detection.

# License
This project is meant for educational purposes only. You agree to take full responsibility in case of collision discovery and any associated movement of assets by cloning and/or running it. For all other cases, this project may not be used or cloned.