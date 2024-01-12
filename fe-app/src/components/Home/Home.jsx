import React from 'react';
import About from './About';
import ContactForm from './Contact';
import { Link  } from 'react-router-dom';

export default function Home() {
  return (
    <>
      <div className="title">
        <h1>Nostalgic Maps</h1>
        <p>
          Sharing memories that will last a lifetime
        </p>

        <Link to="/login" className="link">Get Started</Link>
      </div>

      <About />
      <ContactForm />
    </>
  )
}
