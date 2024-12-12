pub struct CliOption {
    pub name: &'static str,
    pub alias: &'static str,
    pub description: &'static str,
    pub handler: fn(),
}
