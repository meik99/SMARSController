package com.rynkbit.coffeetogo.controlapi.ControlAPI.rest;

import com.auth0.jwt.JWT;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.cloudant.client.api.ClientBuilder;
import com.cloudant.client.api.CloudantClient;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication.Authentication;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.brew.BrewMessageFactory;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt.DefaultMqttConnectionOptionsFactory;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.mqtt.MqttPublisher;
import org.json.simple.JSONObject;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import java.net.MalformedURLException;
import java.net.URL;
import java.util.Date;

@RestController
public class ControlController {

    @GetMapping("/brew")
    ResponseEntity<JSONObject> get(
            @RequestHeader("Authorization") String authHeader
    ) {
        String authorizationCode = "";
        if (authHeader != null) {
            String[] parts = authHeader.split(" ");
            if (parts.length > 1) {
                authorizationCode = parts[1];
            }
        }

        // Note: for Cloudant Local or Apache CouchDB use:
        CloudantClient client;
        try {
            client = ClientBuilder.url(new URL("http://localhost:5984"))
                    .username("admin")
                    .password("mysecretpassword")
                    .build();
        } catch (MalformedURLException e) {
            e.printStackTrace();
            return new ResponseEntity<>(
                    makeResponse(HttpStatus.INTERNAL_SERVER_ERROR.value(), e.getMessage()),
                    HttpStatus.INTERNAL_SERVER_ERROR);
        }

        Authentication authentication = client.database("authentication", false)
                .find(Authentication.class, authorizationCode);

        if (authentication == null) {
            return new ResponseEntity<>(makeResponse(HttpStatus.UNAUTHORIZED.value(), "invalid code"), HttpStatus.UNAUTHORIZED);
        }

        DecodedJWT decodedJWT = JWT.decode(authentication.getToken().getIdToken());
        if (new Date().after(decodedJWT.getExpiresAt())) {
            return new ResponseEntity<>(makeResponse(HttpStatus.UNAUTHORIZED.value(), "code expired"), HttpStatus.UNAUTHORIZED);
        }

        if (!authentication.getEmail().equals("michaelrynkiewicz3@gmail.com")) {
            return new ResponseEntity<>(makeResponse(HttpStatus.UNAUTHORIZED.value(), "invalid code"), HttpStatus.UNAUTHORIZED);
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
