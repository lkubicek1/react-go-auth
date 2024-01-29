import React from 'react'
import ReactDOM from 'react-dom/client'
import {Auth0Provider} from '@auth0/auth0-react';
import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
    <Auth0Provider
        domain={import.meta.env.VITE_AUTH0_DOMAIN}
        clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
        authorizationParams={{
            audience: `https://react-go-auth-api/`,
            redirect_uri: window.location.origin
    }}>
        <React.StrictMode>
            <App/>
        </React.StrictMode>
    </Auth0Provider>,
    );
