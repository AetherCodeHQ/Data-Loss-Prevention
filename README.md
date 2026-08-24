# 🔐 Data Loss Prevention

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Security tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`security` `cryptography` `cli` `golang` `regex` `io`

---

## What is Data-Loss-Prevention?

**Data-Loss-Prevention** is a security-focused tool that analyzes and validates code, configurations, or data for vulnerabilities.

## Features

- ✅ `min()` — Min
- ✅ Streaming file processing
- ✅ Pattern matching and analysis
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Data-Loss-Prevention.git
cd Data-Loss-Prevention

# Build
go build -o data-loss-prevention .

# Run
./data-loss-prevention Usage: dlp-scanner <file>
```

### Or directly with `go run`:
```bash
go run main.go Usage: dlp-scanner <file>
```

## Usage

```bash
# Basic usage
./data-loss-prevention Usage: dlp-scanner <file>

# With flags
./data-loss-prevention Usage: dlp-scanner <file> value Usage: dlp-scanner <file>
```

### Example Output

```
$ ./data-loss-prevention Usage: dlp-scanner <file>
Usage: dlp-scanner <file>
Error:
  [LINE %d] %s detected: %s\n
```

## Project Structure

```
Data-Loss-Prevention/
  main.go          # Entry point (56 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
