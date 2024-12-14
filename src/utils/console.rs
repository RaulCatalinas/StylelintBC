use crate::constants::{DEFAULT_COLOR, MESSAGE_COLORS};
use crate::enums;

pub struct WriteMessageProps {
    pub message_type: enums::MessageType,
    pub message: &'static str,
}

pub fn write_message(props: WriteMessageProps) {
    let default_color_string = DEFAULT_COLOR.to_string();

    let color = MESSAGE_COLORS
        .get(&props.message_type)
        .unwrap_or(&default_color_string);

    println!("{}{}{}", color, props.message, DEFAULT_COLOR);
}
