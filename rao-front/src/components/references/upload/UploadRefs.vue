<template>

  <div class="refsupload">
   

    <h1>
      R<span>eferences (WIP)</span>
    </h1>
    <div>
    
      <google-drive />
    </div>
    <router-link class="navbar-link" to="refs">
      <a class="btn-danger btn"><i class="fa fa-long-arrow-left " aria-hidden="true"></i>
        Back to references list
      </a>
    </router-link>

    <h3>
      Add references
    </h3>
    <transition-group tag="ul" name="reflist">
      <li :key="ref.id" v-for="(ref, index) in this.refs">
            <item-ref :id="ref.id" :refData="refs[index]" @data-change="updateRef" 
            @ask-delete="deleteReference" @file-change="updateRefFiles" />
      </li>
    </transition-group>

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
  import GoogleDrive from '../drive/GoogleDrive'
  
  export default {
    name: 'refs-upload',
    data () {
      return {
        refs: [],
        currentMaxID : 0,
        connected: false
      }
    },
    components: {
      'item-ref': Item,
      'google-drive': GoogleDrive
    },
    created () {   
      this.refs.push({id: this.currentMaxID, client: "", project: "", date: "", keywords: [], attachments: []})
    },
    methods: {
      addOneMoreRef(){
        this.refs.push({id: ++this.currentMaxID, client: "", project: "", date: "", keywords: [], attachments: []})
      },
      updateRef(data){
        let indexOfRefToUpdate = this.refs.findIndex(ref => ref.id === data['id'])
        this.refs[indexOfRefToUpdate] = data
      },
      updateRefFiles(id, files){
        let indexOfRefToUpdate = this.refs.findIndex(ref => ref.id === id)
        this.refs[indexOfRefToUpdate]['attachments'] = files
      },
      sendNewRefs(){  
          console.log("envoi serveur")
          if (this.refs.length < 1)
            console.log("c'est vide")
      },  
      deleteReference(refId){
        let indexOfRefToDelete = this.refs.findIndex(ref => ref.id === refId)
        this.refs.splice(indexOfRefToDelete, 1)
      }
  }}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  @import "../../../variables";

  .refsupload {
    text-align: center;
    margin-bottom: 100px;
    display: block;
    
    h1 {
      margin: 20px auto;
      font-size: $title_high_font_size;
      span {
        font-size: $title_low_font_size;
      }
    }
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
  
    .reflist-enter, .reflist-leave-to{
      opacity: 0;
      transform: translateY(30px);
    }

    .reflist-enter-active, .reflist-leave-active {
      transition: all 0.5s;
    }

    .btn-danger{
      background-color: $red-znk;
      width: 200px;
      font-weight: bold;
      margin-bottom: 20px;
    }
  }


</style>
