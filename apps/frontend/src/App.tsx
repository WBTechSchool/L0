import { useState } from "react";
import ReactJson from "react-json-view";

//@ts-expect-error mock
const BACKEND_PORT = window.BACKEND_PORT || 3001;

export const App = () => {
  const [data, setData] = useState({});
  const [input, setInput] = useState("");

  const onChangeText = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.currentTarget.value);
  };

  console.log(import.meta.env, "test");

  const onFetch = async () => {
    if (input.trim().length) {
      const res = await fetch(
        `http://localhost:${BACKEND_PORT}/order/${input}`
      );
      const data = await res.json();

      if (data) {
        setData(data);
      }
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <div style={{ marginBottom: 10 }}>
        <input
          style={{ padding: 10 }}
          placeholder="orderID"
          value={input}
          onChange={onChangeText}
        />
        <button style={{ padding: 10 }} onClick={onFetch}>
          Search
        </button>
      </div>
      <ReactJson src={data} theme="ocean" />
    </div>
  );
};
