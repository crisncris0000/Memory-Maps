import React from 'react';
import '../../css/register.css'

export default function Register() {
    return (
        <>
<div className="register-container">
      <h2>Register</h2>
      <form>
        <div className="form-group">
          <label className="label">First Name</label>
          <input
            type="text"
            className="input"
            placeholder="Enter your first name"
          />
        </div>

        <div className="form-group">
          <label className="label">Last Name</label>
          <input
            type="text"
            className="input"
            placeholder="Enter your last name"
          />
        </div>

        <div className="form-group">
          <label className="label">Email</label>
          <input
            type="email"
            className="input"
            placeholder="Enter your email"
          />
        </div>

        <div className="form-group">
          <label className="label">Password</label>
          <input
            type="password"
            className="input"
            placeholder="Enter your password"
          />
        </div>

        <button type="submit" className="button">
          Register
        </button>
      </form>
    </div>
        </>
    )
}
