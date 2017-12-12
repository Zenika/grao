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
  const SCOPES = 'https://www.googleapis.com/auth/drive';

export default {
    name: 'google-drive',
    data () {
      return {
        connected: false,
        graoDriveFolderId: "-1",
        refsFilesPaths: []
      }
    },
    props: ['refsToFetchFilesOf', 'uploadRefs','uploadAsked', 'refsToUpload'],
    watch: {
      refsToFetchFilesOf: { 
        handler: function(newVal){
          for (var i in newVal){
            this.fetchRefFilesPaths(newVal[i])
          }
          this.$emit("googleAPIFetchDone", this.refsFilesPaths)
        }
      },
      uploadAsked: {
        handler: function(newVal){
          if (newVal){
            this.uploadFiles()
          }
        }
      }
    }, 
    components: {
        'script2': Script2
    },
    created () {
      Script2.load('https://apis.google.com/js/api.js').then(() => {
        gapi.load('client:auth2', this.setupGoogleDriveAPI)
    })

      //this.$eventHub.$on("newRefToFetchFilesOf", this.fetchRefFilesPaths)
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
        
          gapi.client.load('drive', 'v3', () => {
            gapi.client.drive.files.list(
              {q: "mimeType='application/vnd.google-apps.folder' and name='GRAO-References'"} 
            ).then((response) => {
                if (response.result.files.length < 1){
                  if (!this.uploadRefs){
                    console.error("No GRAO folder was found in specified drive.")
                    this.$emit("googleAPINoGRAOFolder")
                    return
                  }
                  //Create GRAO Folder
                  this.graoDriveFolderId = this.createGraoDriveFolder()
                }
                else {
                  this.graoDriveFolderId = response.result.files[0].id
                }
                this.$emit("googleAPIReady")
            }).catch((response) => {
              if (response.status === 403){
                console.error("You must be connected to Google to do that")
                this.$emit("googleAPINotConnected")
                return
              }
            })
          });  
      },
      createGraoDriveFolder(){
        var fileMetadata = {
          'name': 'GRAO-References',
          'mimeType': 'application/vnd.google-apps.folder'
        };
        gapi.client.drive.files.create({
          resource: fileMetadata,
          fields: 'id'
        }, function (err, file) {
          if (err) {
            console.error(err);
          } else {
            console.log('Folder Id: ', file.id);
            return file.id
          }
        }).execute();
        console.log("GRAO Folder created")
        
      },
      createGraoDriveReferenceFolder(name){
        var fileMetadata = {
          'name': name,
          'mimeType': 'application/vnd.google-apps.folder'
        };
        gapi.client.drive.files.create({
          resource: fileMetadata,
          fields: 'id'
        }, function (err, file) {
          if (err) {
            console.error(err);
          } else {
            console.log('Folder Id: ', file.id);
            return file.id
          }
        }).execute();
        console.log("GRAO Folder created")
      },
      fetchRefFilesPaths(ref){
          let title = ref.objectID+"-"+ref.Date + "-" + ref.Client + "-" + ref.Project
          this.refsFilesPaths[title] = {id: "", files:[]}

          if (this.graoDriveFolderId == "-1"){
              console.error("Async error : cannot fetch reference files paths while graoFolderId is unset")
              return
          }
          
          gapi.client.drive.files.list({
              q: "mimeType='application/vnd.google-apps.folder' and name ='"+title+"' and '"
              +this.graoDriveFolderId+"' in parents"
          })
          .then((response) => { 
              if (response.result.files.length < 1){
                  console.error("Could not find any folder named "+title+" on specified drive.")
                  return
              }
              let refDriveFolderId = response.result.files[0].id
              this.refsFilesPaths[title].id = refDriveFolderId

              gapi.client.drive.files.list({
                  q: "'"+refDriveFolderId+"' in parents"
              })
              .then((resp) => {
                for (let i in resp.result.files)
                  this.refsFilesPaths[title].files.push(resp.result.files[i].name)
              })
          })
      },
      uploadFiles(){
       for (let i in this.refsToUpload){
        for (let j = 0 ; j < this.refsToUpload[i]['attachments'].length ; j++){
          let reader = new FileReader();
          reader.readAsDataURL(this.refsToUpload[i]['attachments'][j])
        
          reader.onload = (event) => {
            let imageData = event.target.result

            this.emitHttpRequestToGoogleApi(
              this.refsToUpload[i]['attachments'][j].name, 
              imageData, 
              this.refsToUpload[i]['attachments'][j].type
            )
          };
        }

       }
      },
      signInToGoogleDrive(){
        gapi.auth2.getAuthInstance().signIn().then( () => {
          this.connected = gapi.auth2.getAuthInstance().isSignedIn.get()
        })
        this.$emit("googleAPIReady")
      },
      signOutFromGoogleDrive(){
        gapi.auth2.getAuthInstance().signOut().then( () => {
          this.connected = gapi.auth2.getAuthInstance().isSignedIn.get()
        })
        this.$emit("googleAPINotConnected")
      },
      emitHttpRequestToGoogleApi(name,data,type,callback) {
        const boundary = '-------314159265358979323846';
        const delimiter = "\r\n--" + boundary + "\r\n";
        const close_delim = "\r\n--" + boundary + "--";
        data = data.split("base64,")[1]
        const contentType = type;

        var metadata = {
            'name': name,
            'mimeType': contentType
          };

          var multipartRequestBody =
              delimiter +
              'Content-Type: application/json\r\n\r\n' +
              JSON.stringify(metadata) +
              delimiter +
              'Content-Transfer-Encoding: base64\r\n' + 
              'Content-Type: ' + contentType + '\r\n' +
              '\r\n' +
              data +
              close_delim + "\r\n\r\n";

          var request = gapi.client.request({
              'path': '/upload/drive/v3/files',
              'method': 'POST',
              'params': {'uploadType': 'multipart'},
              'headers': {
                'Content-Type': 'multipart/form-data; boundary="' + boundary + '"'
              },
              'body': multipartRequestBody});
          if (!callback) {
            callback = function() {
              console.log("Upload successful")
            };
          }
          request.execute(callback);
      }
  }}
</script>

<style>
</style>