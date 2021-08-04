package com.rynkbit.arduino.smarscontroller.logic

import android.bluetooth.BluetoothAdapter
import android.bluetooth.BluetoothDevice
import android.bluetooth.BluetoothSocket
import android.content.res.Resources
import androidx.lifecycle.MutableLiveData
import com.rynkbit.arduino.smarscontroller.R
import java.io.IOException
import java.nio.charset.Charset

class BluetoothCommunicator(
    private val resources: Resources,
    private val bluetoothAdapter: BluetoothAdapter,
    private val bluetoothDevice: BluetoothDevice) : Thread() {

    val status = MutableLiveData<String>()
    private var statusData: String = resources.getString(R.string.disconnected)
        set(value) {
            status.postValue(value)
        }

    private var socket: BluetoothSocket? = null
    var isRunning = true

    override fun start() {
        isRunning = true
        super.start()
    }

    override fun run() {
        super.run()
        try {
            if (bluetoothDevice.uuids.isEmpty()) {
                statusData = resources.getString(R.string.no_uuids_available_for_device, bluetoothDevice.name)
                return
            }

            statusData = resources.getString(R.string.canceling_discovery)
            bluetoothAdapter.cancelDiscovery()

            statusData = resources.getString(R.string.trying_to_connect, bluetoothDevice.name)
            socket = bluetoothDevice.createRfcommSocketToServiceRecord(bluetoothDevice.uuids[0].uuid)

            socket?.let {
                it.connect()
                statusData =
                    resources.getString(R.string.socket_connected, bluetoothDevice.name)
            }
        }catch (e: IOException) {
            statusData = e.localizedMessage ?: e.message ?: resources.getString(R.string.unknown_exception)
        }
    }

    override fun interrupt() {
        super.interrupt()
        closeConnection()
    }

    fun sendData(message: String) {
        try {
            socket?.outputStream?.write(message.toByteArray(Charset.defaultCharset()))
        }catch (e: IOException) {
            statusData = e.localizedMessage ?: e.message ?: resources.getString(R.string.unknown_exception)
        }
    }

    fun stopCommunication() {
        isRunning = false
        closeConnection()
    }

    fun closeConnection() {
        if (socket?.isConnected == true) {
            socket?.close()
        }

        statusData =
            resources.getString(R.string.socket_closed, bluetoothDevice.name)
    }
}