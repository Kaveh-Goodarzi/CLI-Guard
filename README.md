# guard 🔐

A minimal CLI tool written in Go for verifying file integrity using SHA256 hashes.

`guard` helps you detect file changes, additions, and deletions without complex architecture — using only Go’s standard library.

---

## ✨ Features

- Recursive directory scanning
- SHA256 hashing for file integrity
- Simple JSON-based manifest (`.guard.json`)
- Detects:
  - Modified files
  - Added files
  - Removed files
- No external dependencies
- Cross-platform binaries

---

## 📂 Project Structure

```
    CLI-Guard/
    ├── main.go // CLI entry point
    ├── scan.go // Recursive file scanning
    ├── hash.go // SHA256 hashing
    ├── manifest.go // Read/write .guard.json
    └── verify.go // init & verify logic

```


---


## 🚀 Installation

### Download binary (recommended)

Prebuilt binaries are available for:

- Linux (amd64, arm64)
- macOS (Intel & Apple Silicon)
- Windows

👉 See **GitHub Releases**

---

### Build from source

```bash
git clone https://github.com/Kaveh-Goodarzi/CLI-Guard.git
cd CLI-Guard
go build -o guard
```

---

## 🛠 Usage

### Initialize integrity manifest

```bash
guard init <path>
```

Creates a `.guard.json` file containing SHA256 hashes of all regular files in the specified directory and its subdirectories.

### Verify integrity

```bash
guard verify <path>
```

#### Outputs file status:

* OK – unchanged

* FAIL – modified

* REMOVED – deleted

* ADDED – new file

---

## 📄 Example Output

```text
main.go OK
config.yaml FAIL
assets/logo.png REMOVED
README.md ADDED
```

---

## 🐳 Docker

### Build image:

```bash
docker build -t guard .
```

### Run:

```bash
docker run --rm -v $(pwd):/data guard verify /data
```

---

## 🔐 Security & Integrity

This project is scanned using AI Code Guard in GitHub Actions to detect suspicious or unsafe code changes automatically.

---

## 🧪 CI / CD

* Automated build & verification via GitHub Actions

* Multi-platform releases on version tags

* Integrity & security checks on pull requests

