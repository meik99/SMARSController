package com.rynkbit.coffeetogo.android.app.ui.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import android.util.Log
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import androidx.lifecycle.Observer
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.google.android.material.snackbar.Snackbar
import com.rynkbit.coffeetogo.android.app.R
import com.rynkbit.coffeetogo.android.app.logic.settings.Credentials
import com.rynkbit.coffeetogo.android.common.messaging.Publisher
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class MainFragment : Fragment() {

    companion object {
        fun newInstance() = MainFragment()
    }

    private lateinit var viewModel: MainViewModel

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?,
                              savedInstanceState: Bundle?): View {
        return inflater.inflate(R.layout.main_fragment, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel = ViewModelProvider(this, MainViewModelFactory(requireContext())).get(MainViewModel::class.java)

        viewModel.messagingUser.observe(viewLifecycleOwner, {
            updateButton()
        })
        viewModel.messagingPassword.observe(viewLifecycleOwner, {
            updateButton()
        })
    }

    override fun onResume() {
        super.onResume()
        viewModel.updateMessagingCredentials()
    }

    fun updateButton() {
        if (viewModel.messagingUser.value?.isNotEmpty() == true &&
            viewModel.messagingPassword.value?.isNotEmpty() == true) {
            createMakeCoffeeButton()
            Credentials().putDataToWatch(requireContext(), viewModel.viewModelScope)
            Log.i(MainFragment::class.java.simpleName, "Synchronized data")
        } else {
            createSetCredentialsButton()
        }
    }

    private fun createSetCredentialsButton() {
        view?.findViewById<Button>(R.id.btnMakeCoffee)?.apply {
            setText(R.string.setup_messaging_credentials)
            setOnClickListener {
                Navigation
                    .findNavController(requireActivity(), R.id.nav_host)
                    .navigate(R.id.action_mainFragment_to_settingsFragment)
            }
        }
    }

    private fun createMakeCoffeeButton() {
        view?.findViewById<Button>(R.id.btnMakeCoffee)?.apply {
            setText(R.string.make_coffee)
            setOnClickListener {
                val username = viewModel.messagingUser.value ?: ""
                val password = viewModel.messagingPassword.value ?: ""

                viewModel.viewModelScope.launch(Dispatchers.IO) {
                    val sent = Publisher(username, password).trySendMessage("coffee", "make-coffee")

                    if (sent) {
                        Snackbar.make(requireView(), "Make coffee message sent", Snackbar.LENGTH_SHORT).show()
                    } else {
                        Snackbar.make(requireView(), "Could not send message", Snackbar.LENGTH_SHORT).show()
                    }
                }
            }
        }
    }
}