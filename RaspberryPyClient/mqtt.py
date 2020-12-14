import pika
import ssl
import logging
from pika.exceptions import StreamLostError


class Mqtt:
    def __init__(self, username=None, password=None, queue="coffee"):
        self.connection = None
        self.username = username
        self.password = password
        self.channel = None
        self.connected = False
        self.queue = queue
        self.callbacks = []

    def connect(self):
        if self.connected is False:
            logging.info("connecting to mqtt broker")
            self.connected = True
            self.connection = \
                pika.BlockingConnection(
                    pika.ConnectionParameters(
                        host='server.rynkbit.com',
                        port=5672,
                        credentials=
                        pika.PlainCredentials(
                            username=self.username,
                            password=self.password
                        ),
                        ssl_options=pika.SSLOptions(ssl.create_default_context())
                    ))
            self.channel = self.connection.channel()
            self.channel.queue_declare(queue=self.queue)
            logging.info("connected to mqtt broker")
        else:
            logging.error("cannot connect to mqtt broker while connection already exists. call close() first")

    def close(self):
        if self.connected is True:
            logging.info("disconnecting from mqtt broker")
            if self.channel.is_closed is False:
                try:
                    self.channel.close()
                except (AttributeError, StreamLostError):
                    logging.info("channel already closing")
            if self.connection.is_closed is False:
                try:
                    self.connection.close()
                except (AttributeError, StreamLostError):
                    logging.info("connection already closing")
            logging.info("disconnected from mqtt broker")
            self.connected = False
        else:
            logging.error("cannot disconnect from mqtt broker because no connection exists. call connect() first")

    def blocking_retrieve(self):
        if self.connected is False:
            logging.info("cannot retrieve data. connection does not exist. call connect() first")
        else:
            try:
                for method_frame, properties, body in self.channel.consume(self.queue):

                    # Display the message parts
                    logging.debug(method_frame)
                    logging.debug(properties)
                    logging.debug(body)

                    for callback in self.callbacks:
                        callback(body)

                    # Acknowledge the message
                    self.channel.basic_ack(method_frame.delivery_tag)

                    if self.connected is False:
                        break
            except Exception as e:
                if self.connected is False:
                    logging.info("retrieval of messages exited normally")
                else:
                    logging.error(e.with_traceback())
