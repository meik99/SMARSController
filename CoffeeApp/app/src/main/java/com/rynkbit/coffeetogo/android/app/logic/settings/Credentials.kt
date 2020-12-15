package com.rynkbit.coffeetogo.android.app.logic.settings

import android.content.Context
import androidx.preference.PreferenceManager

class Credentials {
    companion object {
        val messagingUserKey = "pref_messaging_username"
        val messagingPasswordKey = "pref_messaging_password"
    }

    fun getMessagingUser(context: Context): String {
        val sharedPreferences = PreferenceManager.getDefaultSharedPreferences(context)
        return sharedPreferences.getString(messagingUserKey, "") ?: ""
    }

    fun getMessagingPassword(context: Context): String {
        val sharedPreferences = PreferenceManager.getDefaultSharedPreferences(context)
        return sharedPreferences.getString(messagingPasswordKey, "") ?: ""
    }
}