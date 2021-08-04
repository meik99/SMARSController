package com.rynkbit.arduino.smarscontroller.ui.main

import android.bluetooth.BluetoothAdapter
import android.bluetooth.BluetoothDevice
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.findNavController
import androidx.recyclerview.widget.DividerItemDecoration
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.rynkbit.arduino.smarscontroller.R
import com.rynkbit.arduino.smarscontroller.data.BluetoothDeviceRepository

class MainFragment : Fragment(), BluetoothDeviceListAdapter.BluetoothDeviceClickListener {

    companion object {
        fun newInstance() = MainFragment()
    }

    private lateinit var viewModel: MainViewModel
    private lateinit var listBluetoothDevices: RecyclerView
    private lateinit var bluetoothListAdapter: BluetoothDeviceListAdapter

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?,
                              savedInstanceState: Bundle?): View {
        return inflater.inflate(R.layout.main_fragment, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel = ViewModelProvider(this).get(MainViewModel::class.java)
        listBluetoothDevices = view.findViewById(R.id.listBluetoothDevices)
        bluetoothListAdapter = BluetoothDeviceListAdapter(listOf())

        listBluetoothDevices.adapter = bluetoothListAdapter
        listBluetoothDevices.layoutManager = LinearLayoutManager(context, LinearLayoutManager.VERTICAL, false)
        listBluetoothDevices.addItemDecoration(DividerItemDecoration(context, LinearLayoutManager.VERTICAL))

        bluetoothListAdapter.bluetoothDevices = BluetoothAdapter.getDefaultAdapter().bondedDevices.toList()
        bluetoothListAdapter.onClickListener = this
    }

    override fun onClick(bluetoothDevice: BluetoothDevice) {
        BluetoothDeviceRepository.INSTANCE.bluetoothDevice = bluetoothDevice
        requireActivity().findNavController(R.id.navigation_view).navigate(R.id.action_mainFragment_to_controllerFragment)
    }

}