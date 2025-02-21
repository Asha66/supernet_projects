package com.rabbitMQ.Sms;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.PropertySource;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
@PropertySource("classpath:application.properties")
public class SmsController {

    @Value("${sms.baseUrl}")
    private String baseUrl;

    @Value("${sms.version}")
    private String version;

    @Value("${sms.method}")
    private String method;

    @Value("${sms.sender}")
    private String sender;

    @Value("${sms.apiKey}")
    private String apiKey;

    @GetMapping("/send-sms")
    public String sendSms(
        @RequestParam String to, @RequestParam String message ) {
        // Construct the URL using injected properties
        String url = baseUrl + "/" + version + "/?api_key=" + apiKey +
            "&method=" + method + "&message=" +   message   + 
            "&to=" + to + "&sender=" + sender;

        // Create a RestTemplate
        RestTemplate restTemplate = new RestTemplate();

        // Send the SMS request
        ResponseEntity<String> response = restTemplate.exchange(
            url,
            HttpMethod.GET,
            null,
            String.class
        );

        // Check the response and return a result
        if (response.getStatusCode().is2xxSuccessful()) {
            return "SMS sent successfully";
        } else {
            return "Failed to send SMS";
        }
    }
}

