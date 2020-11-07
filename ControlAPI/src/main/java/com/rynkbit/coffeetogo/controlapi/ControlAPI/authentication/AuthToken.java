package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

import com.google.gson.annotations.SerializedName;
import org.springframework.data.couchbase.core.mapping.Document;
import org.springframework.data.couchbase.core.mapping.Field;

import java.io.Serializable;

@Document
public class AuthToken implements Serializable {
    @Field("access_token")
    @SerializedName("access_token")
    private String accessToken;
    @Field("expires_in")
    @SerializedName("expires_in")
    private int expiresIn;
    @Field
    @SerializedName("scope")
    private String scope;
    @Field("token_type")
    @SerializedName("token_type")
    private String tokenType;
    @Field("id_token")
    @SerializedName("id_token")
    private String idToken;

    public String getAccessToken() {
        return accessToken;
    }

    public void setAccessToken(String accessToken) {
        this.accessToken = accessToken;
    }

    public int getExpiresIn() {
        return expiresIn;
    }

    public void setExpiresIn(int expiresIn) {
        this.expiresIn = expiresIn;
    }

    public String getScope() {
        return scope;
    }

    public void setScope(String scope) {
        this.scope = scope;
    }

    public String getTokenType() {
        return tokenType;
    }

    public void setTokenType(String tokenType) {
        this.tokenType = tokenType;
    }

    public String getIdToken() {
        return idToken;
    }

    public void setIdToken(String idToken) {
        this.idToken = idToken;
    }
}
