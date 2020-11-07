package com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt;

import org.eclipse.paho.client.mqttv3.MqttConnectOptions;

public interface MqttConnectionOptionsFactory {
    MqttConnectOptions createMqttConnectionOptions();
}
