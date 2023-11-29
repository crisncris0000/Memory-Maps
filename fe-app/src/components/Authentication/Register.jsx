import React from 'react';
import '../../css/register.css'

export default function Register() {
    return (
        <>
            <div className="register-container">
                <h2>Register</h2>
                <form onSubmit={null}>
                    <div className="form-group">
                        <label htmlFor="name">Name</label>
                        <input type="text" id="name" name="name" required/>
                    </div>

                    <div className="form-group">
                        <label htmlFor="email">Email</label>
                        <input type="email" id="email" name="email" required/>
                    </div>

                    <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input type="password" id="password" name="password" required/>
                    </div>

                    <button type="submit">Register</button>
                </form>
            </div>
        </>
    )
}
