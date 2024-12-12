pub mod cli;
pub mod config;
pub mod constants;
pub mod enums;
pub mod handlers;
pub mod options;
pub mod types;
pub mod utils;

fn main() {
    utils::write_message(enums::MessageType::Success, "StylelintBC started");
    utils::write_message(enums::MessageType::Error, "StylelintBC stopped");
    utils::write_message(enums::MessageType::Warning, "StylelintBC warning");
    utils::write_message(enums::MessageType::Info, "StylelintBC info");
}
