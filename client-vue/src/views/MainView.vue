<template>
  <loader v-if="loader"/>
  <div v-if="!loader" class="row">
   <div class="header">
     <h1><span>Wall app</span></h1>
     <p>Just app, where you can write everything</p>
     <textarea></textarea>
     <button>Send</button>
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
      loader: true
    }
  },
  created() {
    let ws = new WebSocket("ws://localhost:3000/")
    ws.onopen = () => {
      ws.send(JSON.stringify({
        action_type:"send-message",
        data:"hi sweety"
      }))
      this.loader = false
    }
  }
}
</script>

<style>
.header{
  width: 90%;
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
  font-family: 'Inter', sans-serif;
  width: 60%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 5%;
}
</style>