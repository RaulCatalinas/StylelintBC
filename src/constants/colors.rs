use crate::enums;
use std::collections::HashMap;

const COLOR_PREFIX: &str = "\033[";
const COLOR_SUFFIX: &str = "m";

pub const DEFAULT_COLOR: &str = COLOR_PREFIX + enums::MessageColor::Default + COLOR_SUFFIX;

pub static MESSAGE_COLORS: HashMap<enums::MessageType, String> = {
    let mut colors: HashMap<enums::MessageType, String> = HashMap::new();

    colors.insert(
        enums::MessageType::Success,
        COLOR_PREFIX + enums::MessageColor::Green + COLOR_SUFFIX,
    );
    colors.insert(
        enums::MessageType::Error,
        COLOR_PREFIX + enums::MessageColor::Red + COLOR_SUFFIX,
    );
    colors.insert(
        enums::MessageType::Warning,
        COLOR_PREFIX + enums::MessageColor::Yellow + COLOR_SUFFIX,
    );
    colors.insert(
        enums::MessageType::Info,
        COLOR_PREFIX + enums::MessageColor::Blue + COLOR_SUFFIX,
    );

    return colors;
};
