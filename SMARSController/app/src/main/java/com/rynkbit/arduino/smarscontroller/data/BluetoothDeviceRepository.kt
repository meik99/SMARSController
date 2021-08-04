package com.rynkbit.arduino.smarscontroller.data

import android.bluetooth.BluetoothDevice

class BluetoothDeviceRepository {
    companion object {
        val INSTANCE = BluetoothDeviceRepository()
    }

    lateinit var bluetoothDevice: BluetoothDevice
}