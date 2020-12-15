package com.rynkbit.coffeetogo.android.common.messaging

import android.util.Log
import com.rabbitmq.client.ConnectionFactory
import java.io.IOException

class Publisher(
        val username: String,
        val password: String) {

    private fun sendMessage(queue: String, message: String) {
        val factory = ConnectionFactory()

        factory.host = "server.rynkbit.com"
        factory.port = 5672
        factory.username = username
        factory.password = password
        factory.useSslProtocol()

        val connection = factory.newConnection()
        val channel = connection.createChannel()

        channel.queueDeclare(queue, false, false, false, null)
        channel.basicPublish("", queue, null, message.toByteArray())
    }

    fun trySendMessage(queue: String, message: String): Boolean {
        try {
            sendMessage(queue, message)
            return true
        }catch (e: IOException) {
            Log.e(Publisher::class.java.simpleName,
                    e.message ?:
                    "IOException occurred while sending message, but exception had to message")
        }
        return false
    }
}