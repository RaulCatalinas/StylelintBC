use crate::types;
use std::fmt::Write;

pub fn show_help(options: Vec<types::CliOption>) {
    let mut builder = String::new();

    builder.push_str("Usage: stylelintbc [options]\n\n");
    builder.push_str("Command line for easy Stylelint configuration\n\n");
    builder.push_str("Options:\n");

    for option in options {
        let _ = write!(
            builder,
            "{:<15} {:<5} {}\n",
            option.name, option.alias, option.description
        );
    }

    builder.push('\n');

    print!("{}", builder);
}
