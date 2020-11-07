package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

import com.cloudant.client.api.ClientBuilder;
import com.cloudant.client.api.CloudantClient;

import java.net.MalformedURLException;
import java.net.URL;

public class DefaultAuthenticationDao implements AuthenticationDao {
    public Authentication findByCode(String code) {
        String host = System.getenv("DB_HOST");
        String username = System.getenv("DB_USER");
        String password = System.getenv("DB_PASSWORD");
        CloudantClient client;
        try {
            client = ClientBuilder.url(new URL("http://" + host + ":5984"))
                    .username(username)
                    .password(password)
                    .build();
        } catch (MalformedURLException e) {
            e.printStackTrace();
            return null;
        }

        Authentication authentication = client.database("authentication", false)
                .find(Authentication.class, code);
        return authentication;
    }
}
