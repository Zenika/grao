<template>

  <div class="refsupload">
   

    <h1>
      R<span>eferences</span>
    </h1>

    <h3>
      Add references
    </h3>

    <li :key="index" v-for="(ref, index) in refs">
          <item-ref class="itemblock" :id="index" :refData="refs[index]" @data-change="updateRef" 
          @ask-delete="deleteReference" @file-change="updateRefFiles" />
    </li>
   
    <a @click="addOneMoreRef()" class="btn btn-default">
      <i class="fa fa-plus-square" aria-hidden="true"></i>
      Add one more reference
    </a>

    <a @click="sendNewRefs()" class="btn btn-default">
      <i class="fa fa-send" aria-hidden="true"></i>
      Send
    </a>

  </div>
</template>

<script>
  /* eslint no-undef: "error" */
 
  import Item from './Item'
  import google from 'googleapis'
  import consts from '../../../constants'
  const REF_MODEL = {client: "", project: "", date: "", keywords: []}

  const DRIVE_CLIENT_ID = consts.DRIVE_CLIENT_ID
  const DRIVE_API_KEY = consts.DRIVE_API_KEY
  const DRIVE_API_URL = consts.DRIVE_API_URL

  export default {
    name: 'refs-upload',
    data () {
      return {
        refs: []
      }
    },
    components: {
      'item-ref': Item
    },
    created () {
      this.refs.push(REF_MODEL)
      
    },
    methods: {
      addOneMoreRef(){
        this.refs.push(REF_MODEL)
      },
      updateRef(index, data){
        this.refs[index] = data
      },
      updateRefFiles(index, files){
        this.refs[index]['attachments'] = files
        console.log(files)
      },
      sendNewRefs(){  
          console.log("envoi serveur")
          gapi.client.init({
            apiKey: DRIVE_API_KEY,
            clientId: DRIVE_CLIENT_ID
           })

          this.$http.post(DRIVE_API_URL, this.refs[0]['attachments']).then(response => {
            console.log(response)
          })

      },
      deleteReference(index){
        this.refs.splice(index, 1)
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  @import "../../../variables";

  .refsupload {
    text-align: center;
    margin-bottom: 100px;

    h3 {
      background-color: $red_znk;
      color: white;
      padding: 5px;
      margin-left: 5%;
      margin-right: 5%;
    }

    li{
      display: block;
      margin-left: 5%;
      margin-right: 5%;
    }

    a{
      margin-top: 10px;
    }

    
  }


</style>
