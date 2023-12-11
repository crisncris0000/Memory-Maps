import React, { useState } from 'react';
import { GoogleLogin, GoogleOAuthProvider } from '@react-oauth/google';
import { useSelector, useDispatch } from 'react-redux';

export default function GoogleAuth() {

  const clientID = process.env.REACT_APP_CLIENT_KEY
  const [email, setEmail] = useState("");

  const user = useSelector((state) => state.user.value);
  const dispatch = useDispatch();

  const onSuccess = (response) => {
    localStorage.setItem("token", response.credential);
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
