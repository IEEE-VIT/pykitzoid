# pykitzoid

![Hacktoberfest 2025 banner](./pykitzoid_hacktoberfest.png)

# Hacktoberfest 2025 with IEEE-VIT :blue_heart:

This is a repository containing a Python package coded in Go, with the motive of providing ML support.

Support open source software by participating in [Hacktoberfest](https://hacktoberfest.digitalocean.com) and get goodies and digital badges! :blue_heart:

> Please check all issues labelled as `hacktoberfest` to start contributing!

Kindly consider leaving a :star: if you like the repository and our organisation.
## Project Goals

### **Purpose**

- Provide clear, well-documented building blocks for ML (data handling, simple models, utilities) implemented primarily in **Go**, with optional Python access points.  

### **Scope**

- Small, focused algorithms and demos you can run **locally**.  
- Friendly for first-time contributors: readable code, bite-sized issues, and a straightforward workflow.

### **Target Audience**

- Students & self-learners who want to understand “how it works under the hood.”  
- Hacktoberfest / open-source contributors who prefer approachable issues.  
- Practitioners curious about Go in ML pipelines.

## Key Features & Overview

- **Go-first implementations**  
  The project’s core is in Go (currently ~93% Go). This keeps binaries fast and the code explicit/learnable.  
  [View on GitHub → IEEE-VIT/pykitzoid](https://github.com/IEEE-VIT/pykitzoid.git)

- **Python wrapper (work-in-progress)**  
  A thin Python interface (see `pyproject.toml`) is scaffolded so users can experiment from notebooks while the logic stays in Go.  
  [View pyproject.toml](https://github.com/IEEE-VIT/pykitzoid/blob/main/pyproject.toml)

- **Self-contained demos**  
  The `demo/` directory includes runnable examples that show how to use the library end-to-end. (Add more demos as algorithms land.)  
  [View demo folder](https://github.com/IEEE-VIT/pykitzoid/tree/main/demo)

- **Contributor-friendly defaults**
  - A `Makefile` with handy targets (build/test/format) for a smooth DX.  
    [View Makefile](https://github.com/IEEE-VIT/pykitzoid/blob/main/Makefile)
  - Clear [CONTRIBUTING guidelines](https://github.com/IEEE-VIT/pykitzoid/blob/main/CONTRIBUTING.md).  
  - MIT-licensed. [View License](https://github.com/IEEE-VIT/pykitzoid/blob/main/LICENSE)


## Getting Started

- Fork it.

- Clone your forked repo and move inside it:

`git clone https://github.com/IEEE-VIT/pykitzoid.git && cd pykitzoid`

- Checkout to a new branch to work on an issue:

`git checkout -b my-amazing-feature`

- Get started working!

- Once you're all done coding, it's time to open a PR :)
  Run the following commands from the root of the project directory:

`git add .`

`git commit -m "A short description about the feature."`

`git push origin <my-amazing-feature>`

Open your forked repo in your browser and then raise a PR to the `main` branch of this repository!

## Contributing

To start contributing, check out [CONTRIBUTING.md](https://github.com/IEEE-VIT/ToDo-iOS/blob/master/contributing.md). New contributors are always welcome to support this project. If you want something gentle to start with, check out issues labelled `easy` or `good-first-issue`. Check out issues labelled `hacktoberfest` if you are up for some fun hacktoberfest goodies! :)

## License

See the [LICENSE](https://github.com/kitrak-rev/pykitzoid/blob/main/LICENSE) file for license rights and limitations (MIT).
