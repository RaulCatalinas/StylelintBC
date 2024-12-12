use crate::constants::{DEFAULT_COLOR, MESSAGE_COLORS};
use crate::enums;

pub fn write_message(message_type: enums::MessageType, message: &str) {
    let color = MESSAGE_COLORS.get(message_type).unwrap_or(DEFAULT_COLOR);

    println!("{}{}{}", color, message, DEFAULT_COLOR);
}
