use tokio;
use std::process::exit;
use reqwest::{Result};
use serde::{Deserialize};
use chrono::{DateTime, Local, Timelike};
use std::{time, thread};
use gpio::GpioOut;

const ERROR_INVALID_ENV_VAR: i32 = 5;
const ERROR_MISSING_ENV_VAR: i32 = 10;
const ERROR_GPIO_NOT_AN_INT: i32 = 15;

#[derive(Deserialize, Debug)]
struct Alarm {
    pub _id: String,
    pub hour: i32,
    pub minute: i32,
    pub days: Vec<String>,
}

#[derive(Clone)]
struct Triggered {
    pub alarm_id: String,
    pub last_triggered: DateTime<Local>,
}

impl Triggered {
    // fn set_id(&mut self, alarm_id: String) {
    //     self.alarm_id = alarm_id;
    // }
    pub fn set_last_triggered(&mut self, last_triggered: DateTime<Local>) {
        self.last_triggered = last_triggered;
    }
}

#[tokio::main]
async fn main() -> Result<()> {
    let alarm_url_result = get_env("COFFEE_ALARM_API");
    let auth_code_result = get_env("COFFEE_AUTH_CODE");
    let gpio_coffee_result = get_env("COFFEE_GPIO");

    exit_on_invalid_env_var(&alarm_url_result);
    exit_on_invalid_env_var(&auth_code_result);
    exit_on_invalid_env_var(&gpio_coffee_result);

    let coffee_pin = parse_pin(gpio_coffee_result.0);

    if coffee_pin < 0 {
        exit(ERROR_GPIO_NOT_AN_INT);
    }

    let mut triggered: Vec<Triggered> = Vec::new();

    loop {
        let latest_alarms = get_alarms(
            &alarm_url_result.0, &auth_code_result.0).await?;
        let current_date = chrono::Local::now();

        for alarm in latest_alarms {
            if alarm.hour as u32 == current_date.hour() && alarm.minute as u32 == current_date.minute() {
                match triggered.iter_mut().find(|trigger| {
                    return trigger.alarm_id == alarm._id;
                }) {
                    Some(triggered_alarm) => {
                        if triggered_alarm.last_triggered.minute() != current_date.minute() ||
                                triggered_alarm.last_triggered.hour() != current_date.hour(){
                            triggered_alarm.set_last_triggered(current_date.clone());
                            trigger_gpio(coffee_pin);
                        }
                    }
                    None => {
                        triggered.push(Triggered {
                            alarm_id: alarm._id,
                            last_triggered: current_date.clone(),
                        });
                        trigger_gpio(coffee_pin);
                    }
                }
            }
        }

        println!("{:?}", current_date);
        thread::sleep(time::Duration::from_secs(25));
    }
}

fn parse_pin(number: String) -> i32 {
    match number.parse::<i32>() {
        Ok(n) => {
            return n;
        }
        Err(err) => {
            println!("{:?}", err);
        }
    }
    return -1;
}

fn trigger_gpio(gpio: i32) {
    println!("triggered pin {}", gpio);
    let pin_result = gpio::sysfs::SysFsGpioOutput::open(gpio as u16);

    if pin_result.is_ok() {
        let mut pin = pin_result.unwrap();
        thread::spawn(move || {
            pin.set_value(true).expect("could not set gpio pin");
            thread::sleep(time::Duration::from_secs(1 * 1000));
            pin.set_value(false).expect("could not set gpio pin");
        });
    }
}

async fn get_alarms(url: &str, code: &str) -> Result<Vec<Alarm>> {
    let client = reqwest::Client::new();
    match client
        .get(url)
        .header("Authorization", format!("Bearer {}", code))
        .send().await {
        Ok(response) => {
            if response.status() == 200 {
                let alarms: Vec<Alarm> = response.json().await?;
                return Result::Ok(alarms);
            }
        }
        Err(err) => {
            println!("{:?}", err)
        }
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
