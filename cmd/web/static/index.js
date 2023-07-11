const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
  go.run(result.instance);

  // golangで実装したhelloが利用可能
  document.body.innerHTML = hello();
});
