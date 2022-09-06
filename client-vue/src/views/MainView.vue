<template>
  <loader v-model:status="status" v-if="loader"/>
  <div v-if="!loader" class="row">
   <div class="header">
     <h1><span>Wall app</span></h1>
     <p>Just app, where you can write everything</p>
     <textarea v-model="text"></textarea>
     <button @click.prevent="sendMessage">Send</button>
   </div>
    <div class="body">
      <div class="card" v-for="message of messages">
        <p>{{message.text}}</p>
        <div class="date">
          <b>{{message.date.getDate()}}.{{message.date.getMonth()}}.{{message.date.getFullYear()}}</b>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import loaderComponent from "@/components/loaderComponent";
export default {
  name: 'HomeView',
  components: {
    "loader": loaderComponent
  },
  data(){
    return {
      loader: true,
      ws: null,
      status: "",
      text: "",
      messages: []
    }
  },
  created() {
    this.status = "start connection to server, please wait"
    let ws = new WebSocket("ws://172.26.20.137:3000/")
    ws.onopen = () => {
      this.ws = ws
      this.status = "Great, have access!"
      setTimeout(() => {
        this.status = "Wait! :D"
      }, 300)
      setTimeout(() => {
        this.loader = false
      }, 2000)
    }
    ws.onmessage = (message) => {
      let decodedMessage = JSON.parse(message.data)
      decodedMessage.date = new Date()
      this.messages.unshift(decodedMessage)
    }
  },
  methods: {
    sendMessage(){
      if (this.ws){
        this.ws.send(JSON.stringify({actionType: "send-message", data: this.text}))
      }
    }
  }
}
</script>

<style>
.date{
  width: 100%;
  display: flex;
  justify-content: end;
}
.date b{
  font-size: 14px;
}
.card{
  margin-top: 15px;
  padding: 10px;
  border-radius: 10px;
  border: 2px solid #282828;
  cursor: pointer;
  transition: 0.3s;
}
.card:hover{
  transition: 0.3s;
  border: 2px solid #dcb639;
  box-shadow: 0px 0px 57px -23px #2828284a;
}
.body{
  width: 450px;
  display: flex;
  justify-content: left;
  flex-direction: column;
}
.header{
  width: 450px;
  display: flex;
  justify-content: left;
  flex-direction: column;
}
.header h1{
  margin: 0px;
  position: relative;
  color: #282828;
}
.header h1::after{
  position: absolute;
  margin: 0px;
  content: "";
  background: #dcb639;
  width: 36%;
  height: 12px;
  display: flex;
  color: #282828;
  bottom: 4px;
  z-index: -1;
}
.header textarea{
  height: 150px;
  border-radius: 10px;
  width: 100%;
  resize: none;
  padding: 6px;
  box-sizing: border-box;
  font-family: 'Inter', sans-serif;
  border: 3px solid #282828;
  transition: 0.3s;
}
.header textarea:focus-visible{
  border: 3px solid #dcb639!important;
  outline: none;
  transition: 0.3s;
  border-radius: 15px;
}
.header button{
  margin-top: 5%;
  padding: 5px;
  text-transform: capitalize;
  font-weight: bold;
  font-size: 18px;
  border: 2px #282828 solid;
  border-radius: 10px;
  background: transparent;
  color: #282828;
  transition: 0.3s;
}
.header button:hover{
  border: 2px #dcb639 solid;
  border-radius: 15px;
  background:  #dcb639;
  color: #282828;
  transition: 0.3s;
  cursor: pointer;
}
.row{
  flex-direction: column;
  font-family: 'Inter', sans-serif;
  width: 60%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 5%;
}
</style>