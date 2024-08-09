import React, { useState } from 'react';
import { sendLLMRequest } from './api';

function App() {
  const [input, setInput] = useState('');
  const [messages, setMessages] = useState([]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const userMessage = { sender: 'user', text: input };
    setMessages([...messages, userMessage]);

    const response = await sendLLMRequest(input);
    const botMessage = { sender: 'bot', text: response.output };

    setMessages([...messages, userMessage, botMessage]);
    setInput(''); // Clear input field
  };

  return (
    <div className="flex flex-col h-screen bg-gray-100">
      <header className="bg-blue-600 text-white p-4 text-center font-bold text-xl">
        LLM Chat
      </header>
      <div className="flex-grow p-4 overflow-auto">
        <div className="flex flex-col space-y-4">
          {messages.map((msg, index) => (
            <div
              key={index}
              className={`p-3 rounded-lg max-w-md ${
                msg.sender === 'user' ? 'bg-blue-500 text-white self-end' : 'bg-gray-200 self-start'
              }`}
            >
              {msg.text}
            </div>
          ))}
        </div>
      </div>
      <footer className="p-4 bg-white flex items-center">
        <form onSubmit={handleSubmit} className="w-full flex">
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            className="border p-2 flex-grow"
            placeholder="Type your message..."
            required
          />
          <button type="submit" className="ml-2 p-2 bg-blue-600 text-white rounded">
            Send
          </button>
        </form>
      </footer>
    </div>
  );
}

export default App;

