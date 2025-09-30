// websocket.js
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function () {
  console.log("WebSocket connection established.");
};

socket.onmessage = function (event) {
  if (event.data === "reload") {
    console.log("Received reload signal, refreshing the page.");
    window.location.reload();
  }
};

socket.onclose = function () {
  console.log("WebSocket connection closed.");
};
