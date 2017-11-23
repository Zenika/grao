<template>
  <div class="ref">
    <div class="idNumber">Reference #{{id+1}}
      <i @click="notifyDeleteToParent" title="Delete this reference" class="hideme fa fa-close" aria-hidden="true"></i>
    </div>
    <form>
      <ul>
        <li>
          <label for="client">Client</label>
          <input @change="notifyChangeToParent" v-model="refDataDisplay['client']" type="text" name="client" />
        </li>
        <li>
          <label for="project">Project</label>
          <input @change="notifyChangeToParent" v-model="refDataDisplay['project']" type="text" name="project" />
        </li>
        <li>
          <label for="dateref">Date</label>
          <input @change="notifyChangeToParent" v-model="refDataDisplay['date']" type="date" name="dateref" class="dateinput"/>
        </li>
        <li>
          <label for="tags">Keywords</label>
          <div class="keywords">
            <li :key="index" v-for="(keyword, index) in refDataDisplay['keywords']">
              <strong class="keyword">{{keyword}}</strong>
            </li>
            <input id="tags" type="text" name="tags" placeholder="Vuejs, Big Data, Symphony..." v-model="currentKeyword"/>
           
            <a @click="addKeyword()" class="btn btn-default">
              <i class="fa fa-plus-square" aria-hidden="true"></i>
              Add keyword
            </a>
          </div>
        </li>
        <li>
          <label for="attachments">Attachments</label>
          <input @change="sendFilesToParent($event)" type="file" name="attachments" multiple/>
        </li>
      </ul>
    </form>
  
  </div>
</template>

<script>
  /* eslint no-undef: "error" */
 
  export default {
    name: 'item-ref',
    data () {
      return {
        refDataDisplay: {client: "", project: "", date: "", keywords: []},
        currentKeyword: "",
      }
    },
    props: ['id', 'refData'],
    watch: {
      refData: function(newVal, oldVal){
        this.refDataDisplay = newVal
      }
    },
    created () {
    },
    methods: {
      addKeyword(){
        if (this.currentKeyword) {
          this.refDataDisplay.keywords.push(this.currentKeyword)
          this.currentKeyword = ""
          document.getElementById("tags").focus()
          this.notifyChangeToParent()
        }
      },
      notifyChangeToParent(){
        this.$emit("data-change", this.id, this.refDataDisplay)
      },
      notifyDeleteToParent(index){
        this.$emit("ask-delete", this.id)
      },
      sendFilesToParent(event){
        this.$emit("file-change", this.id, event.target.files)
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

  @import "../../../variables";

  .ref {
    border: solid 1px #cacaca;
    background-color: #fafafa;
    padding: 15px;
    margin-top: 15px;


    .idNumber{
      font-size: 24px;
      color: black;
      text-align: left;

      i{
        float: right;
        cursor: pointer;  
      }
    }

    .keywords{
      text-align: left;
      display: block;
      margin-top: 20px;
      li{
        display: inline;
        padding: 0;
        border: none;
        margin: 0;
      }
      .keyword{
        display: inline-block;
        background-color: $red_znk;
        margin-left: 5px;
        color: #ffffff;
        padding: 5px;
        margin-bottom: 10px;
        &:hover{
          transform: scale(0.95);
          transition: all 0.2s;
        }
      }
      
      a {
        margin-top: 20px;
      } 
      ::placeholder{
        color: #e5e5e5;
      }
    }

    form {
      max-width: 90%;
      margin-left: 50px;
      margin-right: 50px;
      border-radius: 2px;
      padding: 20px;
      
      li{
        display: block;
        padding: 9px;
        border:1px solid #DDDDDD;
        margin-bottom: 30px;
        border-radius: 3px;
        
        
        label{
          display: block;
          float: left;
          margin-top: -19px;
          height: 14px;
          padding: 2px 5px 2px 5px;
          color: black;
          font-size: 14px;
          margin-bottom: 15px;
        }

        input{
          box-sizing: border-box;
          -webkit-box-sizing: border-box;
          -moz-box-sizing: border-box;
          width: 100%;
          display: block;
          outline: none;
          border: none;
          height: 25px;
          line-height: 25px;
          font-size: 16px;
          padding: 0;
          margin-top: 10px;
          padding-left: 5px;
        }

        .dateinput{
          width: 200px;
        }
      }
      
    }

   

  }


</style>
