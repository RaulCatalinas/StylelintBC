use crate::constants;
use crate::utils;

pub fn handler_option_collaborate() {
    utils::open_url(constants::REPOSITORY);
}

pub fn handler_option_build() {
    println!("Generating Stylelint configuration");
}
