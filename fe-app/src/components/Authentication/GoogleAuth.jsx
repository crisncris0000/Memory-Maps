import React, { useState } from 'react';
import { GoogleLogin } from 'react-google-login';
import axios from 'axios';

export default function GoogleAuth() {

  const clientID = process.env.REACT_APP_CLIENT_KEY
  const [email, setEmail] = useState("");

  const onSuccess = (response) => {
    const obj = response.profileObj

    console.log(obj)
  }

  return (
    <div className="sign-in-button">
      <GoogleLogin 
        clientId={clientID}
        cookiePolicy="single_host_origin"
        onSuccess={onSuccess}
      />
    </div>
  )
}
