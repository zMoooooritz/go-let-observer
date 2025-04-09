# HLL Observer

[![Latest Release](https://img.shields.io/github/release/zMoooooritz/go-let-observer.svg?style=for-the-badge)](https://github.com/zMoooooritz/go-let-observer/releases)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/zMoooooritz/go-let-observer)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/zMoooooritz/go-let-observer/build.yml?branch=master&style=for-the-badge)](https://github.com/zMoooooritz/nachrichten/actions)
[![Go ReportCard](https://goreportcard.com/badge/github.com/zMoooooritz/go-let-observer?style=for-the-badge)](https://goreportcard.com/report/zMoooooritz/go-let-observer)

A **Go-based observer tool** for the game **Hell Let Loose**, providing a graphical interface for monitoring server and player data in real time.

---

## 📸 Preview

| ![HLL-Observer-1](https://github.com/user-attachments/assets/552786c2-6c72-4061-914a-44a07bc1c5e1) | ![HLL-Observer-2](https://github.com/user-attachments/assets/d3643146-d35c-48ea-ac34-2dd43dcf8519) |
|:--------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------|

---

## 🚀 Features

- **Real-time player information**: Displays detailed player stats, including name, role, team, and more.
- **Server monitoring**: View server name, map details, and player counts.
- **Interactive map view**: Zoom, pan, and toggle grid overlays.
- **Customizable UI**: Easily extendable for additional features.

---

## 📦 Installation

### Install the Go package

To install the Go package, simply run:

```bash
go install github.com/zMoooooritz/go-let-observer@latest
```

or download from [releases](https://github.com/zMoooooritz/go-let-observer/releases)

### Build the project

Clone the repository and build the project:

```bash
git clone https://github.com/zMoooooritz/go-let-observer.git
cd go-let-observer
go build .
```

---

## 📖 Usage

### Running the Observer

To start the observer, run the `go-let-observer` command:

```bash
Usage of go-let-observer:
  -host string
        RCON server host
  -password string
        RCON server password
  -port string
        RCON server port
  -size int
        Screen size (default 1000)
  -version
        Display version
```

---

## 🎮 Keybinds

| Keybind              | Description                        |
|----------------------|------------------------------------|
| `LeftClick`          | Show/Hide player info              |
| `RightClick-Drag`    | Move the map view                  |
| `MouseWheel`         | Zoom in/out on the map             |
| `+` / `-`            | Increase/Decrease update interval  |
| `P`                  | Toggle players                     |
| `S`                  | Toggle guesstimated spawns         |
| `G`                  | Toggle grid overlay                |
| `H`                  | Toggle header overlay              |
| `Tab`                | Show scoreboard                    |
| `Esc`, `Q`, `Ctrl+C` | Exit the application               |

---

## 🔧 Built With

- [Ebiten](https://github.com/hajimehoshi/ebiten) - 2D game library for Go
- [go-let-loose](https://github.com/zMoooooritz/go-let-loose) - HLL API written in Go

---

## 📄 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
