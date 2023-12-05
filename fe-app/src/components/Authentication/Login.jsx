import React, { useState } from 'react';
import '../../css/login.css';
import Google from '../../images/google-logo.png';
import Facebook from '../../images/facebook-logo.png';
import { Link } from 'react-router-dom';
import axios from 'axios';
import GoogleAuth from './GoogleAuth';
import { gapi } from 'gapi-script';
import { useEffect } from 'react';

const clientID = process.env.REACT_APP_CLIENT_KEY

export default function Login() {

    useEffect(() => {
        function start() {
            gapi.client.init({
                clientID, 
                scope: ""
            })
        }

        gapi.load("client:auth2", start)
    });

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");


    const handleSubmit = (e) => {
        e.preventDefault();


        axios.post("http://localhost:8080/users/login", {
            email, 
            password,
        }).then((response) => {
            console.log(response.data);
        }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <div className="login-container">
            <h2>Login</h2>
            <form className="form-container" onSubmit={handleSubmit}>
                <div className="input-group">
                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" required 
                    onChange={(e) => setEmail(e.target.value)}/>
                </div>
                
                <div className="input-group">
                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" required 
                    onChange={(e) => setPassword(e.target.value)}/>
                </div>

                <Link to="/register"><button type="submit">Register</button></Link>
                <button type="submit">Login</button>
            </form>

            <p>Or login using one of the following services:</p>
            <GoogleAuth />
            <button className="oauth-btn facebook">
                <img src={Facebook} alt="Facebook Logo" />
                Login with Facebook
            </button>

            <Link to="/">Return to home</Link>
        </div>
    );
}
