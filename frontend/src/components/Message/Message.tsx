import React from "react"
import './Message.css'

interface IProps{
    key: number; 
    message: any; 
}
const Message:React.FC<IProps> = (props) =>{
    console.log(typeof props.message)
    return (
        <div className="message">
            {JSON.parse(props.message.data).body}
        </div> 
    )
}

export default Message;