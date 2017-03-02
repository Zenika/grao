package com.zenika.filestorage

import com.dropbox.core.DbxRequestConfig
import com.dropbox.core.v2.DbxClientV2
import org.springframework.stereotype.Component

/**
 * Created by gwennael.buchet on 22/02/17.
 */
@Component
class DropboxConnector : IStorageConnector {

    fun init() {}

    /* fun connect() {
         println("connect")
     }

     fun getNewFilesInformation() {
         println("Get new info")
     }
    */
    private var client = null
    private val ACCESS_TOKEN = "ZldfFPz6H8AAAAAAAAACZ2qCsf3SCXkTTMxT1IkQp5BRU31Ek12p1YT96sMtGnNd"

    override open fun connect() {
        // Create Dropbox client
        println("Try to connect ...")
        val config = DbxRequestConfig("rao/dropboxconnector")
        client = DbxClientV2(config, ACCESS_TOKEN)

        println("Connected !!!")
    }

    override open fun getNewFilesInformation() {

        println("Try to get infos  ...")
        // Get current account info
        val account = client.users().currentAccount
        println(account.name.displayName)


        // Get files and folder metadata from Dropbox root directory
        var result = client.files().listFolder("")
        while (true) {
            for (metadata in result.entries) {
                System.out.println(metadata.pathLower)
            }

            if (!result.hasMore) {
                break
            }

            result = client.files().listFolderContinue(result.cursor)
        }


        println("Done...")
    }

    override fun disconnect() {
        throw UnsupportedOperationException("not implemented") //To change body of created functions use File | Settings | File Templates.
    }


}

