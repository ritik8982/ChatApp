import React, { useState } from 'react';
import './ChatInput.css'
import { sendMessage } from '../../api';


const ChatInput:React.FC = () =>{
    const[enteredMessage,setEnteredMessage] = useState<string>('');

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) =>{
        e.preventDefault();
        sendMessage(enteredMessage);
        setEnteredMessage('');

    }

    return (
        <div className="chat-input">
            <form onSubmit={handleSubmit}>
                 <input value={enteredMessage} onChange={(e) => setEnteredMessage(e.target.value) } placeholder='type a message.... Hit Enter to send the Message'/>
            </form>
        </div>
    )
}

export default ChatInput;