use crate::cli;
use crate::handlers;
use crate::types;

pub fn get_options() -> Vec<types::CliOption> {
    return vec![
        types::CliOption {
            name: "--version",
            alias: "-v",
            description: "Output the version number",
            handler: cli::show_version,
        },
        types::CliOption {
            name: "--collaborate",
            alias: "-co",
            description: "Open GitHub repository for collaboration",
            handler: handlers::handler_option_collaborate,
        },
        types::CliOption {
            name: "--build",
            alias: "-b",
            description: "Start Stylelint's configuration",
            handler: handlers::handler_option_build,
        },
        types::CliOption {
            name: "--help",
            alias: "-h",
            description: "Display help for command",
            handler: || {
                cli::show_help(get_options());
            },
        },
    ];
}
