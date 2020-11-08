use tokio;
use std::process::exit;
use reqwest::{Result};
use serde::{Deserialize};
use chrono::{DateTime, Utc, Timelike};
use std::{time, thread};

const ERROR_INVALID_ENV_VAR: i32 = 5;
const ERROR_MISSING_ENV_VAR: i32 = 10;

#[derive(Deserialize, Debug)]
struct Alarm {
    pub _id: String,
    pub hour: i32,
    pub minute: i32,
    pub days: Vec<String>
}

struct Triggered {
    pub alarm_id: String,
    pub last_triggered: DateTime<Utc>
}

impl Triggered {
    fn set_id(&mut self, alarm_id: String) {
        self.alarm_id = alarm_id;
    }
    fn set_last_triggered(&mut self, last_triggered: DateTime<Utc>) {
        self.last_triggered = last_triggered;
    }
}

#[tokio::main]
async fn main() -> Result<()> {
    let alarm_url_result = get_env("COFFEE_ALARM_API");
    let auth_code_result = get_env("COFFEE_AUTH_CODE");

    exit_on_invalid_env_var(&alarm_url_result);
    exit_on_invalid_env_var(&auth_code_result);

    let mut triggered: Vec<Triggered> = Vec::new();

    loop {
        let latest_alarms = get_alarms(
            &alarm_url_result.0, &auth_code_result.0).await?;
        let current_date = chrono::Utc::now();

        for alarm in latest_alarms {
            if alarm.hour as u32 == current_date.hour() && alarm.minute as u32 == current_date.minute() {
                println!("triggered");
            }
            println!("{:?}", current_date);
            thread::sleep(time::Duration::from_secs(25));
        }
    }
}

async fn get_alarms(url: &str, code: &str) -> Result<Vec<Alarm>>{
    let client = reqwest::Client::new();
    let response = client
        .get(url)
        .header("Authorization", format!("Bearer {}", code))
        .send().await?;

    if response.status() == 200 {
        let alarms: Vec<Alarm> = response.json().await?;
        return Result::Ok(alarms);
    }
    return Result::Ok(Vec::new());
}

fn exit_on_invalid_env_var(env_var_result: &(String, i32)) {
    if env_var_result.1 != 0 {
        println!("{}", env_var_result.0);
        exit(env_var_result.1);
    }
}

/// retrieves the value stored in environment variable "COFFEE_AUTH"
/// Warning: this function exits the program if the value was not found or
///     cannot be converted to a string
fn get_env(env_name: &str) -> (String, i32) {
    match std::env::var_os(env_name) {
        Some(auth_os_string) => {
            match auth_os_string.to_str() {
                Some(auth_string) => {
                    return (auth_string.to_string(), 0);
                }
                _ => {
                    return (format!("value of env variable {} invalid", env_name).to_string(), ERROR_INVALID_ENV_VAR);
                }
            }
        }
        _ => {
            return (format!("env variable {} is missing", env_name).to_string(), ERROR_MISSING_ENV_VAR);
        }
    }
}
