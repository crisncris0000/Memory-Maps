import React, { useEffect, useState } from 'react';
import '../../css/register.css';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Error from '../Messages/Error';
 
export default function Register() {

    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = (e) => {

        e.preventDefault();

        axios.post("http://localhost:8080/users/new", {
            firstName,
            lastName,
            email,
            password
        }).then((response) => {
            console.log(response);
            navigate('/login', {state: {email}});
        }).catch((error) => {
            setError(true);
            console.log(error.response);
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
                        onChange={(e) => setFirstName(e.target.value)} required/>
                    </div>

                    <div className="form-group">
                        <label className="label">Last Name</label>
                        <input type="text" className="input" placeholder="Enter your last name"
                        onChange={(e) => setLastName(e.target.value)} required/>
                    </div>

                    <div className="form-group">
                        <label className="label">Email</label>
                        <input type="email" className="input" placeholder="Enter your email"
                        onChange={(e) => setEmail(e.target.value)} required/>
                        { error ? <Error errorMessage={"User already exists with that email"} error={error} setError={setError}/> : null }
                    </div>

                    <div className="form-group">
                        <label className="label">Password</label>
                        <input type="password" className="input" placeholder="Enter your password"
                        onChange={(e) => setPassword(e.target.value)} required/>
                    </div>

                    <button type="submit" className="button">
                        Register
                    </button>
                </form>
                <Link to="/" style={{margin: "30px"}}>Return to Home</Link>
                <Link to="/login" style={{margin: "30px"}}>Return to Login</Link>
            </div>
        </>
    )
}
