# commit

🌴 Quickly commits code

## 🚀 Installation

### Via Homebrew

Add my tap by running

```bash
brew tap abroudoux/tap
```

You can now download `commit

```bash
brew install abroudoux/tap/commit
```

Enjoy!

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x commit
```

## 💻 Usage

`commit` allows you to commit your code by running a unique command.

```bash
commit
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request with the details of your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [x] Fix set-upstream creation when first commit
- [ ] Choose upstream during commit
- [x] Rewrite in Go
- [ ] More options during `git add` step
- [x] Installation via Homebrew

## 📑 License

This project is under [MIT License](LICENSE).
