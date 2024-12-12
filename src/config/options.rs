use std::env;
use std::process;

use crate::cli::show_help;
use crate::constants::OPTIONS;
use crate::enums::MessageType;
use crate::options;
use crate::utils::{write_message, WriteMessageProps};

pub fn configure_options() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 2 {
        show_help(OPTIONS);

        process::exit(0);
    }

    let input = &args[1];

    let options = options::get_options();

    for option in options {
        if input == option.name || input == option.alias {
            (option.handler)();

            process::exit(0);
        }
    }

    write_message(WriteMessageProps {
        message_type: MessageType::Error,
        message: "The option you've tried to execute doesn't exist",
    });

    println!();

    show_help(OPTIONS);
}
