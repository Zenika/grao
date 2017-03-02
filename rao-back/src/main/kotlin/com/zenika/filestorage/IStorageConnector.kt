package com.zenika.filestorage

/**
 * Created by gwennael.buchet on 22/02/17.
 */
interface IStorageConnector {
    fun connect();

    fun getNewFilesInformation();

    fun disconnect();
}