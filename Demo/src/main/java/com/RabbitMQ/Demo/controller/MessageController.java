package com.RabbitMQ.Demo.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/Rabbit/msg")
public class MessageController {

	private RabbitMQProducer producer;

	public MessageController(RabbitMQProducer producer) {

		this.producer = producer;
	}

	// http://localhost:8080/Rabbit/msg/publish?message=text message
	@GetMapping("/publish")
	public ResponseEntity<String> sendMessage(@RequestParam("message") String message) {
		producer.sendMessage(message);
		return ResponseEntity.ok("Message sent to RabbitMQ.....");

	}

}
