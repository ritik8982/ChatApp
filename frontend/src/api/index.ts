import { IConnectProps } from "../App";

const socket = new WebSocket('ws://localhost:9000/ws');

const connect = (props:IConnectProps) =>{
    console.log('connecting')

    socket.onopen = () =>{
        console.log('successfully connected')
    }

    // jabhi sendMessage hit hoga uske badd turant ye call ho jayega
    socket.onmessage = (msg) =>{
        console.log('message from websocket:',msg);
        props.setMessages([...props.messages,msg]);
    }

    socket.onclose = (event) =>{
        console.log('socket closed connection:',event)
    }

    socket.onerror = () =>{ 
        console.log('socket error')
    }
};

const sendMessage = (msg:string) =>{
    console.log('sending msg: ',msg);
    socket.send(msg); 
}

export {connect,sendMessage};

