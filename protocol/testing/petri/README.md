# Petri Loadtesting

## Overview

* The load-tests defined here depend on [petri](https://github.com/skip-mev/petri)
* The load-test process is as follows:
    * The load-tests spin up a dydxv4 network with `NUM_NODES` + `NUM_VALIDATORS` all running on the network `dydx-1`, with the `dydxprotocol-base` image
    * Each validator-node is configured to run with a separate oracle-process
    * The network is configured to store all markets defined in `fixtures/markets.json`
    * Each oracle will be reporting the price of all markets in `fixtures/markets.json` as determined by the [gate provider](https://github.com/skip-mev/slinky/tree/main/providers/websockets/gate)
        * The specific configurations used for the oracles are defined in `suite.go::updateOracleConfig`
    * 350 tx/s will be broadcast to the network
    * The `PETRI_LOAD_TEST_KEEP_ALIVE` environment variable will determine whether or not the network should be killed after the tests are finished

## Configuration

* The following environment variables determine how the load-tests should be executed
    * `PETRI_LOAD_TEST_KEEP_ALIVE`: If set to `true` (`export PETRI_LOAD_TEST_KEEP_ALIVE=true`), the network will only be killed after an `INTERRUPT`/`TERMINATE` signal is issued (via `ctrl-c`). If not set, the network will be killed immediately after loadtest completion
    * `PETRI_LOAD_TEST_PROVIDER_TYPE`: oneof(`docker`, `digitalocean`), this determines whether to run the loadtests locally (in docker), or remotely (via digital ocean)
        * If `digitalocean` is provided, a few additional parameters are needed, specifically:
            * `PETRI_LOAD_TEST_DIGITAL_OCEAN_API_KEY`: The digital ocean API-key to use when spawning / killing droplets
            * `PETRI_LOAD_TEST_DIGITAL_OCEAN_IMAGE_ID`: The image (not the chain-binary image, but the OS) that each droplet should be deployed with
                * See [here](https://github.com/skip-mev/petri/blob/main/contrib/docs/digitalocean.md) on creating + pushing the droplet image to use
