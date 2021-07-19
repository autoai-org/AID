#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]
#![allow(
    // Clippy bug: https://github.com/rust-lang/rust-clippy/issues/7422
    clippy::nonstandard_macro_braces,
)]

use tauri::{
  api::process::{Command},
  Manager,
};

fn main() {
  tauri::Builder::default()
    .setup(|app| {
      let _window = app.get_window("main").unwrap();

      tauri::async_runtime::spawn(async move {
        let (mut _rx, mut _child) = Command::new("aid")
          .args(vec!["up"])
          .spawn()
          .expect("failed to setup `aid up` sidecar");
      });
      Ok(())
    })
    .run(tauri::generate_context!())
    .expect("error while running tauri application");
}
