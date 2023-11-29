import React from 'react';
import '../../css/login.css';
import Google from '../../images/google-logo.png';
import Facebook from '../../images/facebook-logo.png';
import { Link } from 'react-router-dom';

export default function Login() {
    return (
        <div className="login-container">
            <h2>Login</h2>
            <div className="form-container">
                <div className="input-group">
                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" required />
                </div>
                
                <div className="input-group">
                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" required />
                </div>

                <Link to="/register"><button type="submit">Register</button></Link>
                <button type="submit">Login</button>
            </div>

            <p>Or login using one of the following services:</p>
            <button className="oauth-btn google">
                <img src={Google} alt="Google Logo"/>
                Login with Google
            </button>
            <button className="oauth-btn facebook">
                <img src={Facebook} alt="Facebook Logo" />
                Login with Facebook
            </button>

            <Link to="/">Return to home</Link>
        </div>
    );
}
