import pika
import ssl
import threading
import logging
from time import sleep
from pika.exceptions import *


FIVE_MINUTES = 5 * 60


class Messaging:
    def __init__(self, username=None, password=None, queue="coffee"):
        self.username = username
        self.password = password
        self.channel = None
        self.closing = False
        self.queue = queue
        self.callbacks = []
        self.connection = self._create_connection()

    def _create_connection(self):
        return pika.BlockingConnection(
            pika.ConnectionParameters(
                host='server.rynkbit.com',
                port=5672,
                credentials=
                pika.PlainCredentials(
                    username=self.username,
                    password=self.password
                ),
                ssl_options=pika.SSLOptions(ssl.create_default_context()),
                heartbeat=600,
                blocked_connection_timeout=300
            ))

    def listen(self):
        while self.connection is not None and not self.closing:
            self._try_listen()

    def _try_listen(self):
        try:
            self._listen()
        except Exception as e:
            logging.error(e)
            self.connection = self._create_connection()

    def _listen(self):
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=self.queue)
        self.channel.basic_consume(queue=self.queue, on_message_callback=self._handle_message)
        self.channel.start_consuming()

    def _handle_message(self, consumer_info, method_frame, properties, body):
        self.channel.basic_ack(delivery_tag=method_frame.delivery_tag)
        for callback in self.callbacks:
            threading.Thread(target=callback, args=[body], daemon=True).start()

    def close(self):
        self.closing = True
        if self.connection is not None:
            try:
                self._close()
            except (StreamLostError, IndexError):
                logging.info("stream already closed")

    def _close(self):
        self.connection.close()

    def open(self):
        self.closing = False
        self.connection = self._create_connection()
