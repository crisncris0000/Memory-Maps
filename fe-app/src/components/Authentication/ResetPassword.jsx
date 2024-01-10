import React, { useState } from 'react';
import axios from 'axios';
import '../../css/reset-password.css';
import Error from '../Messages/Error';

export default function ResetPassword() {

  const [isTokenSent, setIsTokenSent] = useState(false);
  const [email, setEmail] = useState('');

  return(
    <>
      {isTokenSent === false ? <SendEmail isTokenSent={isTokenSent} setIsTokenSent={setIsTokenSent} email={email} setEmail={setEmail} /> : 
      <VerifyToken email={email} />}
    </>
  );
}

function SendEmail({ isTokenSent, setIsTokenSent, email, setEmail }) {

  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();

    axios.post(`http://localhost:8080/reset-token/new`, {
      email,
    }).then((response) => {
      console.log(response);
      setIsTokenSent(true);
    }).catch((error) => {
      console.log(error);
      setError(true);
      setErrorMessage(error.response.data.message);
    })
  };

  return (
    <div className="reset-password-container">
      <form className="reset-password-form" onSubmit={handleSubmit}>
        <h2>Reset Password</h2>
        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <Error error={error} setError={setError} errorMessage={errorMessage}/>
        </div>
        <button type="submit">Reset Password</button>
      </form>
    </div>
  );
}

function VerifyToken({ email }) {
  const [token, setToken] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confrimPassword, setConfirmPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post("http://localhost:8080/reset-token", {
      email,
      password: newPassword,
      token,
    }).then((response) => {
      console.log(response.data);
    }).catch((error) => {
      console.log(error);
    })
  }
  
  return (
    <div className="verify-token-container">
      <form className="verify-token-form" onSubmit={handleSubmit}>
        <h2>Verify Token</h2>
        <div className="form-group">
          <label htmlFor="token">Token:</label>
          <input type="text" placeholder="Enter token" id="token" value={token} onChange={(e) => setToken(e.target.value)} required/>

          <label htmlFor="new-password">New password</label>
          <input type="text" placeholder="Enter your new password" name="new-password" id="new-password" value={newPassword}
          onChange={(e) => setNewPassword(e.target.value)}/>

          <label htmlFor="confirm-password">Confirm password</label>
          <input type="text" placeholder="Enter it again to confirm" name="confirm-password" id="confrim-password" value={confrimPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}/>
        </div>
        <button type="submit">Change Password</button>
      </form>
    </div>
  );
}