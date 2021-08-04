package com.rynkbit.arduino.smarscontroller.ui.controller

import android.bluetooth.BluetoothAdapter
import android.content.pm.ActivityInfo
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.TextView
import androidx.core.view.WindowCompat
import androidx.core.view.WindowInsetsCompat
import androidx.core.view.WindowInsetsControllerCompat
import androidx.lifecycle.Observer
import com.rynkbit.arduino.smarscontroller.R
import com.rynkbit.arduino.smarscontroller.data.BluetoothDeviceRepository
import com.rynkbit.arduino.smarscontroller.logic.BluetoothCommunicator
import io.github.controlwear.virtual.joystick.android.JoystickView

class ControllerFragment : Fragment() {

    companion object {
        fun newInstance() = ControllerFragment()
    }

    private lateinit var viewModel: ControllerViewModel
    private lateinit var bluetoothCommunicator: BluetoothCommunicator

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.controller_fragment, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel = ViewModelProvider(this).get(ControllerViewModel::class.java)

        requireActivity().requestedOrientation = ActivityInfo.SCREEN_ORIENTATION_REVERSE_LANDSCAPE
        WindowCompat.setDecorFitsSystemWindows(requireActivity().window, false)
        WindowInsetsControllerCompat(requireActivity().window, requireView()).let { controller ->
            controller.hide(WindowInsetsCompat.Type.systemBars() or WindowInsetsCompat.Type.statusBars() or WindowInsetsCompat.Type.navigationBars())
            controller.systemBarsBehavior =
                WindowInsetsControllerCompat.BEHAVIOR_SHOW_TRANSIENT_BARS_BY_SWIPE
        }

        startBluetoothSocket()
        view.findViewById<Button>(R.id.btnConnect).setOnClickListener {
            bluetoothCommunicator.stopCommunication()
            startBluetoothSocket()
        }

        view.findViewById<JoystickView>(R.id.joystickMovement).setOnMoveListener { angle, strength ->
            requireView().findViewById<TextView>(R.id.txtStatus).text = getString(R.string.joystick_data, strength, angle)
        }
    }

    private fun startBluetoothSocket() {
        bluetoothCommunicator = BluetoothCommunicator(
            resources,
            BluetoothAdapter.getDefaultAdapter(),
            BluetoothDeviceRepository.INSTANCE.bluetoothDevice
        )
        bluetoothCommunicator.apply {
            start()
        }.status.observe(viewLifecycleOwner, {
            requireView().findViewById<TextView>(R.id.txtStatus).text = it
        })
    }
}