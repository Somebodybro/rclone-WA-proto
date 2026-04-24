// Rclone as a wasm library
//
// This library exports the core rc functionality

//go:build js

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"syscall/js"

	"github.com/Somebodybro/rclone-WA-proto/fs"
	"github.com/Somebodybro/rclone-WA-proto/fs/rc"

	// Core functionality we need
	_ "github.com/Somebodybro/rclone-WA-proto/fs/operations"
	_ "github.com/Somebodybro/rclone-WA-proto/fs/sync"
	_ "github.com/Somebodybro/rclone-WA-proto/librclone/librclone"

	// _ "github.com/Somebodybro/rclone-WA-proto/backend/all" // import all backends

	// Backends
	_ "github.com/Somebodybro/rclone-WA-proto/backend/alias"
	_ "github.com/Somebodybro/rclone-WA-proto/backend/drive"
	_ "github.com/Somebodybro/rclone-WA-proto/backend/ftp"
	_ "github.com/Somebodybro/rclone-WA-proto/backend/http"
	_ "github.com/Somebodybro/rclone-WA-proto/backend/local"
	_ "github.com/Somebodybro/rclone-WA-proto/backend/memory"
)

var (
	document js.Value
	jsJSON   js.Value
)

// errorValue turns an error into a js.Value
func errorValue(method string, in js.Value, err error) js.Value {
	fs.Errorf(nil, "rc: %q: error: %v", method, err)
	// Adjust the error return for some well known errors
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, fs.ErrorDirNotFound) || errors.Is(err, fs.ErrorObjectNotFound):
		status = http.StatusNotFound
	case rc.IsErrParamInvalid(err) || rc.IsErrParamNotFound(err):
		status = http.StatusBadRequest
	}
	return js.ValueOf(map[string]interface{}{
		"status": status,
		"error":  err.Error(),
		"input":  in,
		"path":   method,
	})
}

// rcCallback is a callback for javascript to access the api
func rcCallback(this js.Value, args []js.Value) interface{} {
	method := args[0].String()
	inRaw := args[1]

	if len(args) != 2 {
		return errorValue("", js.Undefined(), errors.New("need two parameters to rc call"))
	}

	handler := js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]
		go func() {
			ctx := context.Background() // FIXME
			log.Printf("rcCallback: this=%v args=%v", this, args)

			var in = rc.Params{}
			switch inRaw.Type() {
			case js.TypeNull:
			case js.TypeObject:
				inJSON := jsJSON.Call("stringify", inRaw).String()
				err := json.Unmarshal([]byte(inJSON), &in)
				if err != nil {
					reject.Invoke(errorValue(method, inRaw, fmt.Errorf("couldn't unmarshal input: %w", err)))
					return
				}
			default:
				reject.Invoke(errorValue(method, inRaw, errors.New("in parameter must be null or object")))
				return
			}

			call := rc.Calls.Get(method)
			if call == nil {
				reject.Invoke(errorValue(method, inRaw, fmt.Errorf("method %q not found", method)))
				return
			}

			out, err := call.Fn(ctx, in)
			if err != nil {
				reject.Invoke(errorValue(method, inRaw, fmt.Errorf("method call failed: %w", err)))
				return
			}
			if out == nil {
				log.Print("resolve nil")
				resolve.Invoke(nil)
				return
			}
			var out2 map[string]interface{}
			err = rc.Reshape(&out2, out)
			if err != nil {
				reject.Invoke(errorValue(method, inRaw, fmt.Errorf("result reshape failed: %w", err)))
				return
			}
			resolve.Invoke(out2)
		}()
		return nil
	})
	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func main() {
	log.Printf("Running on goos/goarch = %s/%s", runtime.GOOS, runtime.GOARCH)
	if js.Global().IsUndefined() {
		log.Fatalf("Didn't find Global - not running in browser")
	}
	document = js.Global().Get("document")
	if document.IsUndefined() {
		log.Fatalf("Didn't find document - not running in browser")
	}

	jsJSON = js.Global().Get("JSON")
	if jsJSON.IsUndefined() {
		log.Fatalf("can't find JSON")
	}

	// Set rc
	js.Global().Set("rc", js.FuncOf(rcCallback))

	// Signal that it is valid
	rcValidResolve := js.Global().Get("rcValidResolve")
	if rcValidResolve.IsUndefined() {
		log.Fatalf("Didn't find rcValidResolve")
	}
	rcValidResolve.Invoke()

	// Wait forever
	select {}
}
