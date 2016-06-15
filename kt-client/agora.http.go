package main

import (
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/agora/runtime"
)

// HTTPMod is an agora NativeModule implementing HTTP(S).
type HTTPMod struct {
	ctx *runtime.Ctx
	obj runtime.Object
}

// ID returns the module's name.
func (mod *HTTPMod) ID() string {
	return "http"
}

// Run returns the module's object.
func (mod *HTTPMod) Run(_ ...runtime.Val) (v runtime.Val, err error) {
	defer runtime.PanicToError(&err)
	if mod.obj == nil {
		mod.obj = runtime.NewObject()
		mod.obj.Set(runtime.String("GET"), runtime.NewNativeFunc(mod.ctx, "mymod.GET", mod.get))
	}
	return mod.obj, nil
}

// SetCtx sets the module's context.
func (mod *HTTPMod) SetCtx(ctx *runtime.Ctx) {
	mod.ctx = ctx
}

func (mod *HTTPMod) get(args ...runtime.Val) runtime.Val {
	runtime.ExpectAtLeastNArgs(1, args)
	url, ok := args[0].(runtime.String)
	if !ok {
		panic("Invalid argument")
	}

	res, err := http.Get(string(url))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return runtime.String(b)
}
