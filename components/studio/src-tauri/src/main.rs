#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

mod cmd;

fn main() {
  tauri::AppBuilder::new()
    .invoke_handler(|_webview, arg| {
      std::thread::spawn(move || {
        spawn_aid_server();
      });
      use cmd::Cmd::*;
      match serde_json::from_str(arg) {
        Err(e) => Err(e.to_string()),
        Ok(command) => {
          match command {
            // definitions for your custom commands from Cmd here
            MyCustomCommand { argument } => {
              //  your command code
              println!("{}", argument);
            }
          }
          Ok(())
        }
      }
    })
    .build()
    .run();
}
fn spawn_aid_server() {
  let (mut rx, mut child) = Command::new_sidecar("my-sidecar")
    .expect("failed to create `my-sidecar` binary command")
    .spawn()
    .expect("Failed to spawn sidecar");

  tauri::async_runtime::spawn(async move {
    // read events such as stdout
    while let Some(event) = rx.recv().await {
      if let CommandEvent::Stdout(line) = event {
        window
          .emit("message", Some(format!("'{}'", line)))
          .expect("failed to emit event");
        // write to stdin
        child.write("message from Rust\n".as_bytes()).unwrap();
      }
    }
  });
  Copy
}
