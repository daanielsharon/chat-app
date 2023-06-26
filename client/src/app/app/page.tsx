"use client";
import ChatBody from "@/components/chat_body";
import httpRequest from "@/helper/axios";
import { AuthContext } from "@/modules/auth_provider";
import { WebsocketContext } from "@/modules/websocket_provider";
import autosize from "autosize";
import { useRouter } from "next/navigation";
import { useContext, useEffect, useRef, useState } from "react";

export type Message = {
  content: string;
  client_id: string;
  username: string;
  room_id: string;
  type: "recv" | "self";
};

const page = () => {
  const [messages, setMessages] = useState<Array<Message>>([]);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  const textArea = useRef<HTMLTextAreaElement>(null);
  const { conn } = useContext(WebsocketContext);
  const { user } = useContext(AuthContext);
  const router = useRouter();

  useEffect(() => {
    // get clients in the room

    if (conn === null) {
      router.push("/");
      return;
    }

    const roomId = conn.url.split("/")[5];
    async function getUsers() {
      try {
        const response = await httpRequest.get(`/ws/getClients/${roomId}`);
        setUsers(response.data);
      } catch (error) {
        console.error(error);
      }
    }

    getUsers();
  }, []);

  useEffect(() => {
    // handle websocket connection

    if (textArea.current) {
      autosize(textArea.current);
    }

    if (conn === null) {
      // connection lost, return to the home page
      router.push("/");
      return;
    }

    conn.onmessage = (message) => {
      // handle messages received from backend
      const m: Message = JSON.parse(message.data);

      console.log("message", message.data);
      if (m.content === "A new user has joined the room") {
        setUsers([...users, { username: m.username }]);
      }

      if (m.content === "User left the chat") {
        const deleteUser = users.filter((user) => user.username != m.username);
        setUsers(deleteUser);
        // display the message
        setMessages([...messages, m]);
        return;
      }

      setMessages([...messages, m]);
    };

    conn.onclose = () => {};

    conn.onerror = () => {};

    conn.onopen = () => {};
  }, [textArea, messages, conn, users]);

  const sendMessage = () => {
    if (!textArea.current?.value) return;
    // connection check
    if (conn === null) {
      router.push("/");
      return;
    }
    conn.send(textArea.current.value);
    textArea.current.value = "";
  };
  return (
    <>
      <div className="flex flex-col w-full">
        <div className="p-4 md:mx-6 mb-14">
          <ChatBody data={messages} user={user} />
        </div>
        <div className="fixed bottom-0 mt-4 w-full">
          <div className="flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md">
            <div className="flex w-full mr-4 rounded-m border border-blue">
              <textarea
                ref={textArea}
                placeholder="Type your message here!"
                className="w-full h-10 p-2 rounded-md focus:outline-none"
                style={{ resize: "none" }}
              />
            </div>
            <div className="flex items-center">
              <button
                className="p-2 rounded-md bg-blue text-white"
                onClick={sendMessage}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default page;
