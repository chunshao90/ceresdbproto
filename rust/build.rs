// Copyright 2022 CeresDB Project Authors. Licensed under Apache-2.0.

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Download the protoc and set the path for tonic_build.
    let protoc_path = protoc_bin_vendored::protoc_bin_path().map_err(|e| Box::new(e))?;
    std::env::set_var("PROTOC", protoc_path.as_os_str());

    tonic_build::configure().out_dir("./src").compile(
        &[
            "../protos/cluster.proto",
            "../protos/xx/common.proto",
            "../protos/meta_event.proto",
            "../protos/meta_service.proto",
            "../protos/meta_storage.proto",
            "../protos/prometheus.proto",
            "../protos/storage.proto",
        ],
        &["../protos"],
    )?;
    Ok(())
}
