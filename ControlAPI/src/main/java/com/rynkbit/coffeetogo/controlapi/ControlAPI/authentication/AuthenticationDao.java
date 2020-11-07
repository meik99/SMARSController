package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

public interface AuthenticationDao {
    Authentication findByCode(String code);
}
