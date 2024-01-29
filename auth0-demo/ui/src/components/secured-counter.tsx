import {useEffect, useState} from 'react';
import axios from 'axios';
import {useAuth0} from "@auth0/auth0-react";

const SecuredCounter = () => {
    const [counter, setCounter] = useState(0);
    const {isAuthenticated, getAccessTokenSilently } = useAuth0();

    useEffect(() => {
        const fetchCounter = async () => {
            if (!isAuthenticated) return;

            try {
                const token = await getAccessTokenSilently({
                    authorizationParams: {
                        audience: import.meta.env.VITE_AUTH0_AUDIENCE,
                    },
                });
                const response = await axios.get(import.meta.env.VITE_API+'/user-counter/current', {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setCounter(response.data.counter);
            } catch (error) {
                console.log(error);
            }
        };

        fetchCounter();
    }, [getAccessTokenSilently, isAuthenticated]);

    const incrementCounter = async () => {
        if (!isAuthenticated) return;

        try {
            const token =  await getAccessTokenSilently({
                authorizationParams: {
                    audience: import.meta.env.VITE_AUTH0_AUDIENCE,
                },
            });
            console.log(token);
            const response = await axios.post(import.meta.env.VITE_API+'/user-counter/increment', {}, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            setCounter(response.data.counter);
        } catch (error) {
            console.error(error);
        }
    };

    return (<div>
            <button onClick={incrementCounter} disabled={!isAuthenticated}>Increment Secured Counter</button>
            <p>User-Specific Counter: {isAuthenticated ? counter : 'Not logged in'}</p>
        </div>);
};

export default SecuredCounter;
