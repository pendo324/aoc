### Usage

Create a .env file and add your AOC_SESSION_COOKIE.

`npm i` then `npm run create-day [year] <day>` to download a day. Re-run to download the second part of the question as markdown after you've unlocked it.

Use `npm run <year>-<day>` to run the relevant `index.js` file for that day.

#### Note

On windows, for some unknown reason, `npm i` will fail if you aren't logged in. Use `npm login` before running `npm i`.

### Maintenance

When this repo was setup, npm had _just_ added multi-project workspaces and ts-node had _just_ added support for Node's native ESModules. Hopefully, eventually, some hacks can be removed, namely:

- the `create-day` npm script flags in `./utils/package.json` (`--no-warnings --experimental-specifier-resolution=node --loader ts-node/esm`)
- the `npm login` warning in the readme?? (still very confused about that one. maybe it will "go away" when `ts-node` is further updated)
