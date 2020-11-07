package com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt;

import org.eclipse.paho.client.mqttv3.MqttConnectOptions;

public class DefaultMqttConnectionOptionsFactory implements MqttConnectionOptionsFactory{
    @Override
    public MqttConnectOptions createMqttConnectionOptions() {
        String mqttUrl = System.getenv("COFFEE_TO_GO_MQTT_URL");
        String username = System.getenv("COFFEE_TO_GO_MQTT_USERNAME");
        String password = System.getenv("COFFEE_TO_GO_MQTT_PASSWORD");
        MqttConnectOptions options = new MqttConnectOptions();

        if (username == null) {
            username = "";
        }
        if (password == null) {
            password = "";
        }
        if (mqttUrl == null) {
            mqttUrl = "tcp://localhost:1883";
        }

        options.setUserName(username);
        options.setPassword(password.toCharArray());
        options.setServerURIs(new String[] { mqttUrl });
        return options;
    }
}
