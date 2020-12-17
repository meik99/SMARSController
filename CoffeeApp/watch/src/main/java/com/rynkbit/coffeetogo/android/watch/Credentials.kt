package com.rynkbit.coffeetogo.android.watch

import android.content.Context
import android.net.Uri
import android.util.Log
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.google.android.gms.wearable.*
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class Credentials(private val context: Context, private val coroutineScope: CoroutineScope) :
    ViewModel() {
    companion object {
        val TAG: String = Credentials::class.java.simpleName
        const val DATA_PATH: String = "/coffee-messaging-credentials"
        const val MESSAGING_USER_KEY = "pref_messaging_username"
        const val MESSAGING_PASSWORD_KEY = "pref_messaging_password"
    }

    private val errorData = MutableLiveData<String>()
    private val usernameData = MutableLiveData<String>()
    private val passwordData = MutableLiveData<String>()

    val error: LiveData<String> = errorData
    val username: LiveData<String> = usernameData
    val password: LiveData<String> = passwordData

    fun findCredentials() {
        findNodes()
    }

    private fun findNodes() {
        val client = Wearable.getNodeClient(context)
        client.connectedNodes
            .addOnCompleteListener {
                Log.d(TAG, "completed querying connected nodes")
                if (it.isSuccessful) {
                    Log.d(TAG, "querying connected nodes was successful")
                    coroutineScope.launch(Dispatchers.IO) {
                        handleNodes(it.result ?: listOf())
                    }
                } else {
                    Log.d(TAG, "querying connected nodes resulted in error")
                    Log.d(TAG, it.exception?.stackTraceToString() ?: "")
                    errorData.postValue("could not query connected nodes")
                }
            }
    }

    private fun handleNodes(result: List<Node>) {
        for (node in result) {
            findDataForNodeId(node.id)
        }
    }

    private fun findDataForNodeId(id: String) {
        val uri = Uri.Builder()
            .scheme(PutDataRequest.WEAR_URI_SCHEME)
            .path(DATA_PATH)
            .authority(id)
            .build()
        Wearable.getDataClient(context).getDataItems(uri)
            .addOnCompleteListener {
                Log.d(TAG, "completed querying data items")
                if (it.isSuccessful) {
                    Log.d(TAG, "querying data items was successful")
                    coroutineScope.launch(Dispatchers.IO) {
                        if (it.result != null) {
                            handleDataItemBuffer(it.result!!)
                        }
                    }
                } else {
                    Log.d(TAG, "querying data resulted in error")
                    Log.d(TAG, it.exception?.stackTraceToString() ?: "")
                    errorData.postValue("could not query data items")
                }
            }
    }

    private fun handleDataItemBuffer(result: DataItemBuffer) {
        for (dataItem in result) {
            parseData(dataItem)
        }
    }

    private fun parseData(dataItem: DataItem) {
        val dataMap = DataMapItem.fromDataItem(dataItem).dataMap
        val user = dataMap.getString(MESSAGING_USER_KEY)
        val password = dataMap.getString(MESSAGING_PASSWORD_KEY)

        Log.d(TAG, "Username: $user")
        Log.d(TAG, "Password: $password")

        usernameData.postValue(user)
        passwordData.postValue(password)
    }

    fun findCredentials(item: DataItem) {
        parseData(item)
    }

    fun reset() {
        usernameData.postValue("")
        passwordData.postValue("")
    }
}