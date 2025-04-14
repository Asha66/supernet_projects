package com.RabbitMQ.Demo.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.RabbitMQ.Demo.model.User;


@Repository
public interface UserRepository extends JpaRepository<User, Long> {
	
}
