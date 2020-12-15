package com.rynkbit.coffeetogo.android.app.ui.main

import android.content.Context
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.rynkbit.coffeetogo.android.app.logic.settings.Credentials

class MainViewModel(val context: Context) : ViewModel() {
    fun updateMessagingCredentials() {
        messagingUser.postValue(Credentials().getMessagingUser(context))
        messagingPassword.postValue(Credentials().getMessagingPassword(context))
    }

    val messagingUser = MutableLiveData(Credentials().getMessagingUser(context))
    val messagingPassword = MutableLiveData(Credentials().getMessagingPassword(context))
}