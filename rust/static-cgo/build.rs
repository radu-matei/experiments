fn main() {
    println!("cargo:rustc-link-search=native={}", "./target");
    println!("cargo:rustc-link-lib=static={}", "wasm2oci");

    #[cfg(target_os = "macos")]
    {
        println!("cargo:rustc-flags=-l framework=CoreFoundation -l framework=Security");
    }
}
