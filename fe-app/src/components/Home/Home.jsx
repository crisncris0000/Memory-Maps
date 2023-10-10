import React from 'react';
import About from './About';
import ContactForm from './Contact';
import { BrowserRouter as Router, Route, Switch, Link  } from 'react-router-dom';

export default function Home() {
  return (
    <>
      <div className="title">
        <h1>Nostalgic Maps</h1>
        <p>
          Sharing memories that will last a lifetime
        </p>

        <button><Link>Get Started</Link></button>
      </div>

      <About />
      <ContactForm />
    </>
  )
}
