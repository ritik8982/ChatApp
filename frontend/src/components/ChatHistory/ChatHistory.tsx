import React from "react"
import Message from "../Message/Message"
import './ChatHistory.css'


interface IProps{
    chatHistory:Array<string>
}
const ChatHistory:React.FC<IProps> = (props) =>{
    return (
        <div className="chat-history">
            <h2>Chat History</h2>
            {props.chatHistory?.map((msg,index) => <Message key={index} message={msg}/>)}
        </div>
    )
}

export default ChatHistory