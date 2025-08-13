# Copilot Stats

<p align="center">
  <img src="./img.png" alt="Copilot Stats screenshot" />
</p>

A simple Go CLI to display your GitHub Copilot usage statistics, including premium interactions, chat, and completions quotas, with styled terminal output.

## Installation
### Via Go tooling
```bash
go install github.com/sahaj-b/copilot-stats@latest
# This will install the go-attend binary in $GOBIN, make sure it's in your PATH
```

### From source
- Clone the repository and build the binary:
   ```bash
   git clone https://github.com/sahaj-b/copilot-stats.git
   cd copilot-stats
   go build -o copilot-stats
   ```
- (Optional) Install to your PATH for easy use:
   ```bash
   install -m 0755 ./copilot-stats ~/.local/bin/copilot-stats
   # ensure ~/.local/bin is on PATH
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   ```
   Or system-wide (requires sudo):
   ```bash
   sudo install -m 0755 ./copilot-stats /usr/local/bin/copilot-stats
   ```

## Setup
Ensure you have a valid Copilot OAuth token. The tool will look for it in:
   - The `GITHUB_TOKEN` environment variable (Codespaces)
   - `~/.config/github-copilot/hosts.json` or `~/.config/github-copilot/apps.json` (if authorized via the Copilot CLI)

## Usage
Run the CLI:
```sh
copilot-stats   # if installed to PATH
# or
./copilot-stats
```

> [!NOTE]
> Disable colors by setting the NO_COLOR environment variable:
> `NO_COLOR=1 copilot-stats`
