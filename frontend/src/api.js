export const sendLLMRequest = async (input) => {
  const response = await fetch('http://localhost:8080/llm', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ input }),
  });
  return response.json();
};

