import { useState } from 'react';
import './App.css';

function App() {

  useState(() => {
    new WebSocket('ws://localhost:3000');
  }, []);
  return (
    <>
    a
    </>
  )
}

export default App
