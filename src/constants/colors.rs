use crate::enums;
use lazy_static::lazy_static;
use std::collections::HashMap;

const COLOR_PREFIX: &str = "\033[";
const COLOR_SUFFIX: &str = "m";
pub const DEFAULT_COLOR: &str = "\033[0m";

lazy_static! {
    pub static ref MESSAGE_COLORS: HashMap<enums::MessageType, String> = {
        let mut map = HashMap::new();

        map.insert(
            enums::MessageType::Success,
            format!(
                "{}{}{}",
                COLOR_PREFIX,
                enums::MessageColor::Green as u8,
                COLOR_SUFFIX
            ),
        );

        map.insert(
            enums::MessageType::Info,
            format!(
                "{}{}{}",
                COLOR_PREFIX,
                enums::MessageColor::Blue as u8,
                COLOR_SUFFIX
            ),
        );

        map.insert(
            enums::MessageType::Error,
            format!(
                "{}{}{}",
                COLOR_PREFIX,
                enums::MessageColor::Red as u8,
                COLOR_SUFFIX
            ),
        );

        map.insert(
            enums::MessageType::Config,
            format!(
                "{}{}{}",
                COLOR_PREFIX,
                enums::MessageColor::White as u8,
                COLOR_SUFFIX
            ),
        );

        map.insert(
            enums::MessageType::Warning,
            format!(
                "{}{}{}",
                COLOR_PREFIX,
                enums::MessageColor::Yellow as u8,
                COLOR_SUFFIX
            ),
        );
        map
    };
}
