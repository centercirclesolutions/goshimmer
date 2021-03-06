# v0.1.0 - 2020-01-31

> Note that this release is a complete breaking change, therefore node operators are instructed to upgrade.

This release mainly integrates the shared codebase between Hornet and GoShimmer called [hive.go](https://github.com/iotaledger/hive.go), 
to foster improvements in both projects simultaneously. Additionally, the autopeering code from [autopeering-sim](https://github.com/iotaledger/autopeering-sim)
has been integrated into the codebase. For developers a Go client library is now available to communicate
with a GoShimmer node in order to issue, get and search zero value transactions and retrieve neighbor information.

A detailed list about the changes in v0.1.0 can be seen under the given [milestone](https://github.com/iotaledger/goshimmer/milestone/1?closed=1).

* Adds config keys to make bind addresses configurable
* Adds a `relay-checker` tool to check transaction propagation
* Adds database versioning
* Adds concrete shutdown order for plugins
* Adds BadgerDB garbage collection routine
* Adds SPA dashboard with TPS and memory chart, neighbors and Tangle explorer (`spa` plugin)
* Adds option to manually define autopeering seed/private keys
* Adds Golang CI and GitHub workflows as part of the continuous integration pipeline
* Adds back off policies when doing network operations
* Adds an open port check which checks whether NATs are configured properly
* Adds `netutil` package
* Adds the glumb visualizer backend as a plugin called `graph` (has to be manually enabled, check the readme)
* Adds rudimentary PoW as a rate control mechanism
* Adds the autopeering code from autopeering-sim
* Adds association between transactions and their address
* Adds the possibility to allow a node to function/start with gossiping disabled
* Adds buffered connections for gossip messages
* Adds an OAS/Swagger specification file for the web API
* Adds web API and refactors endpoints: `broadcastData`, `findTransactionHashes`,
`getNeighbors`, `getTransactionObjectsByHash`, `getTransactionTrytesByHash`, `getTransactionsToApprove`, `spammer`
* Adds a Go client library over the web API
* Adds a complete rewrite of the gossiping layer
* Fixes the autopeering visualizer to conform to the newest autopeering changes.  
  The visualizer can be accessed [here](http://ressims.iota.cafe/).
* Fixes parallel connections by improving the peer selection mechanism
* Fixes that LRU caches are not flushed correctly up on shutdown
* Fixes consistent application naming (removes instances of sole "Shimmer" in favor of "GoShimmer")
* Fixes several analysis server related issues
* Fixes race condition in the PoW code
* Fixes race condition in the `daemon` package
* Fixes several race conditions in the `analysis` package
* Fixes the database not being closed when the node shuts down
* Fixes `webauth` plugin to function properly with the web API
* Fixes solidification related issues
* Fixes several instances of overused goroutine spawning
* Upgrades to [BadgerDB](https://github.com/dgraph-io/badger) v2.0.1
* Upgrades to latest [iota.go](https://github.com/iotaledger/iota.go) library
* Removes sent count from spammed transactions
* Removes usage of `errors.Identifiable` and `github.com/pkg/errors` in favor of standard lib `errors` package
* Use `network`, `parameter`, `events`, `database`, `logger`, `daemon`, `workerpool` and `node` packages from hive.go
* Removes unused plugins (`zmq`, `dashboard`, `ui`)