use std::sync::LazyLock;

pub const REPOSITORY: &str = "https://github.com/RaulCatalinas/StylelintBC";
pub static ISSUES: LazyLock<String> = LazyLock::new(|| REPOSITORY.to_owned() + "/issues");
