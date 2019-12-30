function fetchSystemLog (onEvent: Function) {
    const socket = new WebSocket('ws://localhost:10590/socket/system.log.201912230000');
    socket.addEventListener('message', function (event) {
        onEvent(event.data)
        console.log('Message from server ', event.data);
    });
}

function fetchBuildLog (onEvent:Function, buildName:string) {
    const socket = new WebSocket('ws://localhost:10590/socket/'+buildName);
    socket.addEventListener('message', function (event) {
        onEvent(event.data)
        console.log('Message from server ', event.data);
    });
}

export {
    fetchSystemLog,
    fetchBuildLog
}