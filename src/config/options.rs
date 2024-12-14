use std::env;
use std::process;

use crate::cli::show_help;
use crate::enums::MessageType;
use crate::options::get_options;
use crate::utils::{write_message, WriteMessageProps};

pub fn configure_options() {
    let args: Vec<String> = env::args().collect();
    let options = get_options();

    if args.len() != 2 {
        show_help(options);

        process::exit(0);
    }

    let input = &args[1];

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

    let options = get_options();

    show_help(options);
}
