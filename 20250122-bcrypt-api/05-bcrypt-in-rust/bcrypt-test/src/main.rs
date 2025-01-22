use rand::RngCore;
use base64::{Engine as _, engine::general_purpose::URL_SAFE};
use std::error::Error;

fn random_string(length: usize) -> String {
    let mut bytes = vec![0u8; length];
    rand::thread_rng().fill_bytes(&mut bytes);
    URL_SAFE.encode(&bytes)[..length].to_string()
}

fn main() -> Result<(), Box<dyn Error>> {
    // 18 + 55 + 1 = 74, so above 72 characters' limit of BCrypt
    let user_id = random_string(18);
    let username = random_string(55);
    let password = "super-duper-secure-password";

    let combined_string = format!("{}:{}:{}", user_id, username, password);
    let combined_hash = bcrypt::hash(combined_string.as_bytes(), bcrypt::DEFAULT_COST)?;

    // let's try to break it
    let wrong_password = "wrong-password";
    let wrong_combined_string = format!("{}:{}:{}", user_id, username, wrong_password);

    match bcrypt::verify(wrong_combined_string.as_bytes(), &combined_hash) {
        Ok(true) => println!("Password is correct"),
        Ok(false) => println!("Password is incorrect"),
        Err(e) => println!("{}", e),
    }

    Ok(())
}