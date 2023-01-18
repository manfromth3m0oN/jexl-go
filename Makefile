.PHONY: build-static
build-static:
	@cd lib/jexl-ffi && cargo build --release
	@cp lib/jexl-ffi/target/release/libjexl_ffi.a lib/libjexl.a
	go build