package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

import org.springframework.context.annotation.Configuration;
import org.springframework.data.couchbase.config.AbstractCouchbaseConfiguration;

@Configuration
public class Config extends AbstractCouchbaseConfiguration {
    @Override
    public String getConnectionString() {
        String host = System.getenv("DB_HOST");
        return "couchbase://" + host;
    }

    @Override
    public String getUserName() {
        return System.getenv("DB_USER");
    }

    @Override
    public String getPassword() {
        return System.getenv("DB_PASSWORD");
    }

    @Override
    public String getBucketName() {
        return "authentication";
    }
}
