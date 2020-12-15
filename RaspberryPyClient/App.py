#!/usr/bin/env python
import threading
import logging
import os
import signal
from mqtt import Mqtt
from raspi import RaspberryClient

mqtt_username = os.getenv("COFFEE_MQTT_USERNAME")
mqtt_password = os.getenv("COFFEE_MQTT_PASSWORD")

logging.basicConfig(level=logging.INFO)

mqtt = None
thread = None
raspi_client = RaspberryClient()


def handle_exit(signum, frame):
    if mqtt is not None:
        mqtt.close()
    if thread is not None:
        thread.join()


def received(body):
    logging.debug(body)
    logging.info(f"received message {str(body)}")
    data = body.decode("utf-8")
    raspi_client.handle_request(data)


if __name__ == "__main__":
    if mqtt_username is None:
        logging.error("MQTT username is not set")
        exit(5)

    if mqtt_password is None:
        logging.error("MQTT password is not set")
        exit(10)

    mqtt = Mqtt(username=mqtt_username, password=mqtt_password)
    mqtt.connect()

    signal.signal(signal.SIGINT, handle_exit)
    signal.signal(signal.SIGTERM, handle_exit)

    mqtt.callbacks.append(received)

    thread = threading.Thread(target=mqtt.blocking_retrieve)
    thread.start()
