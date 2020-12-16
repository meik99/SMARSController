package com.rynkbit.coffeetogo.android.watch

import android.app.Activity
import android.net.Uri
import android.os.Bundle
import android.util.Log
import android.widget.Button
import com.google.android.gms.tasks.Tasks
import com.google.android.gms.wearable.*
import com.google.android.material.snackbar.Snackbar
import com.rynkbit.coffeetogo.android.common.messaging.Publisher
import kotlinx.coroutines.*

class MainActivity : Activity(), DataClient.OnDataChangedListener {

    val messagingUserKey = "pref_messaging_username"
    val messagingPasswordKey = "pref_messaging_password"

    /*
    * Nothing in this class works, so I hardcoded the credentials, compiled and removed it again
    * Leaving this comment to remember that I don't get how syncing data between mobile and wearables works
    * */
    var username = ""
    var password = ""

    lateinit var btnMakeCoffee: Button

    override fun onPause() {
        super.onPause()
        Wearable.getDataClient(this).removeListener(this)
    }

    override fun onResume() {
        super.onResume()
        Wearable.getDataClient(this).addListener(this, Uri.parse("wear://"), CapabilityClient.FILTER_REACHABLE)
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)


        val task = Wearable.getDataClient(this).getDataItems(Uri.parse("wear://coffee-messaging-credentials"))
        task.addOnCompleteListener {
            if (it.isSuccessful) {
                it.result?.onEach { item ->
                    checkForCredentialsItem(item)
                }
                runOnUiThread {
                    Snackbar.
                    make(findViewById(R.id.activityView), "Retrieved data", Snackbar.LENGTH_SHORT).
                    show()
                }
            } else {
                runOnUiThread {
                    Snackbar.
                    make(findViewById(R.id.activityView), "Could not retrieve data", Snackbar.LENGTH_SHORT).
                    show()
                }
            }
        }


        btnMakeCoffee = findViewById(R.id.btnMakeCoffee)
        btnMakeCoffee.isEnabled = username.isNotEmpty() && password.isNotEmpty()
        btnMakeCoffee.setOnClickListener {
            runBlocking {
                sendMessage()
            }
        }
    }

    private suspend fun getDataItems() {
        coroutineScope {
            val result = async(Dispatchers.IO) {
                val task = Wearable.getDataClient(this@MainActivity).dataItems
                Tasks.await(task)
            }
            withContext(Dispatchers.Main) {
                val buffer: DataItemBuffer = result.await()
                buffer.onEach { item ->
                    checkForCredentialsItem(item)
                }
                runOnUiThread {
                    Snackbar.
                    make(findViewById(R.id.activityView), "Retrieved data", Snackbar.LENGTH_SHORT).
                    show()
                }
            }
        }
    }

    suspend fun sendMessage() {
        coroutineScope {
            val sent = async(Dispatchers.IO) {
                Publisher(username, password).trySendMessage("coffee", "make-coffee")
            }
            withContext(Dispatchers.Main) {
                if (sent.await()) {
                    Snackbar.make(findViewById(R.id.activityView), "Make coffee message sent", Snackbar.LENGTH_SHORT).show()
                } else {
                    Snackbar.make(findViewById(R.id.activityView), "Could not send message", Snackbar.LENGTH_SHORT).show()
                }
            }
        }
    }

    override fun onDataChanged(dataEvents: DataEventBuffer) {
        dataEvents.forEach { event ->
            if (event.type == DataEvent.TYPE_CHANGED) {
                event.dataItem.also { item ->
                    checkForCredentialsItem(item)
                }
            }
        }
    }

    private fun checkForCredentialsItem(item: DataItem) {
        if (item.uri.path?.compareTo("/coffee-messaging-credentials") == 0) {
            DataMapItem.fromDataItem(item).dataMap.apply {
                username = getString(messagingUserKey)
                password = getString(messagingPasswordKey)
            }
        }

        btnMakeCoffee.isEnabled = username.isNotEmpty() && password.isNotEmpty()
        Log.i(MainActivity::class.java.simpleName, "Synchronized data")
    }

}