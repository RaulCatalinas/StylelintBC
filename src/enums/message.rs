#[derive(Hash, PartialEq, Eq)]
pub enum MessageType {
    Success,
    Error,
    Warning,
    Info,
    Config,
}

pub enum MessageColor {
    Green = 32,
    Blue = 36,
    Red = 31,
    White = 37,
    Yellow = 33,
    Default = 0,
}
