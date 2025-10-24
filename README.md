# 🧭 todo — A Simple Local CLI Task Tracker (written in Go)

**todo** is a lightweight, local-only task and work tracker designed for people who want to stay focused — not buried under complex project management tools.

It stores your data as easy-to-read **JSON files** on disk, so you always own your notes, timelines, and history.

---

## ✨ Features

- 📝 Add, list, and manage tasks quickly from the command line
- ⏱️ Start and stop work sessions to track what you _actually_ worked on
- 🗂️ Tags, areas, and priorities to organize tasks
- 💬 Add notes and comments as you go
- 📦 All data stored locally as `JSON` — easy to back up or export
- 🧾 Export summaries or text dumps (great for journaling or LLMs)
- ⚡ No servers, no sync, no accounts — **just your data**

---

## ⚙️ Installation

You’ll need **Go 1.22+** installed.  
Then clone and build the binary:

```bash
git clone https://github.com/YOURUSERNAME/todo.git
cd todo
go install ./...
```

This will install the `todo` binary into your `$GOPATH/bin` (usually `~/go/bin`).

Make sure that folder is on your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Now you can run:

```bash
todo --help
```

---

## 🚀 How to Use

Here’s how you might use `todo` during a real day of work.

### 🆕 Add new tasks

```bash
todo add "Write project proposal" --body "Draft the client proposal" --priority 1 --area work
todo add "Clean up Go code" --area dev
```

### 📋 List what’s on your plate

```bash
todo list
```

Output:

```
1. [todo] Write project proposal (Priority 1)
2. [todo] Clean up Go code (Priority 2)
```

### ▶️ Start work

Start tracking time automatically:

```bash
todo start "Write project proposal"
```

### 🔄 Switch when something urgent drops in

```bash
todo switch "Fix login bug"
```

This automatically stops the previous session and starts a new one.

### ⏹️ Stop when you’re done

```bash
todo stop
```

### 🏁 Mark a task as complete

```bash
todo done 1
```

### 🗒️ Add a note or comment

```bash
todo note 2 "Need to refactor the DB layer"
todo comment 2 "Reviewed PR #45 and merged"
```

### 📊 Review your day or month

```bash
todo summary --today
todo summary --month oct
```

Output example:

```
# Work Summary (Oct 23)
09:10 - 10:30  Write project proposal
10:35 - 11:50  Fix login bug
13:00 - 14:00  Write project proposal
```

### 📤 Export for journaling or AI tools

```bash
todo export --markdown > summary.md
```

---

## 🧠 Philosophy

`todo` is designed to be:

- **Fast** — minimal friction between idea and action
- **Local-first** — you own your data
- **Honest** — helps you see where your time really went
- **Extensible** — JSON storage makes it easy to integrate with other tools

It’s not a replacement for heavy project management software — it’s a personal tool for focus and reflection.

---

## 🧩 Roadmap

- [ ] Add a full-screen **TUI interface** (like lazygit)
- [ ] Add search and filter commands
- [ ] Auto-close long-running sessions
- [ ] Optional time summaries by area or tag
- [ ] Sync/export integrations (maybe)

---

## 🤝 Contributing

Pull requests, issues, and ideas are very welcome!

If you want to contribute:

1. Fork the repo
2. Create a feature branch (`git checkout -b feature/my-idea`)
3. Submit a pull request

Even small feedback or docs improvements help a lot.

---

## 🌟 Like this project?

If you find it useful —  
please **★ star the repository** on GitHub!  
It helps others discover it and keeps motivation high 🙌

---

## 📂 Data Location

Your data is stored locally at:

```
~/.localtodo/tasks.json
~/.localtodo/current.json
```

You can back up or inspect these files at any time.

---

## 🧑‍💻 License

MIT License — do whatever you like, just don’t blame me if it breaks your day 😉  
See [LICENSE](LICENSE) for details.

---

Built with Go, calm focus, and too much coffee ☕
