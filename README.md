# HLL Observer

[![Latest Release](https://img.shields.io/github/release/zMoooooritz/go-let-observer.svg?style=for-the-badge)](https://github.com/zMoooooritz/go-let-observer/releases)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/zMoooooritz/go-let-observer)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/zMoooooritz/go-let-observer/build.yml?branch=master&style=for-the-badge)](https://github.com/zMoooooritz/nachrichten/actions)
[![Go ReportCard](https://goreportcard.com/badge/github.com/zMoooooritz/go-let-observer?style=for-the-badge)](https://goreportcard.com/report/zMoooooritz/go-let-observer)

An **observer tool** for the game **Hell Let Loose**, providing a graphical interface to monitor server and player data in real time.

---

## 📸 Preview

| ![HLL-Observer-1](https://github.com/user-attachments/assets/552786c2-6c72-4061-914a-44a07bc1c5e1) | ![HLL-Observer-2](https://github.com/user-attachments/assets/d8e1c97f-26f2-4baa-af11-a1647d4ef3e2) |
|:--------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------|

---

## 🚀 Features

- **Real-time player information**: Displays detailed player stats, including name, role, team, and more.
- **Server monitoring**: View server name, map details, and player counts.
- **Interactive map view**: Zoom, pan, and toggle overlays.
- **Recording and Replaying**: Record gameplay or actions and replay them for analysis

---

## 📦 Installation

### Install the Go package

To install the Go package, simply run:

```bash
go install github.com/zMoooooritz/go-let-observer@latest
```

or download from [releases](https://github.com/zMoooooritz/go-let-observer/releases).

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
  -config string
        Path to configuration file
  -directory string
        Path to the replays directory
  -host string
        RCON server host
  -mode string
        Mode to run on startup (viewer, replay, record)
  -password string
        RCON server password
  -port string
        RCON server port
  -size int
        Screen size (default 1000)
  -version
        Display version information
```

---

## 🎮 Keybinds

The following keybinds are configured and can be used:

| Keybind              | Description                        |
|----------------------|------------------------------------|
| `+` / `-`            | Increase/Decrease update interval  |
| `P`                  | Toggle players                     |
| `I`                  | Toggle player info                 |
| `S`                  | Toggle guesstimated spawns         |
| `T`                  | Toggle tanks                       |
| `G`                  | Toggle grid overlay                |
| `H`                  | Toggle header overlay              |
| `C`                  | Configure sectors (1-3, 0=skip)    |
| `Tab`                | Show scoreboard                    |
| `Space`              | Toggle replay pause                |
| `ArrowRight`         | Seek forward in replay             |
| `ArrowLeft`          | Seek backward in replay            |
| `Esc`, `Q`, `Ctrl+C` | Exit the application               |

| Mouse action         | Description                        |
|----------------------|------------------------------------|
| `LeftClick`          | Select a player                    |
| `RightClick-Drag`    | Pan the map                        |
| `MouseWheel`         | Zoom the map                       |

---

## 🔧 Built With

- [Ebiten](https://github.com/hajimehoshi/ebiten) - 2D game library for Go
- [go-let-loose](https://github.com/zMoooooritz/go-let-loose) - HLL API written in Go

---

## 📄 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
