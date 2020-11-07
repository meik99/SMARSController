package com.rynkbit.coffeetogo.controlapi.ControlAPI.rest;

import com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication.DefaultAuthenticationDao;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.brew.BrewMessageFactory;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt.DefaultMqttConnectionOptionsFactory;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt.MqttPublisher;
import org.json.simple.JSONObject;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ControlController {

    @GetMapping("/health")
    ResponseEntity<JSONObject> health() {
        return new ResponseEntity<>(makeResponse(HttpStatus.OK.value(), "success"), HttpStatus.OK);
    }

    @GetMapping("/brew")
    ResponseEntity<JSONObject> brew(
            @RequestHeader("Authorization") String authHeader
    ) {
        AuthorizationHeader authorizationHeader = new AuthorizationHeader(authHeader, new DefaultAuthenticationDao());

        if(!authorizationHeader.isAuthHeaderValid()) {
            return new ResponseEntity<>(makeResponse(HttpStatus.UNAUTHORIZED.value(), "invalid code"), HttpStatus.UNAUTHORIZED);
        }

        if(authorizationHeader.isAuthHeaderExpired()) {
            return new ResponseEntity<>(makeResponse(HttpStatus.UNAUTHORIZED.value(), "code expired"), HttpStatus.UNAUTHORIZED);
        }

        Exception e = new MqttPublisher().publish(
                "coffeetogo/control/brew",
                new BrewMessageFactory(),
                new DefaultMqttConnectionOptionsFactory()
        );

        if (e != null) {
            e.printStackTrace();
            return new ResponseEntity<>(
                    makeResponse(HttpStatus.INTERNAL_SERVER_ERROR.value(), e.getMessage()),
                    HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new ResponseEntity<>(
                makeResponse(HttpStatus.OK.value(), "success"),
                HttpStatus.OK);
    }

    private JSONObject makeResponse(int code, String message) {
        JSONObject response = new JSONObject();
        response.put("code", code);
        response.put("message", message);
        return response;
    }
}
