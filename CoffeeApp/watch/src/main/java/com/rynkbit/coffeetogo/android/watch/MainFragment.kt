package com.rynkbit.coffeetogo.android.watch

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.viewModelScope
import com.google.android.gms.wearable.DataClient
import com.google.android.gms.wearable.DataEvent
import com.google.android.gms.wearable.DataEventBuffer
import com.google.android.gms.wearable.Wearable
import com.google.android.material.snackbar.Snackbar
import com.rynkbit.coffeetogo.android.common.messaging.Publisher
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class MainFragment: Fragment(), DataClient.OnDataChangedListener {
    companion object {
        fun newInstance() = MainFragment()
    }

    lateinit var btnMakeCoffee: Button
    lateinit var viewModel: MainViewModel
    lateinit var credentials: Credentials
    var username = ""
    var password = ""

    override fun onPause() {
        super.onPause()
        if (activity != null) {
            Wearable.getDataClient(requireActivity()).removeListener(this)
        }
    }

    override fun onResume() {
        super.onResume()
        if (activity != null) {
            Wearable.getDataClient(requireActivity()).addListener(this)
        }
    }


    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel = ViewModelProvider(this).get(MainViewModel::class.java)
        btnMakeCoffee = requireView().findViewById(R.id.btnMakeCoffee)

        credentials = Credentials(requireContext(), viewModel.viewModelScope)

        credentials.username.observe(viewLifecycleOwner, {
            username = it
            updateButton()
        })
        credentials.password.observe(viewLifecycleOwner, {
            password = it
            updateButton()
        })
        credentials.error.observe(viewLifecycleOwner, {
            Snackbar.make(requireView(), it, Snackbar.LENGTH_SHORT).show()
        })

        credentials.findCredentials()

    }

    private fun updateButton() {
        btnMakeCoffee.isEnabled = username.isNotEmpty() && password.isNotEmpty()
        btnMakeCoffee.setOnClickListener {
            sendMessage()
        }
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.fragment_main, container, false)
    }

    override fun onDataChanged(dataEvents: DataEventBuffer) {
        dataEvents.forEach { event ->
            if (event.type == DataEvent.TYPE_CHANGED) {
                event.dataItem.also { item ->
                    credentials.findCredentials(item)
                }
            } else {
                credentials.reset()
            }
        }
    }

    fun sendMessage() {
        lifecycleScope.launch(Dispatchers.IO) {
            val sent = async(Dispatchers.IO) {
                Publisher(username, password).trySendMessage("coffee", "make-coffee")
            }
            withContext(Dispatchers.Main) {
                if (sent.await()) {
                    Snackbar.make(requireView(), "Make coffee message sent", Snackbar.LENGTH_SHORT).show()
                } else {
                    Snackbar.make(requireView(), "Could not send message", Snackbar.LENGTH_SHORT).show()
                }
            }
        }
    }


}