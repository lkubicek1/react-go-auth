import {useEffect, useState} from "react";
import axios from 'axios';

const ApiCounter = () => {
    const [apiCount, setApiCount] = useState(0);

    // Function to fetch the current counter value from the API
    const fetchCurrentCount = async () => {
        try {
            const response = await axios.get(import.meta.env.VITE_API+'/counter/current');
            setApiCount(response.data.counter);
        } catch (error) {
            console.error(error);
        }
    };

    // Function to increment the counter
    const incrementApiCounter = async () => {
        try {
            const response = await axios.post(import.meta.env.VITE_API+'/counter/increment');
            setApiCount(response.data.counter);
        } catch (error) {
            console.error(error);
        }
    };

    // useEffect hook to fetch the current count on component mount
    useEffect(() => {
        fetchCurrentCount();
    }, []); // Empty dependency array ensures this runs only once on mount

    return (<div>
            <button onClick={incrementApiCounter}>Increment API Counter</button>
            <p>API Counter: {apiCount}</p>
        </div>);
};

export default ApiCounter;