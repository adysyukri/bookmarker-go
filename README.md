# bookmarker-go

[Browser bookmark](https://github.com/adysyukri/browser-bookmarking) project but with Go and backend DB instead of localStorage

---

## Run development server

If using Nix, in the root folder, simply run

```
nix develop
```

After success, run `make run`

If not using Nix, install these:

- [Go](https://go.dev/doc/install) (1.22)
- [Node Js](https://nodejs.org/en/download/package-manager) (21)

Make sure to have Go bin folder in PATH variable.

In terminal at {home} directory, type

```bash
nano .bashrc
```

Then add at the bottom line of the file and save

```bash
export PATH=$PATH:/home/{your username}/go/bin
```

Install go libs dependancies:

- Install **air**

```
go install github.com/cosmtrek/air@latest
```

- Install **templ**

```
go install github.com/a-h/templ/cmd/templ@latest
```

Copy `.air.toml.example` file to `.air.toml`

Then simply run `make run`

> Server by default will run on port 3000
