import React, { useState } from 'react';
import axios from 'axios';

export default function ContactForm() {

    const [email, setEmail] = useState("");
    const [subject, setSubject] = useState("");
    const [body, setBody] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();

        axios.post("http://localhost:8080/users/send-email", {
            email,
            subject,
            body
        }).then((response) => {
            console.log(response.data);
        }).catch((error) => {
            console.log(error);
        })
    }

    return (
        <div className="contact-me">
            <h3>Technical issues regarding this website? Email me using this contact form:</h3>
            <form className="form-container" onSubmit={handleSubmit}>
                <div className="input-group">
                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" onChange={(e) => setEmail(e.target.value)} required />
                </div>

                <div className="input-group">
                    <label htmlFor="subject">Subject</label>
                    <input type="text" id="subject" name="subject" onChange={(e) => setSubject(e.target.value)} required />
                </div>

                <div className="input-group">
                    <label htmlFor="message">Message</label>
                    <textarea id="message" name="message" rows="4" onChange={(e) => setBody(e.target.value)} required></textarea>
                </div>
                
                <button type="submit">Send</button>
            </form>
        </div>
    );
}
