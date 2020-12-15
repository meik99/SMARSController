package com.rynkbit.coffeetogo.android.common.messaging

import com.rabbitmq.client.ConnectionFactory
import org.junit.Assert
import org.junit.Assert.*
import org.junit.Test

class PublisherTest {
    @Test
    fun testGetEnvVars() {
        val mqttUser = System.getenv("COFFEE_MQTT_USERNAME")
        val mqttPass = System.getenv("COFFEE_MQTT_PASSWORD")

        assertNotNull(mqttUser)
        assertNotNull(mqttPass)

        assertTrue(mqttUser?.isNotEmpty() ?: false)
        assertTrue(mqttPass?.isNotEmpty() ?: false)
    }

    @Test
    fun testConnection() {
        val mqttUser = System.getenv("COFFEE_MQTT_USERNAME")
        val mqttPass = System.getenv("COFFEE_MQTT_PASSWORD")
        val factory = ConnectionFactory()

        factory.host = "server.rynkbit.com"
        factory.port = 5672
        factory.username = mqttUser
        factory.password = mqttPass
        factory.useSslProtocol()

        val connection = factory.newConnection()
        val channel = connection.createChannel()

        channel.queueDeclare("coffee", false, false, false, null)
        channel.basicPublish("", "coffee", null, "make-coffee".toByteArray())
    }
}