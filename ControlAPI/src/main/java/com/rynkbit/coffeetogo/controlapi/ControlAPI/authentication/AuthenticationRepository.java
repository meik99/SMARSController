package com.rynkbit.coffeetogo.controlapi.ControlAPI.authentication;

import org.springframework.data.repository.CrudRepository;

public interface AuthenticationRepository extends CrudRepository<Authentication, String> {
}
