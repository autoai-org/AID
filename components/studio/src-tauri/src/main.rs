#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]
#![allow(
    // Clippy bug: https://github.com/rust-lang/rust-clippy/issues/7422
    clippy::nonstandard_macro_braces,
)]

#[cfg(not(feature = "ui"))]
mod rust {
  use tauri::{api::process::Command, Manager};

  #[tauri::command]
  fn close_splashscreen() {}

  pub fn main() {
    tauri::Builder::default()
      .setup(|app| {
        let splashscreen_window = app.get_window("splashscreen").unwrap();
        let _window = app.get_window("main").unwrap();
        tauri::async_runtime::spawn(async move {
          let (mut _rx, mut _child) = Command::new("aid")
            .args(vec!["up"])
            .spawn()
            .expect("failed to setup `aid up` sidecar");
          splashscreen_window.close().unwrap();
          _window.show().unwrap()
        });
        Ok(())
      })
      .invoke_handler(tauri::generate_handler![close_splashscreen])
      .run(tauri::generate_context!())
      .expect("error while running tauri application");
  }
}

#[cfg(feature = "ui")]
mod ui {
  use std::sync::{Arc, Mutex};
  use tauri::{Manager, State, Window};

  struct SplashscreenWindow(Arc<Mutex<Window>>);
  struct MainWindow(Arc<Mutex<Window>>);

  #[tauri::command]
  fn close_splashscreen(
    _: Window, // force inference of P
    splashscreen: State<SplashscreenWindow>,
    main: State<MainWindow>,
  ) {
    splashscreen.0.lock().unwrap().close().unwrap();
    main.0.lock().unwrap().show().unwrap();
  }
  pub fn main() {
    tauri::Builder::default()
      .setup(|app| {
        // set the splashscreen and main windows to be globally available with the tauri state API
        app.manage(SplashscreenWindow(Arc::new(Mutex::new(
          app.get_window("splashscreen").unwrap(),
        ))));
        app.manage(MainWindow(Arc::new(Mutex::new(
          app.get_window("main").unwrap(),
        ))));
        Ok(())
      })
      .invoke_handler(tauri::generate_handler![close_splashscreen])
      .run(tauri::generate_context!())
      .expect("error while running aid desktop");
  }
}

fn main() {
  #[cfg(feature = "ui")]
  ui::main();
  #[cfg(not(feature = "ui"))]
  rust::main();
}
