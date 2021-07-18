import React, {useEffect, useState} from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card } from 'antd';
import ChatBox from 'react-chat-plugin';
import { w3cwebsocket as W3CWebSocket } from "websocket";

const username = localStorage.getItem("FullStack-username");
const user = localStorage.getItem("FullStack-user");

const client = new W3CWebSocket('ws://127.0.0.1:8080/ws' + "?username=" + JSON.parse(username));

export default () => {

  const userName = JSON.parse(username);



  const message = [
      {
        text: userName + ' has joined the conversation',
        timestamp: +new Date(),
        type: 'notification',
      },
  ];

  const [messages, concatMessage] = useState(message);

  useEffect(()=> {

    client.onopen = () => {
      console.log('WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      const data = JSON.parse(message.data);
      console.log(data);

      concatMessage([
        ...messages,
        data
      ])
    };
  })

  const handleOnSendMessage = (text) => {

    const msgBox = {
      author: {
        username: userName,
        id: userName,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: text,
      timestamp: +new Date(),
      type: 'text',
    }

    client.send(JSON.stringify(msgBox));

  };

  return (
    <PageContainer>
      <Card
        title = "Full Stack Chat Room"
      >


        <ChatBox
          messages={messages}
          userId={userName}
          onSendMessage={handleOnSendMessage}
          width={'1000px'}
          height={'500px'}
        />
      </Card>



    </PageContainer>
  );
};
