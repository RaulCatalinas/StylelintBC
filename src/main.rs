pub mod cli;
pub mod config;
pub mod constants;
pub mod enums;
pub mod handlers;
pub mod options;
pub mod types;
pub mod utils;

fn main() {
    config::configure_options();
}
