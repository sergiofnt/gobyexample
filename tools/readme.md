## Build Tools

### `format`
Bash scripts that runs gofmt and replace tabs with spaces (to save space in generated) html files on all *.go files.

```bash
tools/format
```

### `build`
Small bash scripts that runs `foramt` and both builders. Will regenerate full site.

```bash
tools/build
```

### `build-Loop`
Runs `updated` every 10 seconds.

```bash
tools/build-loop
```

### `updated`
`updated` is a copy of `format` that also check sha1 (on `*.go` file )before formatting and generation.

### `measure`
soon

### `generate`
soon

