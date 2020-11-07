package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

import com.google.gson.annotations.SerializedName;
import org.springframework.data.annotation.Id;
import org.springframework.data.couchbase.core.mapping.Document;
import org.springframework.data.couchbase.core.mapping.Field;

import java.io.Serializable;

@Document
public class Authentication implements Serializable {
    @Id
    @Field("_id")
    @SerializedName("_id")
    private String id;

    @Field("email")
    private String email;

    @Field("token")
    @SerializedName("token")
    private AuthToken token;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public AuthToken getToken() {
        return token;
    }

    public void setToken(AuthToken token) {
        this.token = token;
    }
}
