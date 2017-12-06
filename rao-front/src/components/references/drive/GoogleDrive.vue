<template>
    <div>
        <button v-if="!connected" @click="signInToGoogleDrive()">Connect to Google Drive</button>
        <button v-if="connected" @click="signOutFromGoogleDrive()">Disconnect from Google Drive</button>
    </div>
</template>

<script>

  import consts from '../../../constants'
  import Script2 from 'vue-script2'
  const OAUTH_CLIENT_ID = "404476430683-b5e3agvralurokmvduaidae29131o8tc.apps.googleusercontent.com"
  const DRIVE_API_KEY = 'AIzaSyCSF7d53JN4xETyownTOsVfavbe3jXW984'
  const SCOPES = 'https://www.googleapis.com/auth/drive.metadata.readonly';

export default {
    name: 'google-drive',
    data () {
      return {
        connected: false,
        graoDriveFolderId: "-1"
      }
    },
    components: {
        'script2': Script2
    },
    created () {
      Script2.load('https://apis.google.com/js/api.js').then(() => {
        gapi.load('client:auth2', this.setupGoogleDriveAPI)
      })

      this.$eventHub.$on("newRefToFetchFilesOf", this.fetchRefFilesPaths)
    },
    methods: {
      setupGoogleDriveAPI(){
        gapi.client.init({
          apiKey: DRIVE_API_KEY,
          clientId: OAUTH_CLIENT_ID,
          scope: SCOPES
        }).then(() => {
            this.connected = gapi.auth2.getAuthInstance().isSignedIn.get()
            this.fetchGraoFolderId()
        })
      },
      fetchGraoFolderId(){
          gapi.client.load('drive', 'v2', () => {
            gapi.client.drive.files.list(
              {q: "mimeType='application/vnd.google-apps.folder' and title='GRAO-References'"} 
            ).then((response) => {
                this.graoDriveFolderId = response.result.items[0].id
                this.$emit("googleAPIReady")
            })
          });  
      },
      fetchRefFilesPaths(ref){
          let title = ref.Date + "-" + ref.Client + "-" + ref.Project

          if (this.graoDriveFolderId == "-1"){
              console.error("Async error : cannot fetch reference files paths while graoFolderId is unset")
              return
          }
          
          gapi.client.drive.files.list({
              q: "mimeType='application/vnd.google-apps.folder' and title='"+title+"' and '"
              +this.graoDriveFolderId+"' in parents"
          })
          .then((response) => {
              if (response.result.items.length < 1){
                  console.error("Could not find any folder named "+title+" on specified drive.")
                  return
              }
              let refDriveFolderId = response.result.items[0].id
              
              gapi.client.drive.files.list({
                  q: "'"+refDriveFolderId+"' in parents"
              })
              .then((resp) => {
                for (let i in resp.result.items)
                  console.log(resp.result.items[i].title) // ici
              })
          })
      },
      getAllFilesFromGraoFolder(){
          gapi.client.drive.files.list({
            q: "'"+graoDriveFolderId+"' in parents"
          }).then((response) => {
            response.result.items.forEach((item)=>console.log(item.title))
          })
      },
      signInToGoogleDrive(){
        gapi.auth2.getAuthInstance().signIn().then( () => {
          this.connected = gapi.auth2.getAuthInstance().isSignedIn.get()
        })
      },
      signOutFromGoogleDrive(){
        gapi.auth2.getAuthInstance().signOut().then( () => {
          this.connected = gapi.auth2.getAuthInstance().isSignedIn.get()
        })
      }
  }}
</script>

<style>
</style>