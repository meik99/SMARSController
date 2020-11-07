package com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt;

import com.rynkbit.coffeetogo.controlapi.ControlAPI.rest.ControlController;
import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttConnectOptions;
import org.eclipse.paho.client.mqttv3.MqttException;

public class MqttPublisher {
    public MqttException publish(
            String topic,
            MqttMessageFactory messageFactory,
            MqttConnectionOptionsFactory optionsFactory
    ) {
        MqttClient client = null;
        MqttException exception = null;

        try {
            MqttConnectOptions options = optionsFactory
                    .createMqttConnectionOptions();
            String serverUri = "";

            if (options.getServerURIs().length > 0) {
                serverUri = options.getServerURIs()[0];
            }

            client = new MqttClient(serverUri, ControlController.class.getName());
            client.connect(options);
            client.publish(topic, messageFactory.createMessage());
        } catch (MqttException e) {
            exception = e;
        } finally {
            try {
                disconnectMqttClient(client);
            } catch (MqttException e) {
                exception = e;
            }
        }
        return exception;
    }

    private void disconnectMqttClient(MqttClient client) throws MqttException {
        if (client != null && client.isConnected()) {
            client.disconnect();
        }
    }
}
