# pykitzoid

![Hacktoberfest 2025 banner](./pykitzoid_hacktoberfest.png)

## Hacktoberfest 2025 with IEEE-VIT

`pykitzoid` is an educational repository for machine learning implementations in Go, with Python-facing demos and wrappers where useful.

Support open source by participating in [Hacktoberfest](https://hacktoberfest.digitalocean.com).

> Check issues labeled `hacktoberfest`, `good first issue`, or `documentation` to start contributing.

## Project Goals

- Build clear, well-documented ML building blocks in Go.
- Keep examples small and runnable on local machines.
- Make contribution flow friendly for first-time contributors.
- Bridge Go implementations with Python usage patterns for learning.

## Repository Structure

- `algorithms/`: core algorithm implementations and sample datasets.
- `demo/`: demo Go library and Python script that loads the shared object.
- `src/`: Python package scaffolding.
- `Makefile`: helper targets for build and cleanup.

## Local Usage

### Prerequisites

- Go (1.20+ recommended)
- Python (3.10+ recommended)
- `make`

### Run the demo

```bash
make build_demo
python3 demo/demo_app.py
```

This builds `demo/build/library.so` from `demo/library.go` and executes the Python ctypes demo.

### Build algorithm shared libraries

```bash
make build
ls src/bin
```

This compiles Go sources under `algorithms/` into shared libraries in `src/bin/`.

### Clean build artifacts

```bash
make clear      # remove src/bin/*
make clearall   # remove src/bin/* and demo/build/*
```

## Getting Started as a Contributor

1. Fork this repository.
2. Clone your fork:

```bash
git clone https://github.com/<your-username>/pykitzoid.git
cd pykitzoid
```

3. Create a branch:

```bash
git checkout -b docs/improve-readme
```

4. Make changes, then commit and push:

```bash
git add .
git commit -m "docs: improve README"
git push origin docs/improve-readme
```

5. Open a pull request against the `main` branch.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for contribution guidelines and workflow.

## License

Distributed under the [MIT License](./LICENSE).
