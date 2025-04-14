package com.supernet.rabbitmqdemo.config;

import org.springframework.amqp.core.AmqpTemplate;
import org.springframework.amqp.core.Binding;
import org.springframework.amqp.core.BindingBuilder;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.core.TopicExchange;
import org.springframework.amqp.support.converter.Jackson2JsonMessageConverter;
import org.springframework.amqp.support.converter.MessageConverter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
@Configuration
public class RabbitMQConfig {

	@Value("${rabbitmq.exchange.name}")
	private String exchange;
	
	@Value("${rabbitmq.queue.name}")
	private String queue;
	
	@Value("${rabbitmq.routingkey}")
	private String routingkey;

	/*
	 * @Value("${rabbitmq.json_queue.name}") private String json_queue;
	 * 
	 * @Value("${rabbitmq.json_routingkey}") private String json_routingkey;
	 * 	
	//spring bean for queue to store json messages
	@Bean
	public Queue jsonQueue() {
		return new Queue(json_queue);
	}
	
	
	// binding between json queue and exchange using rounting key
	@Bean
	public Binding json_binding() {
		return BindingBuilder.bind(jsonQueue()).to(exchange()).with(json_routingkey);
	}
	
		@Bean
	public MessageConverter converter()
	{
		return new Jackson2JsonMessageConverter();
	}

	 */
	// spring beans for rabbitmq queue
	@Bean
	public Queue queue() {
		return new Queue(queue);
	}
	

	// binding between queue and exchange using rounting key
	@Bean
	public Binding binding() {
		return BindingBuilder.bind(queue()).to(exchange()).with(routingkey);
	}


	   @Bean
	    public AmqpTemplate amqpTemplate(ConnectionFactory connectionFactory){
	        RabbitTemplate rabbitTemplate = new RabbitTemplate(connectionFactory);
	        rabbitTemplate.setMessageConverter(converter());
	        return rabbitTemplate;
	    }

}
