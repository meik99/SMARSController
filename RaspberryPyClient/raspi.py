import logging

from gpiozero import LED
from gpiozero.pins.mock import MockFactory
from time import sleep


class RaspberryClient:
    def __init__(self, pin_number=17):
        self.pin_number = pin_number
        try:
            self.pin = LED(self.pin_number)
        except Exception as e:
            logging.error(e)
            self.pin = LED(self.pin_number, pin_factory=MockFactory())

    def handle_request(self, message):
        if message == "make-coffee":
            self.pin.on()
            sleep(5 * 60)
            self.pin.off()

