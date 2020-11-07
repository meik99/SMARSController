package com.rynkbit.coffeetogo.controlapi.ControlAPI.brew;

import com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt.MqttMessageFactory;
import org.eclipse.paho.client.mqttv3.MqttMessage;
import org.json.simple.JSONObject;

import java.time.Instant;

public class BrewMessageFactory implements MqttMessageFactory {
    public MqttMessage createMessage() {
        JSONObject message = new JSONObject();

        String currentDateString = Instant.now().toString();
        message.put("command", "brew");
        message.put("date",  currentDateString);
        return new MqttMessage(message.toString().getBytes());
    }
}
