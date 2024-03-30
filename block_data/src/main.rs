use reqwest::StatusCode;
use serde::{Deserialize, Serialize};
 
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Confidence {
    pub block: u32,
    pub confidence: f64,
    pub serialised_confidence: Option<String>,
}
 
const LIGHT_CLIENT_URL: &str = "http://127.0.0.1:7000";
 
#[tokio::main]
async fn main() {
    let block_number = 586959;
    let confidence_url = format!("{LIGHT_CLIENT_URL}/v1/confidence/{block_number}");
    let response = reqwest::get(&confidence_url).await.unwrap();
 
    if response.status() == StatusCode::OK {
        let confidence: Confidence =
            serde_json::from_str(&response.text().await.unwrap()).unwrap();
        println!("{confidence:?}");
    } else {
        eprintln!("Failed to get confidence: {}", response.status());
    }
    // ...error handling...
}