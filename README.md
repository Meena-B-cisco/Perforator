# Perforator (v0.0.5)

This repository contains the prebuilt binary and setup for **[yandex/perforator](https://github.com/yandex/perforator)** version `v0.0.5`.

Perforator is a lightweight Linux profiler that uses the `perf_event` interface to record CPU activity and generate flamegraphs.

---

## 🔧 Setup

### Clone the Repository

```bash
git clone https://github.com/Meena-B-cisco/Perforator.git
cd Perforator
git lfs pull

```
### Build from source


If you're using the source code under `perforator-0.0.5`, build it using:

```cd perforator-0.0.5
./ya make -r perforator/bundle/
```


### Run the CLI Binary

To use the `perforator` CLI binary directly, set an environment variable pointing to its full path:

```bash
export perforator_cli=/full/path/to/perforator
```

Follow the tutorial here: https://perforator.tech/docs/en/tutorials/python-profiling
