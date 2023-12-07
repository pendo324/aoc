### Usage

Create a config.toml file and add your aoc_session_cookie, i.e.

```toml
aoc_session_cookie = "my_cookie"
```

`make` and then `aoc create-day <year> <day>` to download a day. Re-run to download the second part of the question as markdown after you've unlocked it.

Use `aoc run year day` to run the relevant `main.go` file for that day.

#### Note

This project also comes with a VSCode `launch.json` which allows you to press F5 and then execute any of the aforementioned commands with a debugger.
