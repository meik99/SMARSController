package com.rynkbit.coffeetogo.android.app.logic.settings

import android.content.Context
import android.util.Log
import androidx.preference.PreferenceManager
import com.google.android.gms.wearable.PutDataMapRequest
import com.google.android.gms.wearable.Wearable

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

    fun putDataToWatch(context: Context) {
        val putDataRequest = PutDataMapRequest.create("/coffee-messaging-credentials").run {
            dataMap.putString(messagingUserKey, getMessagingUser(context))
            dataMap.putString(messagingPasswordKey, getMessagingPassword(context))
            setUrgent()
            asPutDataRequest()
        }
        val task = Wearable.getDataClient(context).putDataItem(putDataRequest)
        task.addOnCompleteListener {
            Log.i(Credentials::class.java.simpleName, "Task completed")
        }
    }
}