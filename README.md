# go-lint-generics
a linting learning experiment

This repo was deliberately initialised to use Go version `1.17`

The code in main.go makes an API request to [Cat Facts](https://catfact.ninja/) to get some data

To install `golangci-lint` locally
```
brew install golangci-lint
```
```
brew upgrade golangci-lint
```
To run `golangci-lint` locally
```
golangci-lint run
```

# Observations
1. Linting running and working fine on Go version `1.17`
2. Opened a PR to bump Go version to `1.19` [https://github.com/Rosalita/go-lint-generics/pull/1](https://github.com/Rosalita/go-lint-generics/pull/1) this PR encountered the following error

```
  Running [/home/runner/golangci-lint-1.46.2-linux-amd64/golangci-lint run --out-format=github-actions] in [] ...
  panic: load embedded ruleguard rules: rules/rules.go:13: can't load fmt
  
  goroutine 1 [running]:
  github.com/go-critic/go-critic/checkers.init.22()
  	github.com/go-critic/go-critic@v0.6.3/checkers/embedded_rules.go:47 +0x4b4
```

This error is explicitly listed on [https://github.com/golangci/golangci-lint/issues/2649](https://github.com/golangci/golangci-lint/issues/2649)

Tried bumping `golangci-lint` up to `1.49.0` this fixes gocritics rule guard failling to load rules.

