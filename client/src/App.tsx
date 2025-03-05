import { useEffect, useState } from 'react';
import './App.css';

function App() {

  useEffect(() => {
    new WebSocket('ws://localhost:8080/ws');
  }, []);
  return (
    <>
    a
    </>
  )
}

export default App
