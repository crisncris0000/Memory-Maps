import React, { useState } from 'react';
import '../../css/login.css';
import { Link } from 'react-router-dom';
import axios from 'axios';
import GoogleAuth from './GoogleAuth';
import { gapi } from 'gapi-script';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { setUser } from '../../user/userSlice';
import { jwtDecode } from "jwt-decode";
import { useLocation } from 'react-router-dom';
import Error from '../Messages/Error';
import { useNavigate } from 'react-router-dom';

const clientID = process.env.REACT_APP_CLIENT_KEY

export default function Login() {

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const dispatch = useDispatch();
    const location = useLocation();
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        function start() {
            gapi.client.init({
                clientID, 
                scope: ""
            })
        }
        gapi.load("client:auth2", start)
    });


    const handleSubmit = (e) => {
        e.preventDefault();

        axios.post("http://localhost:8080/users/login", {
            email,
            password,
        }).then((response) => {
            const decodedJWT = jwtDecode(response.data.token);
            dispatch(setUser(decodedJWT))
            localStorage.setItem("token", response.data.token);

            navigate("/maps")
        }).catch((error) => {
            setErrorMessage(error.response.data.message);
            setError(true);
            console.log(error);
        })
    }

    return (
        <div className="login-container">
            <h2>Login</h2>
            <form className="form-container" onSubmit={handleSubmit}>
                <div className="input-group">
                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" required value={location.state?.email} 
                    onChange={(e) => setEmail(e.target.value)}/>
                </div>
                
                <div className="input-group">
                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" required 
                    onChange={(e) => setPassword(e.target.value)}/>
                </div>
                
                <Error error={error} setError={setError} errorMessage={errorMessage}/>

                <Link to="/register"><button type="submit">Register</button></Link>
                <button type="submit">Login</button>
            </form>

            <p>Or login using one of the following services:</p>
            <GoogleAuth />

            <Link to="/">Return to home</Link>
        </div>
    );
}
