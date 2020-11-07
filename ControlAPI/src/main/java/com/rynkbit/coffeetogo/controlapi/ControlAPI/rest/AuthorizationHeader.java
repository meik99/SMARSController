package com.rynkbit.coffeetogo.controlapi.ControlAPI.rest;

import com.auth0.jwt.JWT;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication.Authentication;
import com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication.AuthenticationDao;

import java.util.Date;

public class AuthorizationHeader {

    private final String authorizationHeader;
    private final AuthenticationDao authenticationDao;

    public AuthorizationHeader(String authorizationHeader,
                               AuthenticationDao authenticationDao) {
        this.authorizationHeader = authorizationHeader;
        this.authenticationDao = authenticationDao;
    }

    public String getCode() {
        String authorizationCode = "";
        if (authorizationHeader != null) {
            String[] parts = authorizationHeader.split(" ");
            if (parts.length > 1) {
                authorizationCode = parts[1];
            }
        }
        return authorizationCode;
    }

    public boolean isAuthHeaderValid() {
        String authorizationCode = getCode();
        Authentication authentication = authenticationDao.findByCode(authorizationCode);
        // Note: for Cloudant Local or Apache CouchDB use:

        return !(authentication == null || !authentication.getEmail().equals("michaelrynkiewicz3@gmail.com"));
    }

    public boolean isAuthHeaderExpired() {
        String authorizationCode = getCode();
        Authentication authentication = authenticationDao.findByCode(authorizationCode);
        DecodedJWT decodedJWT = JWT.decode(authentication.getToken().getIdToken());
        return new Date().after(decodedJWT.getExpiresAt());
    }
}
