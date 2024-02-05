const WebSocket = require('ws')





function connect(){

  const url = 'ws://localhost:8080/ws'
   socket = new WebSocket(url)
  
  
  console.log("Attempting Connection...");
  
  
  
  
  socket.onopen = () => {
    console.log("Successfully Connected");
    sendMsg()
    
  };
  
  
  socket.onmessage = msg => {
    console.log(msg);
  };
  
  socket.onclose = event => {
    console.log("Socket Closed Connection: ");
    retry();
 
   // setInterval(connect,2000);
  };
  
  socket.onerror = error => {
    console.log("Socket Error: ");
   retry();
   
  };
  
  
  function sendMsg(){
    socket.send("hello from the other side");
  }
  
  
  
  
}
connect()









  function retry(){

    setInterval(retry2,2000);

  }
  function retry2(){
    try{
      connect()
    }
    catch(err){
      console.log(err)
    }
  }


  


//