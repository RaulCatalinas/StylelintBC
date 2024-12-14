pub mod cli;
pub mod config;
pub mod constants;
pub mod enums;
pub mod handlers;
pub mod options;
pub mod types;
pub mod utils;

fn main() {
    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Success,
        message: "StylelintBC started",
    });

    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Error,
        message: "StylelintBC stopped",
    });

    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Info,
        message: "StylelintBC info",
    });

    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Warning,
        message: "StylelintBC warning",
    });

    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Config,
        message: "StylelintBC config",
    });
}
