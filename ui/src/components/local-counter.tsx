import {useState} from "react";

const ApiCounter = () => {
    const [count, setCount] = useState(0)
    // Function to fetch the current counter value from the API

    return (<div>
            <button onClick={() => setCount((count) => count + 1)}>Increment Local Counter</button>
            <p>Local Counter: {count}</p>
        </div>);
};

export default ApiCounter;