use webbrowser;

pub fn open_url(url: &str) -> bool {
    return webbrowser::open(url).is_ok();
}
