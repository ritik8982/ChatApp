import React, { useEffect, useState } from 'react';
import './App.css';
import { connect } from './api';
import Header from './components/Header';
import ChatHistory from './components/ChatHistory';
import ChatInput from './components/ChatInput';

export interface IConnectProps{
    messages : Array<string>
    setMessages: React.Dispatch<React.SetStateAction<any[]>>
}
const App:React.FC = () => {

  const [chatHistory,setChatHistory] = useState<Array<string>>([]);

  useEffect(()=>{
    let props:  IConnectProps = {
      messages: chatHistory,
      setMessages: setChatHistory
    };
    connect(props);
  })
  return (
    <div className="App">
      <Header/>
      <ChatHistory chatHistory={chatHistory}/>
      <ChatInput/>

    </div>
  );
}
 
export default App;
