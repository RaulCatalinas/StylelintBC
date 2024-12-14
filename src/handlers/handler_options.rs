use crate::constants;
use crate::enums;
use crate::utils;

use std::process;

pub fn handler_option_collaborate() {
    utils::write_message(utils::WriteMessageProps {
        message_type: enums::MessageType::Info,
        message: "Opening the GitHub repository...",
    });

    let url_opened = utils::open_url(constants::REPOSITORY);

    if !url_opened {
        utils::write_message(utils::WriteMessageProps {
            message_type: enums::MessageType::Error,
            message: "Failed to open the GitHub repository",
        });

        process::exit(1);
    }
}

pub fn handler_option_build() {
    println!("Generating Stylelint configuration");
}
