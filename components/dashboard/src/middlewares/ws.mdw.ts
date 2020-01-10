function watchLog(logid: number, onEvent: Function) {
    const socket = new WebSocket('ws://localhost:10590/socket/'+logid);
    socket.addEventListener('message', function (event) {
        onEvent(event.data)
    });
}

export {
    watchLog
}