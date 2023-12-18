import React, { useState } from 'react';
import { GoogleLogin, GoogleOAuthProvider } from '@react-oauth/google';
import { useSelector, useDispatch } from 'react-redux';
import { jwtDecode } from 'jwt-decode';
import { setUser } from '../../user/userSlice';

export default function GoogleAuth() {

  const clientID = process.env.REACT_APP_CLIENT_KEY
  const dispatch = useDispatch();

  const onSuccess = (response) => {
    localStorage.setItem("token", response.credential);

    const decodedCredentials = jwtDecode(response.credential);

    const user = {
      email: decodedCredentials.email,
      firstName: decodedCredentials.given_name,
      lastName: decodedCredentials.family_name,
      role: 1,
    }

    dispatch(setUser(user));
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
