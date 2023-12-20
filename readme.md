# Litmus Chaos thin client

Just a Litmus Chaos Control Plane thin client with minimal dependencies

If you're looking for the official thing, please check the [litmuschaos/litmusctl](https://github.com/litmuschaos/litmusctl) repository.

This repo is mainly for using on my Litmus Chaos terraform provider as the official client imports way too much stuff.

## Depdencies

This project depends only on hasura-go-client and it's transient dependencies.

If in the future I'm feeling to... I'll remove it too and to plain graphql queries.
