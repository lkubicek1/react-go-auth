import { useAuth0 } from "@auth0/auth0-react";
import LoginButton from "./login.tsx";
import LogoutButton from "./logout.tsx";

const Profile = () => {
    const { user, isAuthenticated, isLoading } = useAuth0();

    if (isLoading) {
        return <div style={{ color: 'blue' }}>Loading ...</div>;
    }

    if (isAuthenticated) {
        return (
            <div>
                <div style={{color: 'green'}}>
                    <h2>Welcome back, {user?.name}!</h2>
                    <p>Email: {user?.email}</p>
                </div>
                <LogoutButton/>
            </div>
        );
    } else {
        return (<div>
            <div style={{color: 'red'}}>User is not logged in.</div>
            <LoginButton/>
        </div>);
    }
};

export default Profile;
