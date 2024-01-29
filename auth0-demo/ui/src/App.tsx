import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import goLogo from './assets/Go-Logo_Blue.svg'
import './App.css'
import Profile from "./components/profile.tsx";
import ApiCounter from "./components/api-counter.tsx";
import LocalCounter from "./components/local-counter.tsx";
import SecuredCounter from "./components/secured-counter.tsx";

function App() {


    return (<>
        <div>
            <a href="https://vitejs.dev" target="_blank">
                <img src={viteLogo} className="logo" alt="Vite logo"/>
            </a>
            <a href="https://react.dev" target="_blank">
                <img src={reactLogo} className="logo react" alt="React logo"/>
            </a>
            <a href="https://go.dev/" target="_blank">
                <img src={goLogo} className="logo react" alt="Go logo"/>
            </a>
        </div>
        <h1>Vite + React + Go</h1>
        <div className="card">
            <LocalCounter/>
            <ApiCounter/>
            <SecuredCounter/>
            <Profile/>
        </div>
        <p className="read-the-docs">
            Click on the Vite, React and Go logos to learn more
        </p>
    </>)
}

export default App
