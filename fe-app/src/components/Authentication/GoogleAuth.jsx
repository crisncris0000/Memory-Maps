import React, { useState } from 'react';
import { GoogleLogin, GoogleOAuthProvider } from '@react-oauth/google';
import axios from 'axios';

export default function GoogleAuth() {

  const clientID = process.env.REACT_APP_CLIENT_KEY
  const [email, setEmail] = useState("");

  const onSuccess = (response) => {
    console.log(response)
  }

  const onError = (error) => {
    console.log(error)
  }

  return (
    <GoogleOAuthProvider clientId={clientID}>
      <div className="sign-in-button">
        <GoogleLogin
          onSuccess={onSuccess}
          onError={onError}
        />
      </div>
    </GoogleOAuthProvider>
  )
}
