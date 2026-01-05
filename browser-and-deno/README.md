# Use go-jsonnet to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

Tested with Go 1.25.5, Bun 1.3.5, Deno 2.6.3, go-jsonnet [v0.21.0](https://github.com/google/go-jsonnet/tree/v0.21.0).

> [!NOTE]
> This example uses Jsonnet as an expression language. 
> 
> If your use case is employing Jsonnet as configuration language,
> take in account the [hermeticity](https://github.com/google/jsonnet/blob/v0.21.0/doc/articles/comparisons.html#L189),
> there external calls like HTTP requests are discouraged.

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-and-deno
```

2. Install go-jsonnet:

```sh
go get github.com/google/go-jsonnet
```

3. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

4. Copy the glue JS from Golang distribution to example's folder (note using `/lib/wasm/` because Go 1.24+):

```sh
cp $(go env GOROOT)/lib/wasm/wasm_exec.js ./
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
~/.deno/bin/deno run --allow-net --allow-read jsr:@std/http/file-server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As `index.html` and a **15M**-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Test with Node.js

Impossible yet due to https://github.com/golang/go/issues/59605.

### Test with Bun

1. Install Bun:

```sh
curl -fsSL https://bun.sh/install | bash
```

2. Run with Bun:

```sh
~/.bun/bin/bun bun.js
```

### Test with Deno

1. Run with Deno:

```sh
~/.deno/bin/deno run --allow-read --allow-net deno.js
```

### Finish

Perform your own experiments if desired.
