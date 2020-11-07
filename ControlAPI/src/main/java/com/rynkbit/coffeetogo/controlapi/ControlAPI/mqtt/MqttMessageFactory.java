package com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt;

import org.eclipse.paho.client.mqttv3.MqttMessage;

public interface MqttMessageFactory {
    MqttMessage createMessage();
}
