import { useState, useEffect } from 'react';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import axios from 'axios';

interface User {
  username: string;
}

function App() {
  const [count, setCount] = useState(0);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [user, setUser] = useState<User | null>(null);


  useEffect(() => {
    checkSession();
  }, []);

  const checkSession = async () => {
    try {
      const response = await axios.get('http://localhost:8080/session', { withCredentials: true });
      if (response.data.username) {
        setUser({ username: response.data.username });
      } else {
        setUser(null);
      }
    } catch (error) {
      console.error('Session check failed', error);
      setUser(null); // Optionally handle error, e.g., by clearing user state
    }
  };

  const handleRegister = async () => {
    try {
      await axios.post('http://localhost:8080/register', { username, password });
      alert('Registration successful');
    } catch (error) {
      console.error('Registration failed', error);
      alert('Registration failed');
    }
  };

  const handleLogin = async () => {
    try {
      await axios.post('http://localhost:8080/login', { username, password }, { withCredentials: true });
      setUser({ username });
      alert('Login successful');
    } catch (error) {
      console.error('Login failed', error);
      alert('Login failed');
    }
  };

  const handleLogout = async () => {
    try {
      await axios.get('http://localhost:8080/logout', { withCredentials: true });
      setUser(null);
      alert('Logout successful');
    } catch (error) {
      console.error('Logout failed', error);
    }
  };

  return (
      <>
        <div>
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <h1>Vite + React</h1>
        {user ? (
            <div>
              <h2>Welcome, {user.username}</h2>
              <button onClick={handleLogout}>Logout</button>
            </div>
        ) : (
            <div className="card">
              <input type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
              <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
              <button onClick={handleRegister}>Register</button>
              <button onClick={handleLogin}>Login</button>
            </div>
        )}
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </>
  );
}

export default App;
