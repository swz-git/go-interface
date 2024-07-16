# rlbot/go-interface

Go interface for the RLBot v5 socket api

## Todo

* Nicer way to handle flatbuffers? (ideally we dont want the T suffix for the object api)
  * Wrappers? (much manual work)
  * Codegen? (rely upon a lot of custom codegen, eww)
* Helper stuff (niceties)
  * Rendering
  * Statesetting?
  * Basic Agent

Also see TODO comments

## Building

If you're on linux and have [just](https://github.com/casey/just) installed, do `just build`. If not, see `./justfile`.
