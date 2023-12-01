import React, { useState } from 'react';
import '../../css/register.css';
import { Link } from 'react-router-dom';
import axios from 'axios';
 
export default function Register() {

    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = () => {
        axios.post("http://localhost:8080/users/new", {
            firstName,
            lastName,
            email,
            password
        }).then((response) => {
            console.log(response.message);
        }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <>
            <div className="register-container">
                <h2>Register</h2>
                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label className="label">First Name</label>
                        <input type="text" className="input" placeholder="Enter your first name"
                        onChange={(e) => setFirstName(e.target.value)}/>
                    </div>

                    <div className="form-group">
                        <label className="label">Last Name</label>
                        <input type="text" className="input" placeholder="Enter your last name"
                        onChange={(e) => setLastName(e.target.value)}/>
                    </div>

                    <div className="form-group">
                        <label className="label">Email</label>
                        <input type="email" className="input" placeholder="Enter your email"
                        onChange={(e) => setEmail(e.target.value)}/>
                    </div>

                    <div className="form-group">
                        <label className="label">Password</label>
                        <input type="password" className="input" placeholder="Enter your password"
                        onChange={(e) => setPassword(e.target.value)}/>
                    </div>

                    <button type="submit" className="button">
                        Register
                    </button>
                </form>
                <Link to="/" style={{"margin": "20px"}}>Return to Home</Link>
                <Link to="/login" style={{"margin": "20px"}}>Return to Login</Link>
            </div>
        </>
    )
}
