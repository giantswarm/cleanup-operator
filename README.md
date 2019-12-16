[![CircleCI](https://circleci.com/gh/giantswarm/cleanup-operator.svg?&style=shield)](https://circleci.com/gh/giantswarm/cleanup-operator) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/cleanup-operator/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/cleanup-operator)

# cleanup-operator

> An operator that disposes of any evidence of test stuff in your cluster.

This [Kubernetes][] operator watches for test resources in the cluster and
automatically disposes of any that are stale (by default older than 8 hours).

## Installing

TODO

## Usage 🚀

TODO

## Uninstalling

TODO

## Contributing 🤝

See known [issues][], if you found one that's not on the list or have a
suggestion for improvement, open a new issue. If you can, fork and send a PR,
it will be appreciated 💖.

## Hacking 🧰

### Building

`go build` will build the code.

TODO test, style

### Code

TODO

## Acknowledgements 👍

- [operatorkit][]

## License 📝

[![license-badge][]](LICENSE)


[kubernetes]: https://kubernetes.io/
[issues]: https://github.com/giantswarm/cleanup-operator/issues
[operatorkit]: https://github.com/giantswarm/operatorkit
[license-badge]: https://img.shields.io/github/license/giantswarm/cleanup-operator?style=for-the-badge
