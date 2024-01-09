import React, { useState } from 'react';
import axios from 'axios';
import '../../css/reset-password.css';

export default function ResetPassword() {

  const [isTokenSent, setIsTokenSent] = useState(false);

  return(
    <>
      {isTokenSent === false ? <SendEmail /> : <VerifyToken />}
    </>
  );
}

function SendEmail({isTokenSent, setIsTokenSent}) {
  const [email, setEmail] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();

    axios.post(`http://localhost:8080/reset-token/new`, {
      email,
      token: "123",
    }).then((response) => {
      console.log(response);
    }).catch((error) => {
      console.log(error);
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
        </div>
        <button type="submit">Reset Password</button>
      </form>
    </div>
  );
}

function VerifyToken() {
  const [token, setToken] = useState('');
  const [verificationResult, setVerificationResult] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(`Verifying token: ${token}`);
  }
  
  return (
    <div className="verify-token-container">
      <form className="verify-token-form" onSubmit={handleSubmit}>
        <h2>Verify Token</h2>
        <div className="form-group">
          <label htmlFor="token">Token:</label>
          <input
            type="text"
            id="token"
            value={token}
            onChange={(e) => setToken(e.target.value)}
            required
          />
        </div>
        <button type="submit">Verify Token</button>
      </form>
      <div className="verification-result">
        {verificationResult && <p>{verificationResult}</p>}
      </div>
    </div>
  );
}