<div align="center">
  <h1><strong><span style="color: #0F234B;">m</span><span style="color: #00A078;">env</span></strong></h1>
  <p><b>Modern Environment Management for HPC</b></p>
  <p>
    <img src="https://img.shields.io/badge/Go-%3E%3D%201.21-00A078?labelColor=0F234B&style=flat-square&logo=go" alt="Go Version">
    <img src="https://img.shields.io/badge/Version-v0.1.1-00A078?labelColor=0F234B&style=flat-square" alt="Version">
    <img src="https://img.shields.io/badge/Build-Passing-00A078?labelColor=0F234B&style=flat-square" alt="Build Status">
    <img src="https://img.shields.io/badge/License-MIT-00A078?labelColor=0F234B&style=flat-square" alt="License">
  </p>
</div>
![Version](https://img.shields.io/github/v/release/Thedtk24/menv?color=00A078&style=flat-square)
![License](https://img.shields.io/github/license/Thedtk24/menv?color=0F234B&style=flat-square)

---

## The Problem

While Lmod's `module save` feature is useful, it has several limitations in modern High-Performance Computing (HPC) workflows:
- **Portability limitations** across different clusters.
- Relies on obscure or hidden file formats instead of clear, human-readable definitions.
- Misses the ability to seamlessly store specific environment variables alongside module loads.
- Lacks a native way to quickly audit your environment or seamlessly auto-load modules when entering a specific project directory.

**menv** solves this by providing reproducible, human-readable environments designed for researchers and developers.

## Key Features

- **Atomic Snapshots (YAML-based):** State-of-the-art environment definitions stored in transparent, versionable YAML files.
- **Automatic Hook (Auto-load on `cd`):** Generate shell wrappers that automatically load the environments when entering your project folders.
- **Doctor Command:** Easily perform portability checks of your environments when jumping between different HPC clusters.
- **Zero Overhead:** Written in Go, it runs natively with lightning-fast performance and no measurable overhead.

## Installation

You can install `menv` globally on your HPC account via our robust installation script:

```bash
curl -sSL https://raw.githubusercontent.com/Thedtk24/menv/main/install.sh | bash
```

> **Note:** The script will automatically fetch the correct binary for your OS and architecture via the GitHub API and install it in `~/.local/bin`.

## Quick Start

Save your current module configuration:
```bash
menv save my-project
```

Load your saved environment seamlessly:
```bash
menv load my-project
```

Add the zero-overhead auto-loader to your shell:
```bash
# Add this to your ~/.zshrc or ~/.bashrc
menv hook >> ~/.bashrc
```
*(Every time you `cd` into a directory containing a `.menv.lock` file, your environment is automatically loaded!)*

## Architecture

Under the hood, `menv` is distributed as a single static Go binary. It evaluates the active loaded modules by communicating with Lmod strictly through environment variables. For loading, it outputs pure shell instructions (`module purge`, `module load ...`) to ensure complete compatibility across any shell wrapper.

**Example of an `menv` YAML save file (`~/.menv/my-project.yaml`):**
```yaml
name: my-project
created: 2026-03-27T12:00:00Z
modules:
  - gcc/11.3.0
  - openmpi/4.1.4
  - cuda/11.8
env_vars: {}
```
This transparent architecture means you are always in control of your stack.

**`menv` is designed to be the perfect middle-ground:**
- **Rock-solid Security:** Operates completely in userspace—no root, no daemons. Keeps your SysAdmins happy!
- **Native Performance:** ZERO virtualization layers, allowing full and direct access to InfiniBand fabrics, NVLink, and specialized cluster hardware.
- **Cluster Friendly:** Perfect synergy with the cluster's existing Lmod software stack engineered by your HPC admins.

## Contributing & License

We welcome contributions to make `menv` the standard for HPC environment management! Feel free to open issues or submit pull requests.

This project is licensed under the **MIT License**. See the `LICENSE` file for more details.
