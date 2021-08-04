package com.rynkbit.arduino.smarscontroller.ui.main

import android.bluetooth.BluetoothDevice
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import androidx.recyclerview.widget.RecyclerView.ViewHolder
import com.rynkbit.arduino.smarscontroller.R

class BluetoothDeviceListAdapter(var bluetoothDevices: List<BluetoothDevice>) :
    RecyclerView.Adapter<BluetoothDeviceListAdapter.BluetoothDeviceViewHolder>() {

    class BluetoothDeviceViewHolder(itemView: View) : ViewHolder(itemView) {
        val txtDeviceName = itemView.findViewById<TextView>(R.id.txtDeviceName)
        val txtDeviceMac = itemView.findViewById<TextView>(R.id.txtDeviceMac)
    }

    interface BluetoothDeviceClickListener {
        fun onClick(bluetoothDevice: BluetoothDevice)
    }

    var onClickListener: BluetoothDeviceClickListener? = null

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): BluetoothDeviceViewHolder {
        val itemView = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_bluetooth_device, parent, false)
        return BluetoothDeviceViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: BluetoothDeviceViewHolder, position: Int) {
        bluetoothDevices[position].apply {
                holder.txtDeviceName.text = name
                holder.txtDeviceMac.text = address
                holder.itemView.setOnClickListener {
                    onClickListener?.onClick(this)
                }
            }
    }

    override fun getItemCount(): Int = bluetoothDevices.size
}