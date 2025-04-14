package com.RabbitMQ.Demo.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.stereotype.Repository;
import com.RabbitMQ.Demo.model.User;
import com.RabbitMQ.Demo.repository.UserRepository;

@Service
public class UserServiceImpl implements UserService {

	@Autowired
	private UserRepository userrepository;

	@Override
	public User saveUser(User user) {

		return userrepository.save(user);
	}

}
